package blob

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- test helpers ---

type mockELFetcher struct {
	block *types.Block
	err   error
}

func (m *mockELFetcher) BlockByNumber(_ context.Context, _ *big.Int) (*types.Block, error) {
	return m.block, m.err
}

func mockELBlock(beaconRoot common.Hash) *types.Block {
	header := &types.Header{
		Number:           big.NewInt(100),
		ParentBeaconRoot: &beaconRoot,
	}
	return types.NewBlockWithHeader(header)
}

func fakeVersionedHash(label string) ([]byte, string, common.Hash) {
	commitment := make([]byte, 48)
	copy(commitment, label)
	commitmentHex := "0x" + hex.EncodeToString(commitment)
	h := sha256.Sum256(commitment)
	h[0] = 0x01
	return commitment, commitmentHex, common.Hash(h)
}

// canonicalHeader is the /eth/v1/beacon/headers?parent_root= reply for one canonical child at slot.
func canonicalHeader(slot string) string {
	return `{"data":[{"canonical":true,"header":{"message":{"slot":"` + slot + `"}}}]}`
}

// --- tests ---

func TestBeaconClient_FetchBlob_Success(t *testing.T) {
	t.Parallel()

	testPayload := []byte("hello blob world")
	blob, err := EncodeBlobData(testPayload)
	require.NoError(t, err)
	blobHex := "0x" + hex.EncodeToString(blob[:])

	_, commitmentHex, expectedHash := fakeVersionedHash("test-commitment")
	beaconRoot := common.HexToHash("0xbeaconroot")

	beacon := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/eth/v1/beacon/headers" && r.URL.Query().Get("parent_root") == beaconRoot.Hex():
			_, _ = fmt.Fprint(w, canonicalHeader("101"))
		case r.URL.Path == "/eth/v1/beacon/blob_sidecars/101":
			resp := blobSidecarsResponse{
				Data: []BlobSidecar{{Index: "0", Blob: blobHex, KzgCommitment: commitmentHex}},
			}
			_ = json.NewEncoder(w).Encode(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer beacon.Close()

	bc := NewBeaconClient(beacon.URL, &mockELFetcher{block: mockELBlock(beaconRoot)}, 5*time.Second, 0)
	result, err := bc.FetchBlob(context.Background(), expectedHash, 100)
	require.NoError(t, err)
	assert.Equal(t, testPayload, result)
}

// The slot after the parent can be empty (a proposer missed it): the block's
// beacon block then sits later than parent+1, and its sidecars with it.
func TestBeaconClient_FetchBlob_EmptySlotAfterParent(t *testing.T) {
	t.Parallel()

	testPayload := []byte("after an empty slot")
	blob, err := EncodeBlobData(testPayload)
	require.NoError(t, err)
	blobHex := "0x" + hex.EncodeToString(blob[:])

	_, commitmentHex, expectedHash := fakeVersionedHash("empty-slot-commitment")
	beaconRoot := common.HexToHash("0xbeaconroot")

	// parent at slot 100, slot 101 empty, the block's own beacon block at 102
	beacon := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/eth/v1/beacon/headers" && r.URL.Query().Get("parent_root") == beaconRoot.Hex():
			_, _ = fmt.Fprint(w, `{"data":[{"canonical":false,"header":{"message":{"slot":"101"}}},{"canonical":true,"header":{"message":{"slot":"102"}}}]}`)
		case r.URL.Path == "/eth/v1/beacon/blob_sidecars/101":
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprint(w, `{"message":"Block not found: no blocks found at slot 101","code":404}`)
		case r.URL.Path == "/eth/v1/beacon/blob_sidecars/102":
			resp := blobSidecarsResponse{
				Data: []BlobSidecar{{Index: "0", Blob: blobHex, KzgCommitment: commitmentHex}},
			}
			_ = json.NewEncoder(w).Encode(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer beacon.Close()

	bc := NewBeaconClient(beacon.URL, &mockELFetcher{block: mockELBlock(beaconRoot)}, 5*time.Second, 0)
	result, err := bc.FetchBlob(context.Background(), expectedHash, 100)
	require.NoError(t, err)
	assert.Equal(t, testPayload, result)
}

func TestBeaconClient_FetchBlob_NotFound(t *testing.T) {
	t.Parallel()

	beaconRoot := common.HexToHash("0xbeaconroot")

	beacon := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/eth/v1/beacon/headers" && r.URL.Query().Get("parent_root") == beaconRoot.Hex():
			_, _ = fmt.Fprint(w, canonicalHeader("101"))
		case r.URL.Path == "/eth/v1/beacon/blob_sidecars/101":
			resp := blobSidecarsResponse{Data: []BlobSidecar{}}
			_ = json.NewEncoder(w).Encode(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer beacon.Close()

	bc := NewBeaconClient(beacon.URL, &mockELFetcher{block: mockELBlock(beaconRoot)}, 5*time.Second, 0)
	_, err := bc.FetchBlob(context.Background(), common.HexToHash("0xdeadbeef"), 100)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found in blocks")
}

func TestBeaconClient_FetchBlob_BeaconError(t *testing.T) {
	t.Parallel()

	beaconRoot := common.HexToHash("0xbeaconroot")

	beacon := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "internal server error")
	}))
	defer beacon.Close()

	bc := NewBeaconClient(beacon.URL, &mockELFetcher{block: mockELBlock(beaconRoot)}, 5*time.Second, 0)
	_, err := bc.FetchBlob(context.Background(), common.Hash{}, 100)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "blob search failed")
}

