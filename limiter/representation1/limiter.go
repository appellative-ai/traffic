package representation1

import (
	"github.com/appellative-ai/core/fmtx"
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

const (
	Fragment           = "v1"
	RateLimitKey       = "rate-limit"
	RateBurstKey       = "rate-burst"
	PeakDurationKey    = "peak-duration"
	OffPeakDurationKey = "off-peak-duration"
	WindowSizeKey      = "window-size"
	ReviewDurationKey  = "review-duration"
)

const (
	limit           = rate.Limit(50)
	burst           = 10
	offPeakDuration = time.Minute * 5
	peakDuration    = time.Minute * 2
	windowSize      = 200
)

// Limiter - values used by the agent for rate limiting.
// //         Peak,off-peak and load size adjust how often events are dequeued and sent to the master
//
//	to be analyzed via linear regression.
//
//	The load size is the threshold between peak and off-peak durations.
//
//	Limit and burst are the starting values for rate-limiting. These get changed based on regression
//	analysis of the events.
type Limiter struct {
	Limit           rate.Limit    `json:"limit"`
	Burst           int           `json:"burst"`
	WindowSize      int           `json:"window-size"` // Number of events to track before analysis
	PeakDuration    time.Duration `json:"peak-duration"`
	OffPeakDuration time.Duration `json:"off-peak-duration"`
	ReviewDuration  time.Duration `json:"review-duration"`
}

func Initialize(m map[string]string) *Limiter {
	l := &Limiter{
		Limit:           limit,
		Burst:           burst,
		PeakDuration:    peakDuration,
		OffPeakDuration: offPeakDuration,
		WindowSize:      windowSize,
	}
	parseLimiter(l, m)
	return l
}

func (l *Limiter) Update(m map[string]string) bool {
	return parseLimiter(l, m)
}

func parseLimiter(l *Limiter, m map[string]string) (changed bool) {
	if l == nil || m == nil {
		return
	}
	s := m[RateLimitKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil && i > 0 {
			if l.Limit != rate.Limit(i) {
				l.Limit = rate.Limit(i)
				changed = true
			}
		}
	}
	s = m[RateBurstKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil && i > 0 {
			if l.Burst != i {
				l.Burst = i
				changed = true
			}
		}
	}
	s = m[PeakDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if l.PeakDuration != dur {
				l.PeakDuration = dur
				changed = true
			}
		}
	}
	s = m[OffPeakDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if l.OffPeakDuration != dur {
				l.OffPeakDuration = dur
				changed = true
			}
		}
	}
	s = m[WindowSizeKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil && i > 0 {
			if l.WindowSize != i {
				l.WindowSize = i
				changed = true
			}
		}
	}
	s = m[ReviewDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if l.ReviewDuration != dur {
				l.ReviewDuration = dur
				changed = true
			}
		}
	}
	return
}
