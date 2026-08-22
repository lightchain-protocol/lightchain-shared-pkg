package blob

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SearchWindowBlocks is the number of EL blocks to search backward when the
// blob is not found at the expected slot. This handles the two-TX flow where
// the gateway submits the blob TX in an earlier block than the user's
// submitJob TX.
const SearchWindowBlocks = 100

// MaxMissedSlots bounds how far past the parent beacon block to look for the
// block that actually carried a payload.
//
// EIP-4788 gives an execution block the root of the *parent* beacon block, and
// the payload normally rides in the very next slot. That only holds when no
// slot is missed. On a small validator set misses are routine - roughly 8% of
// slots on this devnet - and every one of them shifts the payload a slot later.
// Assuming slot+1 made blob fetches fail permanently whenever the intervening
// slot had no proposer, which surfaced as chat requests hanging forever.
const MaxMissedSlots = 32

// ELBlockFetcher is the subset of ethclient used to fetch execution-layer blocks.
// Defined as an interface for testability.
type ELBlockFetcher interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

// BeaconClient fetches blob sidecars from the CL Beacon API.
type BeaconClient struct {
	beaconURL  string
	elFetcher  ELBlockFetcher
	httpClient *http.Client
	maxRetries int
}

// NewBeaconClient creates a BeaconClient for fetching blob data.
func NewBeaconClient(beaconURL string, elFetcher ELBlockFetcher, timeout time.Duration, maxRetries int) *BeaconClient {
	return &BeaconClient{
		beaconURL:  beaconURL,
		elFetcher:  elFetcher,
		httpClient: &http.Client{Timeout: timeout},
		maxRetries: maxRetries,
	}
}

// beaconHeaderResponse is the JSON envelope for GET /eth/v1/beacon/headers/{id}.
type beaconHeaderResponse struct {
	Data struct {
		Header struct {
			Message struct {
				Slot string `json:"slot"`
			} `json:"message"`
		} `json:"header"`
	} `json:"data"`
}

// beaconHeaderListResponse is the envelope for the filtered form of the same
// endpoint, GET /eth/v1/beacon/headers?parent_root=..., which returns an array.
type beaconHeaderListResponse struct {
	Data []struct {
		Header struct {
			Message struct {
				Slot string `json:"slot"`
			} `json:"message"`
		} `json:"header"`
	} `json:"data"`
}

// blobSidecarsResponse is the JSON envelope for GET /eth/v1/beacon/blob_sidecars/{slot}.
type blobSidecarsResponse struct {
	Data []BlobSidecar `json:"data"`
}

// BlobSidecar represents a single blob sidecar from the Beacon API.
type BlobSidecar struct {
	Index         string `json:"index"`
	Blob          string `json:"blob"`
	KzgCommitment string `json:"kzg_commitment"`
}

// FetchBlob retrieves blob data for a given versioned hash from the Beacon API.
// Steps: EL block → parentBeaconBlockRoot → CL header (slot) → blob sidecars → match by versioned hash.
func (bc *BeaconClient) FetchBlob(ctx context.Context, versionedHash common.Hash, blockNumber uint64) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt <= bc.maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 100ms, 200ms, 400ms, ...
			backoff := time.Duration(100<<uint(attempt-1)) * time.Millisecond
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		data, err := bc.fetchBlobOnce(ctx, versionedHash, blockNumber)
		if err != nil {
			lastErr = err
			continue
		}
		return data, nil
	}
	return nil, fmt.Errorf("fetch blob after %d retries: %w", bc.maxRetries, lastErr)
}

func (bc *BeaconClient) fetchBlobOnce(ctx context.Context, versionedHash common.Hash, blockNumber uint64) ([]byte, error) {
	// Search from the submitJob block backward through recent blocks.
	// The blob TX was submitted by the gateway before the user called
	// submitJob, so it may be in any of the preceding blocks.
	startBlock := blockNumber
	endBlock := uint64(0)
	if startBlock > SearchWindowBlocks {
		endBlock = startBlock - SearchWindowBlocks
	}

	var lastErr error
	for blk := startBlock; blk >= endBlock && blk <= startBlock; blk-- {
		data, err := bc.fetchBlobAtBlock(ctx, versionedHash, blk)
		if err == nil {
			return data, nil
		}
		lastErr = err
		// "not found" errors → try the previous block.
		// RPC or decode errors → stop immediately.
		if !IsNotFoundError(err) {
			return nil, fmt.Errorf("blob search failed at block %d: %w", blk, err)
		}
	}

	return nil, fmt.Errorf("blob with versioned hash %s not found in blocks %d..%d: %w", versionedHash.Hex(), endBlock, startBlock, lastErr)
}

