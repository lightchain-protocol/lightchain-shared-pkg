package types_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/lightchain/pkg/types"
)

func TestWorkerQueueName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		addr common.Address
		want string
	}{
		{
			name: "checksummed address",
			addr: common.HexToAddress("0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B"),
			want: "worker:0xab5801a7d398351b8be11c439e05c5b3259aec9b",
		},
		{
			name: "lowercase address",
			addr: common.HexToAddress("0xab5801a7d398351b8be11c439e05c5b3259aec9b"),
			want: "worker:0xab5801a7d398351b8be11c439e05c5b3259aec9b",
		},
		{
			name: "zero address",
			addr: common.Address{},
			want: "worker:0x0000000000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := types.WorkerQueueName(tt.addr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWorkerQueueName_Consistency(t *testing.T) {
	t.Parallel()

	// Same address in different formats must produce the same queue name.
	addr1 := common.HexToAddress("0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B")
	addr2 := common.HexToAddress("0xab5801a7d398351b8be11c439e05c5b3259aec9b")

	assert.Equal(t, types.WorkerQueueName(addr1), types.WorkerQueueName(addr2))
}

func TestDisputeQueueName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "disputes", types.DisputeQueueName)
}
