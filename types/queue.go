package types

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// WorkerQueueName returns the canonical Asynq queue name for a worker.
// Both dispatcher and worker must use this to ensure queue name agreement.
// Format: "worker:0x{lowercase_hex}"
func WorkerQueueName(addr common.Address) string {
	return fmt.Sprintf("worker:%s", strings.ToLower(addr.Hex()))
}

// DisputeQueueName returns the canonical Asynq queue name for dispute tasks.
const DisputeQueueName = "disputes"