func (bc *BeaconClient) fetchBlobAtBlock(ctx context.Context, versionedHash common.Hash, blockNumber uint64) ([]byte, error) {
	block, err := bc.elFetcher.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("fetch EL block %d: %w", blockNumber, err)
	}

	parentBeaconRoot := block.BeaconRoot()
	if parentBeaconRoot == nil {
		return nil, fmt.Errorf("block %d has no parentBeaconBlockRoot (pre-Deneb?)", blockNumber)
	}

	// Ask the beacon directly which block has this parent. That is the block
	// carrying our payload, in one request, regardless of how many slots were
	// missed in between.
	blobSlot, err := bc.childSlotOf(ctx, parentBeaconRoot.Hex())
	if err != nil {
		// Older or non-conforming beacons may not support the parent_root
		// filter. Fall back to resolving the parent's slot and walking forward.
		parentSlot, perr := bc.getBeaconSlot(ctx, parentBeaconRoot.Hex())
		if perr != nil {
			return nil, fmt.Errorf("get beacon slot for block %d: %w", blockNumber, perr)
		}
		var sidecars []BlobSidecar
		sidecars, blobSlot, err = bc.sidecarsAfterSlot(ctx, parentSlot)
		if err != nil {
			return nil, fmt.Errorf("get blob sidecars after slot %d (block %d): %w", parentSlot, blockNumber, err)
		}
		return bc.matchSidecar(sidecars, versionedHash, blobSlot, blockNumber)
	}

	sidecars, err := bc.getBlobSidecars(ctx, blobSlot)
	if err != nil {
		return nil, fmt.Errorf("get blob sidecars at slot %d (block %d): %w", blobSlot, blockNumber, err)
	}

	return bc.matchSidecar(sidecars, versionedHash, blobSlot, blockNumber)
}

// matchSidecar picks the sidecar carrying versionedHash and decodes it.
func (bc *BeaconClient) matchSidecar(
	sidecars []BlobSidecar,
	versionedHash common.Hash,
	blobSlot uint64,
	blockNumber uint64,
) ([]byte, error) {
	for _, sc := range sidecars {
		commitHash := KzgToVersionedHash(sc.KzgCommitment)
		if commitHash == (common.Hash{}) {
			continue // skip sidecar with malformed KZG commitment
		}
		if commitHash == versionedHash {
			blobData, err := hex.DecodeString(StripHexPrefix(sc.Blob))
			if err != nil {
				return nil, fmt.Errorf("decode blob hex: %w", err)
			}

			payload, err := DecodeBlobData(blobData)
			if err != nil {
				return nil, fmt.Errorf("decode blob payload: %w", err)
			}
			return payload, nil
		}
	}

	return nil, fmt.Errorf("blob with versioned hash %s not found in slot %d sidecars (block %d)", versionedHash.Hex(), blobSlot, blockNumber)
}

// childSlotOf returns the slot of the beacon block whose parent is blockRoot.
//
// An execution block carries its *parent* beacon block root (EIP-4788), so this
// is what turns that into the slot actually holding the payload. Asking the
// beacon to filter by parent_root costs one request and is exact, which matters
// because the caller may repeat this for up to SearchWindowBlocks blocks: the
// previous approach of walking forward slot by slot multiplied out to thousands
// of requests and made a fetch take a minute.
func (bc *BeaconClient) childSlotOf(ctx context.Context, blockRoot string) (uint64, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers?parent_root=%s", bc.beaconURL, blockRoot)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return 0, fmt.Errorf("create child header request: %w", err)
	}

	resp, err := bc.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("child header request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read child header response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("child header returned status %d: %s", resp.StatusCode, string(body))
	}

	var listResp beaconHeaderListResponse
	if err := json.Unmarshal(body, &listResp); err != nil {
		return 0, fmt.Errorf("unmarshal child header: %w", err)
	}
	if len(listResp.Data) == 0 {
		return 0, fmt.Errorf("no beacon block has parent_root %s", blockRoot)
	}

	var slot uint64
	if _, err := fmt.Sscanf(listResp.Data[0].Header.Message.Slot, "%d", &slot); err != nil {
		return 0, fmt.Errorf("parse child slot %q: %w", listResp.Data[0].Header.Message.Slot, err)
	}
	return slot, nil
}

