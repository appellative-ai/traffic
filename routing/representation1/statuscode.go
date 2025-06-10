package representation1

// StatusCodeMetrics - redirect status code thresholds
type StatusCodeMetrics struct {
	MaxFailures int
	Failures    int
	Status2xx   int
	Status4xx   int
	Status5xx   int
}

// Failed - threshold has been exceeded
func (s *StatusCodeMetrics) Failed() bool {
	return s.Failures >= s.MaxFailures
}

// AddFailure - add a failure
func (s *StatusCodeMetrics) AddFailure() {
	s.Failures++
}
