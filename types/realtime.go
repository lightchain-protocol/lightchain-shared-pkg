package types

// MessageType identifies the kind of realtime pub/sub message or WebSocket frame.
type MessageType string

const (
	MessageTypeChunk    MessageType = "chunk"
	MessageTypeComplete MessageType = "complete"
	MessageTypeError    MessageType = "error"
)

// PubSubMessage is the envelope workers publish to Redis for relay fan-out.
type PubSubMessage struct {
	Type          MessageType `json:"type"`
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
