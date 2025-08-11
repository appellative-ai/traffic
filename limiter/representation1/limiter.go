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
	LoadSizeKey        = "load-size"
	ReviewLengthKey    = "review-length"
)

const (
	limit           = rate.Limit(50)
	burst           = 10
	offPeakDuration = time.Minute * 5
	peakDuration    = time.Minute * 2
	loadSize        = 200
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
	Limit           rate.Limit
	Burst           int
	PeakDuration    time.Duration
	OffPeakDuration time.Duration
	LoadSize        int
	ReviewLength    int // minutes
}

func Initialize(m map[string]string) *Limiter {
	l := &Limiter{
		Limit:           limit,
		Burst:           burst,
		PeakDuration:    peakDuration,
		OffPeakDuration: offPeakDuration,
		LoadSize:        loadSize,
	}
	parseLimiter(l, m)
	return l
}

func (l *Limiter) Update(m map[string]string) {
	parseLimiter(l, m)
}

func parseLimiter(l *Limiter, m map[string]string) {
	if l == nil || m == nil {
		return
	}
	s := m[RateLimitKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			l.Limit = rate.Limit(i)
		}
	}
	s = m[RateBurstKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			l.Burst = i
		}
	}
	s = m[PeakDurationKey]
	if s != "" {
		dur, err := fmtx.ParseDuration(s)
		if err != nil {
			//messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
			return
		}
		l.PeakDuration = dur
	}
	s = m[OffPeakDurationKey]
	if s != "" {
		dur, err := fmtx.ParseDuration(s)
		if err != nil {
			//messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
			return
		}
		l.OffPeakDuration = dur
	}

	s = m[LoadSizeKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			l.LoadSize = i
		}
	}
}
