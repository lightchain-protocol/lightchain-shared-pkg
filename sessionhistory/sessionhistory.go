// Package sessionhistory rebuilds the conversation a chat job continues: the
// earlier jobs in its session, replayed as turns. The worker serves a job on
// that history and the disputer re-runs the job on it; both use this one
// implementation so they pick the same jobs and replay them the same way.
package sessionhistory

import (
	"context"
	"fmt"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/lightchain/pkg/blob"
	"github.com/lightchain/pkg/chain/bindings"
	"github.com/lightchain/pkg/crypto"
	"github.com/lightchain/pkg/promptenv"
	"github.com/lightchain/pkg/searchaug"
)

// LookbackBlocks is how far before a job's submit block its session's
// earlier jobs are looked for (~1.15 days at 2 s blocks). It is a protocol
// constant rather than a setting: a worker and the disputer that looked back
// different distances would rebuild different histories for the same job.
// It must stay inside the beacon chain's blob retention, or an earlier job's
// prompt could no longer be fetched.
const LookbackBlocks = 50000

// Lister lists a session's jobs from the chain.
type Lister interface {
	// SessionJobIDs returns the ids of sessionID's jobs whose JobSubmitted
	// event lies in blocks [fromBlock, toBlock], in any order.
	SessionJobIDs(ctx context.Context, sessionID, fromBlock, toBlock uint64) ([]uint64, error)
}

// PriorJobs returns the jobs whose turns jobID's history replays, oldest
// first: its session's jobs submitted before it in the LookbackBlocks up to
// submitBlock, its own submit block. A later job in the same block is left
// out by its higher id.
func PriorJobs(ctx context.Context, l Lister, sessionID, jobID, submitBlock uint64) ([]uint64, error) {
	var from uint64
	if submitBlock > LookbackBlocks {
		from = submitBlock - LookbackBlocks
	}
	ids, err := l.SessionJobIDs(ctx, sessionID, from, submitBlock)
	if err != nil {
		return nil, err
	}
	ids = slices.DeleteFunc(ids, func(id uint64) bool { return id >= jobID })
	slices.Sort(ids)
	return ids, nil
}

// Jobs reads where a job's prompt and answer are.
type Jobs interface {
	// GetJobBlobInfo returns the job's prompt and response blob hashes, zero
	// while absent, and the blocks of the transactions that carried them.
	GetJobBlobInfo(ctx context.Context, jobID uint64) (promptHash, responseHash common.Hash, submitBlock, completionBlock uint64, err error)
}

// Build replays jobIDs, in order, as the turns of a conversation: each job's
// prompt as a user turn (promptenv.Replay), then its answer, if it has one,
// as the assistant's. A self-contained job and its answer are left out: it
// was another conversation in the same session. Any fetch or decrypt failure
// fails the whole history, so no caller serves or judges a job on part of
// one.
//
// Blobs are decrypted with the current session's key, so a job id from
// another session fails to decrypt rather than leak into this one.
func Build(ctx context.Context, jobs Jobs, blobs blob.BlobFetcher, sessionKey []byte, jobIDs []uint64) ([]promptenv.Message, error) {
	var turns []promptenv.Message
	for _, jobID := range jobIDs {
		promptHash, responseHash, submitBlock, completionBlock, err := jobs.GetJobBlobInfo(ctx, jobID)
		if err != nil {
			return nil, fmt.Errorf("get blob hashes for job %d: %w", jobID, err)
		}
		if promptHash != (common.Hash{}) {
			prompt, err := fetchDecrypt(ctx, blobs, sessionKey, promptHash, submitBlock)
			if err != nil {
				return nil, fmt.Errorf("prompt of job %d: %w", jobID, err)
			}
			turn, ok := promptenv.Replay(prompt)
			if !ok {
				continue
			}
			turns = append(turns, turn)
		}
		if responseHash != (common.Hash{}) {
			response, err := fetchDecrypt(ctx, blobs, sessionKey, responseHash, completionBlock)
			if err != nil {
				return nil, fmt.Errorf("response of job %d: %w", jobID, err)
			}
			turns = append(turns, promptenv.Message{Role: "assistant", Content: searchaug.DecodeResponse(response).Answer})
		}
	}
	return turns, nil
}

func fetchDecrypt(ctx context.Context, blobs blob.BlobFetcher, sessionKey []byte, hash common.Hash, block uint64) ([]byte, error) {
	b, err := blobs.FetchBlob(ctx, hash, block)
	if err != nil {
		return nil, fmt.Errorf("fetch blob: %w", err)
	}
	plain, err := crypto.Decrypt(sessionKey, b)
	if err != nil {
		return nil, fmt.Errorf("decrypt: %w", err)
	}
	return plain, nil
}

// FilterSessionJobs is SessionJobIDs over the JobRegistry's JobSubmitted
// events, indexed by session, for a chain client to serve as its Lister.
func FilterSessionJobs(ctx context.Context, f *bindings.JobRegistryFilterer, sessionID, fromBlock, toBlock uint64) ([]uint64, error) {
	end := toBlock
	opts := &bind.FilterOpts{Context: ctx, Start: fromBlock, End: &end}
	iter, err := f.FilterJobSubmitted(opts, nil, []*big.Int{new(big.Int).SetUint64(sessionID)})
	if err != nil {
		return nil, fmt.Errorf("filter JobSubmitted session %d [%d,%d]: %w", sessionID, fromBlock, toBlock, err)
	}
	defer func() { _ = iter.Close() }()
	var ids []uint64
	for iter.Next() {
		if ev := iter.Event; ev != nil && ev.JobId != nil {
			ids = append(ids, ev.JobId.Uint64())
		}
	}
	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("iterate JobSubmitted session %d: %w", sessionID, err)
	}
	return ids, nil
}
