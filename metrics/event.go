package metrics

import "time"

type Event struct {
	Duration   time.Duration
	StatusCode int
}
