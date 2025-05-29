package representation1

import (
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/fmtx"
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

const (
	Fragment        = "v1"
	RateLimitKey    = "rate-limit"
	RateBurstKey    = "rate-burst"
	OriginalPathKey = "original-path"
	NewPathKey      = "new-path"
	IntervalKey     = "interval-key"

	defaultLimit = rate.Limit(50)
	defaultBurst = 10
	maxInterval  = time.Minute * 2
)

// Redirect - configuration
// TODO : Add thresholds for status code and latency failures
type Redirect struct {
	Running      bool
	Limit        rate.Limit
	Burst        int
	Interval     time.Duration
	OriginalPath string
	NewPath      string
	Codes        *StatusCodeThreshold
	Latency      *PercentileThreshold
}

func Initialize() *Redirect {
	r := new(Redirect)
	r.Limit = defaultLimit
	r.Burst = defaultBurst
	r.Interval = maxInterval
	r.Codes = new(StatusCodeThreshold)
	r.Latency = new(PercentileThreshold)
	return r
}

func NewRedirect(name string) *Redirect {
	m, _ := resource.Resolve[map[string]string](name, Fragment, resource.Resolver)
	return newRedirect(m)
}

func newRedirect(m map[string]string) *Redirect {
	r := new(Redirect)
	parseRedirect(r, m)
	return r
}

func (r *Redirect) Enabled() bool {
	return r.OriginalPath != "" && r.NewPath != ""
}

func (r *Redirect) Failed() bool {
	return r.Latency.Failed() || r.Codes.Failed()
}

func (r *Redirect) Update(m map[string]string) {
	if m == nil {
		return
	}
}

func parseRedirect(r *Redirect, m map[string]string) {
	if r == nil || m == nil {
		return
	}
	s := m[RateLimitKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			r.Limit = rate.Limit(i)
		}
	}
	s = m[RateBurstKey]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			r.Burst = i
		}
	}
	s = m[OriginalPathKey]
	if s != "" {
		r.OriginalPath = s
	}
	s = m[NewPathKey]
	if s != "" {
		r.NewPath = s
	}

	s = m[IntervalKey]
	if s != "" {
		dur, err := fmtx.ParseDuration(s)
		if err != nil {
			//messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
			return
		}
		r.Interval = dur
	}
}

// StatusCodeThreshold - redirect status code thresholds
type StatusCodeThreshold struct {
	MaxFailures int
	Failures    int
	Status2xx   int
	Status4xx   int
	Status5xx   int
}

// Failed - threshold has been exceeded
func (s *StatusCodeThreshold) Failed() bool {
	return s.Failures >= s.MaxFailures
}

// AddFailure - add a failure
func (s *StatusCodeThreshold) AddFailure() {
	s.Failures++
}

// PercentileThreshold - redirect configured latency threshold
type PercentileThreshold struct {
	MaxFailures int
	Failures    int
	Score       int
	Latency     int // milliseconds
}

// Failed - latency threshold exceeded
func (p *PercentileThreshold) Failed() bool {
	return p.Failures >= p.MaxFailures
}

// AddFailure - add a failure
func (p *PercentileThreshold) AddFailure() {
	p.Failures++
}
