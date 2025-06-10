package representation1

import (
	"github.com/behavioral-ai/core/fmtx"
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

const (
	//Fragment            = "v1"
	RateLimitKey        = "rate-limit"
	RateBurstKey        = "rate-burst"
	OriginalPathKey     = "original-path"
	NewPathKey          = "new-path"
	IntervalKey         = "interval"
	StatusCodeThreshold = "status-code-threshold"
	PercentileThreshold = "percentile-threshold"

	defaultLimit               = rate.Limit(50)
	defaultBurst               = 10
	maxInterval                = time.Minute * 2
	defaultStatusCodeThreshold = 10 // Percentage traffic
	defaultPercentileThreshold = 95 // Milliseconds for 95 percentile
)

// Redirect - configuration
// TODO : document
type Redirect2 struct {
	Running             bool
	Limit               rate.Limit
	Burst               int
	Interval            time.Duration
	OriginalPath        string
	NewPath             string
	StatusCodeThreshold int
	PercentileThreshold int
	Codes               *StatusCodeMetrics
	Latency             *PercentileMetrics
}

func Initialize2(m map[string]string) *Redirect2 {
	r := new(Redirect2)
	r.Limit = defaultLimit
	r.Burst = defaultBurst
	r.Interval = maxInterval
	r.StatusCodeThreshold = defaultStatusCodeThreshold
	r.PercentileThreshold = defaultPercentileThreshold
	r.Codes = new(StatusCodeMetrics)
	r.Latency = new(PercentileMetrics)
	parseRedirect(r, m)
	return r
}

/*
func NewRedirect(name string) *Redirect {
	m, _ := resource.Resolve[map[string]string](name, Fragment, resource.Resolver)
	return newRedirect(m)
}
*/

func (r *Redirect2) Enabled() bool {
	return r.OriginalPath != "" && r.NewPath != ""
}

func (r *Redirect2) Failed() bool {
	return r.Latency.Failed() || r.Codes.Failed()
}

func (r *Redirect) Update(m map[string]string) {
	if m == nil {
		return
	}
}

func parseRedirect(r *Redirect2, m map[string]string) {
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

	s = m[StatusCodeThreshold]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			r.StatusCodeThreshold = i
		}
	}
	s = m[PercentileThreshold]
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			r.PercentileThreshold = i
		}
	}
}