// sidecarsAfterSlot returns the sidecars of the first beacon block found after
// parentSlot, along with the slot it came from.
//
// Missed slots are skipped rather than treated as failures: the beacon replies
// 404 "no blocks found at slot N" for a slot nobody proposed, which says
// nothing about whether the payload exists. Only a genuine transport or decode
// error aborts the walk.
func (bc *BeaconClient) sidecarsAfterSlot(ctx context.Context, parentSlot uint64) ([]BlobSidecar, uint64, error) {
	var lastErr error
	for offset := uint64(1); offset <= MaxMissedSlots; offset++ {
		slot := parentSlot + offset

		sidecars, err := bc.getBlobSidecars(ctx, slot)
		if err == nil {
			return sidecars, slot, nil
		}
		if !isMissingSlot(err) {
			return nil, slot, err
		}
		lastErr = err

		if ctx.Err() != nil {
			return nil, slot, ctx.Err()
		}
	}

	return nil, 0, fmt.Errorf("no beacon block in slots %d..%d: %w",
		parentSlot+1, parentSlot+MaxMissedSlots, lastErr)
}

// isMissingSlot reports whether the beacon simply had no block at that slot, as
// opposed to failing to answer. Prysm returns 404 with "no blocks found at
// slot" for a missed proposal.
func isMissingSlot(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "status 404") ||
		strings.Contains(msg, "Block not found") ||
		strings.Contains(msg, "NOT_FOUND")
}

func (bc *BeaconClient) getBeaconSlot(ctx context.Context, blockRoot string) (uint64, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%s", bc.beaconURL, blockRoot)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return 0, fmt.Errorf("create beacon header request: %w", err)
	}

	resp, err := bc.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("beacon header request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read beacon header response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("beacon header returned status %d: %s", resp.StatusCode, string(body))
	}

	var headerResp beaconHeaderResponse
	if err := json.Unmarshal(body, &headerResp); err != nil {
		return 0, fmt.Errorf("unmarshal beacon header: %w", err)
	}

	var slot uint64
	if _, err := fmt.Sscanf(headerResp.Data.Header.Message.Slot, "%d", &slot); err != nil {
		return 0, fmt.Errorf("parse slot %q: %w", headerResp.Data.Header.Message.Slot, err)
	}

	return slot, nil
}

func (bc *BeaconClient) getBlobSidecars(ctx context.Context, slot uint64) ([]BlobSidecar, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/blob_sidecars/%d", bc.beaconURL, slot)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create blob sidecars request: %w", err)
	}

	resp, err := bc.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("blob sidecars request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read blob sidecars response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("blob sidecars returned status %d: %s", resp.StatusCode, string(body))
	}

	var sidecarsResp blobSidecarsResponse
	if err := json.Unmarshal(body, &sidecarsResp); err != nil {
		return nil, fmt.Errorf("unmarshal blob sidecars: %w", err)
	}

	return sidecarsResp.Data, nil
}

// IsNotFoundError returns true if the error indicates the blob was simply not
// present for the given block (expected during backward search). Returns false
// for infrastructure errors (RPC failures, decode errors) that should stop the search.
//
// This must recognise a missed slot as well as an absent blob. It previously
// matched only "not found in slot", so the beacon's 404 for an unproposed slot
// was classified as infrastructure failure and aborted the entire backward
// search on the first missed slot - turning a recoverable miss into a permanent
// job failure.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "not found in slot") ||
		strings.Contains(msg, "not found in blocks") ||
		strings.Contains(msg, "no beacon block in slots") ||
		isMissingSlot(err)
}
