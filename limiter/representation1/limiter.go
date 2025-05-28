package representation1

import (
	"github.com/behavioral-ai/core/fmtx"
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
	ThresholdKey       = "threshold"
)

const (
	offPeakDuration = time.Minute * 5
	peakDuration    = time.Minute * 2
	limit           = rate.Limit(50)
	burst           = 10
	loadSize        = 200
	threshold       = 3000 // milliseconds
)

type Limiter struct {
	Running         bool
	Enabled         bool
	Limit           rate.Limit
	Burst           int
	PeakDuration    time.Duration
	OffPeakDuration time.Duration
	LoadSize        int
	Threshold       int
}

func Initialize() *Limiter {
	return &Limiter{
		Limit:           limit,
		Burst:           burst,
		PeakDuration:    peakDuration,
		OffPeakDuration: offPeakDuration,
		LoadSize:        loadSize,
		Threshold:       threshold,
	}
}

func NewLimiter(name string) *Limiter {
	m := make(map[string]string)
	return newLimiter(name, m)
}

func newLimiter(name string, m map[string]string) *Limiter {
	l := Initialize()
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
	s = m[ThresholdKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			l.Threshold = i
		}
	}
}
