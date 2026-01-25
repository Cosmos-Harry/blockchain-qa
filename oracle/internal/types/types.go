package types

import "time"

// ResponseMode defines how the oracle should respond to requests
type ResponseMode int

const (
	// OnTime - Oracle responds at the exact deadline
	OnTime ResponseMode = iota
	// Late - Oracle responds 5-15 minutes after deadline
	Late
	// Invalid - Oracle sends invalid/malformed data
	Invalid
	// NoResponse - Oracle doesn't respond (simulates downtime)
	NoResponse
)

func (m ResponseMode) String() string {
	switch m {
	case OnTime:
		return "OnTime"
	case Late:
		return "Late"
	case Invalid:
		return "Invalid"
	case NoResponse:
		return "NoResponse"
	default:
		return "Unknown"
	}
}

// PollCloseRequest represents a request to close a poll
type PollCloseRequest struct {
	PollAddress string
	Deadline    time.Time
	RequestedAt time.Time
}
