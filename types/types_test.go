package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJobState(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		state JobState
		value uint8
		str   string
	}{
		{name: "Submitted", state: JobStateSubmitted, value: 0, str: "Submitted"},
		{name: "Acknowledged", state: JobStateAcknowledged, value: 1, str: "Acknowledged"},
		{name: "Completed", state: JobStateCompleted, value: 2, str: "Completed"},
		{name: "TimedOut", state: JobStateTimedOut, value: 3, str: "TimedOut"},
		{name: "Disputed", state: JobStateDisputed, value: 4, str: "Disputed"},
		{name: "Resolved", state: JobStateResolved, value: 5, str: "Resolved"},
		{name: "Released", state: JobStateReleased, value: 6, str: "Released"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.value, uint8(tt.state))
			assert.Equal(t, tt.str, tt.state.String())
		})
	}
}

func TestJobState_unknown(t *testing.T) {
	t.Parallel()
	unknown := JobState(99)
	assert.Equal(t, "JobState(99)", unknown.String())
}

func TestSessionStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status SessionStatus
		value  uint8
		str    string
	}{
		{name: "Active", status: SessionStatusActive, value: 0, str: "Active"},
		{name: "Reassigning", status: SessionStatusReassigning, value: 1, str: "Reassigning"},
		{name: "Closed", status: SessionStatusClosed, value: 2, str: "Closed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.value, uint8(tt.status))
			assert.Equal(t, tt.str, tt.status.String())
		})
	}
}

func TestSessionStatus_unknown(t *testing.T) {
	t.Parallel()
	unknown := SessionStatus(99)
	assert.Equal(t, "SessionStatus(99)", unknown.String())
}

func TestWorkerStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status WorkerStatus
		value  uint8
		str    string
	}{
		{name: "Active", status: WorkerStatusActive, value: 0, str: "Active"},
		{name: "Inactive", status: WorkerStatusInactive, value: 1, str: "Inactive"},
		{name: "Suspended", status: WorkerStatusSuspended, value: 2, str: "Suspended"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.value, uint8(tt.status))
			assert.Equal(t, tt.str, tt.status.String())
		})
	}
}

func TestWorkerStatus_unknown(t *testing.T) {
	t.Parallel()
	unknown := WorkerStatus(99)
	assert.Equal(t, "WorkerStatus(99)", unknown.String())
}

func TestWorkerAddr_ToHex(t *testing.T) {
	t.Parallel()

	addr := WorkerAddr{0xde, 0xad, 0xbe, 0xef}
	hex := addr.ToHex()
	require.Contains(t, hex, "0x")
	require.Contains(t, hex, "deadbeef")
}

func TestZeroValues(t *testing.T) {
	t.Parallel()

	t.Run("JobState zero is Submitted", func(t *testing.T) {
		t.Parallel()
		var s JobState
		assert.Equal(t, JobStateSubmitted, s)
	})

	t.Run("SessionStatus zero is Active", func(t *testing.T) {
		t.Parallel()
		var s SessionStatus
		assert.Equal(t, SessionStatusActive, s)
	})

	t.Run("WorkerStatus zero is Active", func(t *testing.T) {
		t.Parallel()
		var s WorkerStatus
		assert.Equal(t, WorkerStatusActive, s)
	})

	t.Run("Job zero value has Submitted state", func(t *testing.T) {
		t.Parallel()
		var j Job
		assert.Equal(t, JobStateSubmitted, j.State)
	})

	t.Run("Session zero value has Active status", func(t *testing.T) {
		t.Parallel()
		var s Session
		assert.Equal(t, SessionStatusActive, s.Status)
	})
}
