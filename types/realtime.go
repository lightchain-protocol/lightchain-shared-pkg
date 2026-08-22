package types

// MessageType identifies the kind of realtime pub/sub message or WebSocket frame.
type MessageType string

const (
	MessageTypeChunk    MessageType = "chunk"
	MessageTypeComplete MessageType = "complete"
	MessageTypeError    MessageType = "error"
	MessageTypeMetadata MessageType = "metadata"
)

// FrameKind names the content channel a frame carries. It is orthogonal to
// MessageType: Type describes where a frame sits in the response lifecycle
// (in-flight chunk vs terminal complete), Kind describes what is inside it.
//
// A single response can interleave several kinds. Only FrameKindText
// contributes to the full-response ciphertext anchored in the blob and hashed
// into completeJob - anything else would change the hash the contract checks
// in disputeResponseMismatch and break settlement.
type FrameKind string

const (
	// FrameKindText is the assistant's visible answer. The empty string is
	// treated as text so frames from workers built before kinds existed
	// still decode correctly.
	FrameKindText FrameKind = "text"
	// FrameKindReasoning is a reasoning model's chain of thought. The worker
	// used to read Ollama's `thinking` field and discard it, so users waited
	// through reasoning they never saw.
	FrameKindReasoning FrameKind = "reasoning"
	// FrameKindArtifact is a descriptor for out-of-band content (a PDF, an
	// image, a file). The payload is JSON metadata, not the artifact itself:
	// a response blob caps at 126,972 bytes, so the bytes live off chain and
	// only their hash is committed.
	FrameKindArtifact FrameKind = "artifact"
	// FrameKindAudio is synthesized speech for voice replies.
	FrameKindAudio FrameKind = "audio"
	// FrameKindStats carries the model's own timing for the generation -
	// token counts and throughput. Ollama reports these and the worker used
	// to discard them, leaving the UI unable to show how fast a response
	// actually was.
	FrameKindStats FrameKind = "stats"
)

// Normalize maps the zero value onto text. Kind is serialized with omitempty
// so that older workers, which never set it, stay wire-compatible.
func (k FrameKind) Normalize() FrameKind {
	if k == "" {
		return FrameKindText
	}
	return k
}

// PubSubMessage is the envelope workers publish to Redis for relay fan-out.
type PubSubMessage struct {
	Type          MessageType `json:"type"`
	Kind          FrameKind   `json:"kind,omitempty"`
	JobID         JobID       `json:"jobId"`
	SessionID     SessionID   `json:"sessionId"`
	Sequence      uint32      `json:"seq"`
	TotalChunks   uint32      `json:"totalChunks"`
	Payload       []byte      `json:"payload"`
	Signature     string      `json:"signature"`
	CorrelationID string      `json:"correlationId"`
	Timestamp     int64       `json:"ts"`
}

// WSFrame is the consumer-facing WebSocket frame. It mirrors PubSubMessage.
type WSFrame struct {
	Type          MessageType `json:"type"`
	Kind          FrameKind   `json:"kind,omitempty"`
	JobID         JobID       `json:"jobId"`
	SessionID     SessionID   `json:"sessionId"`
	Sequence      uint32      `json:"seq"`
	TotalChunks   uint32      `json:"totalChunks"`
	Payload       []byte      `json:"payload"`
	Signature     string      `json:"signature"`
	CorrelationID string      `json:"correlationId"`
	Timestamp     int64       `json:"ts"`
}

// WSErrorFrame is sent when relay delivery fails in a way that would create a sequence gap.
type WSErrorFrame struct {
	Type          MessageType `json:"type"`
	Code          string      `json:"code"`
	JobID         JobID       `json:"jobId"`
	SessionID     SessionID   `json:"sessionId"`
	DroppedSeq    uint32      `json:"droppedSeq"`
	Message       string      `json:"message"`
	CorrelationID string      `json:"correlationId"`
	Timestamp     int64       `json:"ts"`
}
