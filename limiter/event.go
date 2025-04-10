package limiter

import "time"

type event struct {
	Duration   time.Duration `json:"duration"`
	StatusCode int           `json:"status-code"`
}

//Path       string        `json:"path"` // uri path
//	Start      time.Time     `json:"start-ts"`
