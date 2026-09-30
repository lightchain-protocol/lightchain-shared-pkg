package sessionhistory

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lightchain/pkg/chain/bindings"
	"github.com/lightchain/pkg/crypto"
	"github.com/lightchain/pkg/promptenv"
)

// listerFunc lists a session's jobs from a fake chain's JobSubmitted events.
type listerFunc func(sessionID, fromBlock, toBlock uint64) ([]uint64, error)

func (f listerFunc) SessionJobIDs(_ context.Context, sessionID, fromBlock, toBlock uint64) ([]uint64, error) {
	return f(sessionID, fromBlock, toBlock)
}

// A session's earlier jobs are its JobSubmitted events in the 50000 blocks up
// to the job's own, however many job ids behind it they sit, oldest first.
// The job itself and any later job in its block are not earlier jobs.
func TestPriorJobs_ListsTheSessionsEarlierJobsInTheLookBackWindow(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name                  string
		submitBlock, wantFrom uint64
	}{
		{name: "deep in the chain", submitBlock: 120000, wantFrom: 70000},
		{name: "near genesis", submitBlock: 1200, wantFrom: 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var query []uint64
			events := listerFunc(func(sessionID, fromBlock, toBlock uint64) ([]uint64, error) {
				query = []uint64{sessionID, fromBlock, toBlock}
				return []uint64{140, 201, 20, 200, 100}, nil
			})

			ids, err := PriorJobs(context.Background(), events, 7, 200, tc.submitBlock)

			require.NoError(t, err)
			assert.Equal(t, []uint64{7, tc.wantFrom, tc.submitBlock}, query)
			assert.Equal(t, []uint64{20, 100, 140}, ids)
		})
	}
}

// jobSubmittedLogs is a node's log filter holding JobSubmitted logs. It
// records the query it is asked, and leaves the matching to the real node.
type jobSubmittedLogs struct {
	query ethereum.FilterQuery
	logs  []types.Log
}

func (l *jobSubmittedLogs) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	l.query = q
	return l.logs, nil
}

func (l *jobSubmittedLogs) SubscribeFilterLogs(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error) {
	return nil, errors.New("not used")
}

func jobSubmittedLog(t *testing.T, jobID, sessionID uint64) types.Log {
	t.Helper()
	parsed, err := bindings.JobRegistryMetaData.GetAbi()
	require.NoError(t, err)
	return types.Log{
		Topics: []common.Hash{
			parsed.Events["JobSubmitted"].ID,
			common.BigToHash(new(big.Int).SetUint64(jobID)),
			common.BigToHash(new(big.Int).SetUint64(sessionID)),
		},
		Data: common.LeftPadBytes(common.HexToAddress("0x0b").Bytes(), 32),
	}
}

// The node is asked for the session's JobSubmitted logs in the window, and
// the job ids come back from them.
func TestFilterSessionJobs_AsksTheNodeForTheSessionsJobSubmittedLogsInTheWindow(t *testing.T) {
	t.Parallel()
	registry := common.HexToAddress("0x1e")
	node := &jobSubmittedLogs{logs: []types.Log{jobSubmittedLog(t, 20, 7), jobSubmittedLog(t, 100, 7)}}
	f, err := bindings.NewJobRegistryFilterer(registry, node)
	require.NoError(t, err)

	ids, err := FilterSessionJobs(context.Background(), f, 7, 70000, 120000)

	require.NoError(t, err)
	assert.Equal(t, []uint64{20, 100}, ids)
	assert.Equal(t, []common.Address{registry}, node.query.Addresses)
	assert.Equal(t, big.NewInt(70000), node.query.FromBlock)
	assert.Equal(t, big.NewInt(120000), node.query.ToBlock)
	require.Len(t, node.query.Topics, 3)
	assert.Empty(t, node.query.Topics[1], "any job id")
	assert.Equal(t, []common.Hash{common.BigToHash(big.NewInt(7))}, node.query.Topics[2], "only the session's jobs")
}

// chainJobs is a fake chain whose jobs' prompts and answers sit in blobs
// encrypted with the session key.
type chainJobs struct {
	t          *testing.T
	sessionKey []byte
	blobs      map[common.Hash][]byte
	jobs       map[uint64][2]common.Hash
}

func (c *chainJobs) add(id uint64, prompt, answer string) {
	p, a := common.Hash{0xaa, byte(id)}, common.Hash{0xbb, byte(id)}
	c.jobs[id] = [2]common.Hash{p, a}
	for h, plain := range map[common.Hash]string{p: prompt, a: answer} {
		enc, err := crypto.Encrypt(c.sessionKey, []byte(plain))
		require.NoError(c.t, err)
		c.blobs[h] = enc
	}
}

func (c *chainJobs) GetJobBlobInfo(_ context.Context, jobID uint64) (prompt, response common.Hash, submitBlock, completionBlock uint64, err error) {
	h, ok := c.jobs[jobID]
	if !ok {
		return prompt, response, 0, 0, fmt.Errorf("no job %d", jobID)
	}
	return h[0], h[1], 10 * jobID, 10*jobID + 1, nil
}

func (c *chainJobs) FetchBlob(_ context.Context, hash common.Hash, _ uint64) ([]byte, error) {
	b, ok := c.blobs[hash]
	if !ok {
		return nil, fmt.Errorf("no blob %s", hash)
	}
	return b, nil
}

// A multimodal prompt replays as its text with its images, not as the raw
// envelope; the answer follows it.
func TestBuild_ReplaysAMultimodalPromptWithItsImages(t *testing.T) {
	t.Parallel()
	key := bytes.Repeat([]byte{7}, 32)
	c := &chainJobs{t: t, sessionKey: key, blobs: map[common.Hash][]byte{}, jobs: map[uint64][2]common.Hash{}}
	c.add(20, "plain question", "plain answer")
	c.add(60, `{"v":1,"text":"what is in this picture?","images":["aW1n"]}`, "a cat")

	turns, err := Build(context.Background(), c, c, key, []uint64{20, 60}, 1000)

	require.NoError(t, err)
	assert.Equal(t, []promptenv.Message{
		{Role: "user", Content: "plain question"},
		{Role: "assistant", Content: "plain answer"},
		{Role: "user", Content: "what is in this picture?", Images: []string{"aW1n"}},
		{Role: "assistant", Content: "a cat"},
	}, turns)
}

// An answer mined after the current job's submit block is left out, though
// its prompt is replayed: whether the worker saw it on chain depends on when
// it served the job, so the disputer could not know. An answer mined in that
// very block is on chain for both.
func TestBuild_LeavesOutAnAnswerMinedAfterTheJobWasSubmitted(t *testing.T) {
	t.Parallel()
	key := bytes.Repeat([]byte{7}, 32)
	c := &chainJobs{t: t, sessionKey: key, blobs: map[common.Hash][]byte{}, jobs: map[uint64][2]common.Hash{}}
	c.add(20, "first question", "first answer")   // answer mined in block 201
	c.add(60, "second question", "second answer") // answer mined in block 601

	turns, err := Build(context.Background(), c, c, key, []uint64{20, 60}, 600)

	require.NoError(t, err)
	assert.Equal(t, []promptenv.Message{
		{Role: "user", Content: "first question"},
		{Role: "assistant", Content: "first answer"},
		{Role: "user", Content: "second question"},
	}, turns)

	turns, err = Build(context.Background(), c, c, key, []uint64{20, 60}, 601)

	require.NoError(t, err)
	assert.Len(t, turns, 4, "an answer mined in the job's own submit block is replayed")
}
