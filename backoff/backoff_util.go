package backoff

import "time"

type ExponentialBackoff interface {
	NextDelay() time.Duration
	Reset()
	CurrentDelay() time.Duration
}

// Implement the ExponentialBackoff interface
