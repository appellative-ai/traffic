package representation1

import "golang.org/x/time/rate"

const (
	defaultLimit = rate.Limit(50)
	defaultBurst = 10
)

// Redirect - configuration
type Redirect struct {
	Running      bool
	Limit        rate.Limit
	Burst        int
	OriginalPath string
	NewPath      string
	Codes        *StatusCodeThreshold
	Latency      *PercentileThreshold
}

func NewRedirect(name string) *Redirect {
	r := new(Redirect)
	r.Limit = defaultLimit
	r.Burst = defaultBurst
	r.Codes = new(StatusCodeThreshold)
	r.Latency = new(PercentileThreshold)
	return r
}

func (r *Redirect) Enabled() bool {
	return r.OriginalPath != "" && r.NewPath != ""
}

func (r *Redirect) Failed() bool {
	return r.Latency.Failed() || r.Codes.Failed()
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

func Initialize() *Redirect {
	r := new(Redirect)
	r.Limit = defaultLimit
	r.Burst = defaultBurst
	r.Codes = new(StatusCodeThreshold)
	r.Latency = new(PercentileThreshold)
	return r
}
