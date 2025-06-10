package representation1

// PercentileMetrics - redirect configured latency threshold
type PercentileMetrics struct {
	MaxFailures int
	Failures    int
	Score       int
	Latency     int // milliseconds
}

// Failed - latency threshold exceeded
func (p *PercentileMetrics) Failed() bool {
	return p.Failures >= p.MaxFailures
}

// AddFailure - add a failure
func (p *PercentileMetrics) AddFailure() {
	p.Failures++
}
