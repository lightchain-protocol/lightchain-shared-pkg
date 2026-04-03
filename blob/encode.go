// Package blob provides shared EIP-4844 blob encoding, decoding, and hashing utilities.
package blob

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// FieldElementSize is the size of a single BLS12-381 field element (32 bytes).
	FieldElementSize = 32
	// FieldElementDataSize is the usable data per field element (31 bytes, high byte must be zero).
	FieldElementDataSize = 31
	// BlobLengthPrefixSize is the size of the big-endian length prefix (4 bytes).
	BlobLengthPrefixSize = 4
	// MaxBlobDataSize is the total data capacity of a blob before length prefix.
	MaxBlobDataSize = params.BlobTxFieldElementsPerBlob * FieldElementDataSize
	// MaxBlobPayloadSize is the maximum user payload per blob (after reserving the length prefix).
	MaxBlobPayloadSize = MaxBlobDataSize - BlobLengthPrefixSize
)

// EncodeBlobData packs raw bytes into canonical 4844 field elements.
// Each 32-byte field element stores 31 data bytes with a zero high byte.
// The payload is prefixed with a 4-byte big-endian length to preserve exact trailing bytes.
func EncodeBlobData(data []byte) (kzg4844.Blob, error) {
	var blob kzg4844.Blob

	if len(data) > MaxBlobPayloadSize {
		return blob, fmt.Errorf("blob payload too large: got %d bytes, max %d", len(data), MaxBlobPayloadSize)
	}

	framedData := make([]byte, BlobLengthPrefixSize+len(data))
	binary.BigEndian.PutUint32(framedData[:BlobLengthPrefixSize], uint32(len(data)))
	copy(framedData[BlobLengthPrefixSize:], data)

	for i := 0; i < params.BlobTxFieldElementsPerBlob; i++ {
		srcStart := i * FieldElementDataSize
		if srcStart >= len(framedData) {
			break
		}

		dstStart := i * FieldElementSize
		copyLen := min(FieldElementDataSize, len(framedData)-srcStart)
		copy(blob[dstStart+1:dstStart+1+copyLen], framedData[srcStart:srcStart+copyLen])
	}

	return blob, nil
}

// DecodeBlobData reverses EncodeBlobData by stripping the 31-byte packing and
// using the framed payload length to restore the exact original bytes.
func DecodeBlobData(blobData []byte) ([]byte, error) {
	if len(blobData)%FieldElementSize != 0 {
		return nil, fmt.Errorf("blob sidecar size %d is not a multiple of %d", len(blobData), FieldElementSize)
	}
	if len(blobData) > params.BlobTxFieldElementsPerBlob*FieldElementSize {
		return nil, fmt.Errorf("blob sidecar too large: got %d bytes", len(blobData))
	}

	unpacked := make([]byte, 0, len(blobData)/FieldElementSize*FieldElementDataSize)
	for i := 0; i < len(blobData); i += FieldElementSize {
		if blobData[i] != 0 {
			return nil, fmt.Errorf("blob sidecar field element %d is not canonical", i/FieldElementSize)
		}
		unpacked = append(unpacked, blobData[i+1:i+FieldElementSize]...)
	}

	if len(unpacked) < BlobLengthPrefixSize {
		return nil, fmt.Errorf("blob sidecar missing payload length prefix")
	}

	payloadLen := int(binary.BigEndian.Uint32(unpacked[:BlobLengthPrefixSize]))
	if payloadLen < 0 || payloadLen > MaxBlobPayloadSize {
		return nil, fmt.Errorf("blob payload length %d exceeds max %d", payloadLen, MaxBlobPayloadSize)
	}
	if BlobLengthPrefixSize+payloadLen > len(unpacked) {
		return nil, fmt.Errorf("blob payload length %d exceeds unpacked size %d", payloadLen, len(unpacked)-BlobLengthPrefixSize)
	}

	payload := make([]byte, payloadLen)
	copy(payload, unpacked[BlobLengthPrefixSize:BlobLengthPrefixSize+payloadLen])
	return payload, nil
}
