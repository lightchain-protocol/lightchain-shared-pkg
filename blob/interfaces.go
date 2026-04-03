package blob

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// BlobFetcher fetches EIP-4844 blob data given a versioned hash and block number.
type BlobFetcher interface {
	FetchBlob(ctx context.Context, versionedHash common.Hash, blockNumber uint64) ([]byte, error)
}

// BlobSubmitter submits EIP-4844 blob transactions and returns the versioned hashes.
type BlobSubmitter interface {
	SubmitBlobTx(ctx context.Context, data []byte) ([][32]byte, error)
}