func TestBeaconClient_FetchBlob_ELBlockError(t *testing.T) {
	t.Parallel()

	bc := NewBeaconClient("http://unused", &mockELFetcher{err: fmt.Errorf("rpc unavailable")}, 5*time.Second, 0)
	_, err := bc.FetchBlob(context.Background(), common.Hash{}, 100)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "blob search failed")
}

func TestBeaconClient_FetchBlob_RetryOnFailure(t *testing.T) {
	t.Parallel()

	testPayload := []byte("retry-payload")
	blob, err := EncodeBlobData(testPayload)
	require.NoError(t, err)
	blobHex := "0x" + hex.EncodeToString(blob[:])

	_, commitmentHex, expectedHash := fakeVersionedHash("retry-commitment")
	beaconRoot := common.HexToHash("0xbeaconroot")

	callCount := 0
	beacon := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/eth/v1/beacon/headers" && r.URL.Query().Get("parent_root") == beaconRoot.Hex():
			callCount++
			if callCount <= 1 {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprint(w, "temporary error")
				return
			}
			_, _ = fmt.Fprint(w, canonicalHeader("51"))
		case r.URL.Path == "/eth/v1/beacon/blob_sidecars/51":
			resp := blobSidecarsResponse{
				Data: []BlobSidecar{{Index: "0", Blob: blobHex, KzgCommitment: commitmentHex}},
			}
			_ = json.NewEncoder(w).Encode(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer beacon.Close()

	bc := NewBeaconClient(beacon.URL, &mockELFetcher{block: mockELBlock(beaconRoot)}, 5*time.Second, 2)
	result, err := bc.FetchBlob(context.Background(), expectedHash, 100)
	require.NoError(t, err)
	assert.Equal(t, testPayload, result)
}

func TestBeaconClient_GetBeaconSlot_ServerError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "internal error")
	}))
	defer srv.Close()

	bc := NewBeaconClient(srv.URL, nil, 5*time.Second, 0)
	_, err := bc.getBeaconSlot(context.Background(), "0xdeadbeef")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "status 500")
}

func TestBeaconClient_GetBlobSidecars_WithData(t *testing.T) {
	t.Parallel()

	blobData := make([]byte, 64)
	for i := range blobData {
		blobData[i] = byte(i)
	}
	blobHex := "0x" + hex.EncodeToString(blobData)

	_, commitHex, _ := fakeVersionedHash("sidecar-test")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		resp := blobSidecarsResponse{
			Data: []BlobSidecar{{Index: "0", Blob: blobHex, KzgCommitment: commitHex}},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	bc := NewBeaconClient(srv.URL, nil, 5*time.Second, 0)
	sidecars, err := bc.getBlobSidecars(context.Background(), 42)
	require.NoError(t, err)
	require.Len(t, sidecars, 1)
	assert.Equal(t, commitHex, sidecars[0].KzgCommitment)

	expectedHash := KzgToVersionedHash(commitHex)
	assert.Equal(t, byte(0x01), expectedHash[0])
}

func TestBeaconClient_SearchWindowBlocks(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint64(100), uint64(SearchWindowBlocks))
}

func TestBeaconClient_ContextCancellation(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(10 * time.Second)
	}))
	defer srv.Close()

	bc := NewBeaconClient(srv.URL, &mockELFetcher{block: mockELBlock(common.Hash{})}, 5*time.Second, 2)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	_, err := bc.FetchBlob(ctx, common.Hash{}, 100)
	require.Error(t, err)
}

func TestIsNotFoundError(t *testing.T) {
	t.Parallel()

	assert.False(t, IsNotFoundError(nil))
	assert.True(t, IsNotFoundError(fmt.Errorf("blob with hash 0xabc not found in slot 42 sidecars (block 100)")))
	assert.False(t, IsNotFoundError(fmt.Errorf("beacon header returned status 500")))
	assert.False(t, IsNotFoundError(fmt.Errorf("rpc unavailable")))
}

// Compile-time assertion.
var _ BlobFetcher = (*BeaconClient)(nil)
