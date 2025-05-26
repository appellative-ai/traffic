package representation1

import (
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

const (
	Fragment           = "v1"
	rateLimitKey       = "rate-limit"
	rateBurstKey       = "rate-burst"
	peakDurationKey    = "peak-duration"
	offPeakDurationKey = "off-peak-duration"
	loadSizeKey        = "load-size"
	thresholdKey       = "threshold"
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

func NewLimiter(name string) *Limiter {
	return new(nil)
}

func new(m map[string]string) *Limiter {
	value := Initialize()
	if m == nil {
		return value
	}

	s := m[rateLimitKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			value.Limit = rate.Limit(i)
		}
	}
	s = m[rateBurstKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			value.Burst = i
		}
	}
	return value
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
