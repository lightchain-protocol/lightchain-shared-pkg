package blob

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeBlobData_ProducesCanonicalBlob(t *testing.T) {
	t.Parallel()

	input := make([]byte, 1024)
	for i := range input {
		input[i] = byte(i % 251)
	}

	blob, err := EncodeBlobData(input)
	require.NoError(t, err)

	for i := 0; i < params.BlobTxFieldElementsPerBlob; i++ {
		assert.Equal(t, byte(0), blob[i*FieldElementSize], "field element %d must be canonical", i)
	}

	decoded, err := DecodeBlobData(blob[:])
	require.NoError(t, err)
	assert.Equal(t, input, decoded)

	commitment, err := kzg4844.BlobToCommitment(&blob)
	require.NoError(t, err)

	_, err = kzg4844.ComputeBlobProof(&blob, commitment)
	require.NoError(t, err)
}

func TestEncodeBlobData_RejectsOversizePayload(t *testing.T) {
	t.Parallel()

	_, err := EncodeBlobData(make([]byte, MaxBlobPayloadSize+1))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "blob payload too large")
}

func TestEncodeBlobData_RoundTripsTrailingZeros(t *testing.T) {
	t.Parallel()

	input := append([]byte("ciphertext-with-trailing-zeros"), 0x00, 0x00, 0x00)

	blob, err := EncodeBlobData(input)
	require.NoError(t, err)

	decoded, err := DecodeBlobData(blob[:])
	require.NoError(t, err)
	assert.Equal(t, input, decoded)
}

func TestEncodeBlobData_EmptyPayload(t *testing.T) {
	t.Parallel()

	blob, err := EncodeBlobData([]byte{})
	require.NoError(t, err)

	decoded, err := DecodeBlobData(blob[:])
	require.NoError(t, err)
	assert.Equal(t, []byte{}, decoded)
}

func TestEncodeBlobData_MaxPayload(t *testing.T) {
	t.Parallel()

	input := make([]byte, MaxBlobPayloadSize)
	for i := range input {
		input[i] = byte(i % 256)
	}

	blob, err := EncodeBlobData(input)
	require.NoError(t, err)

	decoded, err := DecodeBlobData(blob[:])
	require.NoError(t, err)
	assert.Equal(t, input, decoded)
}

func TestDecodeBlobData_NonCanonicalFieldElement(t *testing.T) {
	t.Parallel()

	blobData := make([]byte, FieldElementSize)
	blobData[0] = 0x01 // non-canonical high byte

	_, err := DecodeBlobData(blobData)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not canonical")
}

func TestDecodeBlobData_InvalidSize(t *testing.T) {
	t.Parallel()

	_, err := DecodeBlobData(make([]byte, 33)) // not a multiple of 32
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not a multiple of")
}
