package config

type StatusCode struct {
	MaxFailures int
	Failures    int
	Status2xx   int
	Status4xx   int
	Status5xx   int
}

func (s *StatusCode) Failed() bool {
	return s.Failures >= s.MaxFailures
}

func (s *StatusCode) AddFailure() {
	s.Failures++
}

type Percentile struct {
	MaxFailures int
	Failures    int
	Score       int
	Latency     int // milliseconds
}

func (p *Percentile) Failed() bool {
	return p.Failures >= p.MaxFailures
}

func (p *Percentile) AddFailure() {
	p.Failures++
}

type Redirect struct {
	OriginalPath string
	NewPath      string
	Codes        *StatusCode
	Latency      *Percentile
}

func NewRedirect() *Redirect {
	r := new(Redirect)
	r.Codes = new(StatusCode)
	r.Latency = new(Percentile)
	return r
}

func (r *Redirect) Enabled() bool {
	return r.OriginalPath != "" && r.NewPath != ""
}

func (r *Redirect) Failed() bool {
	return r.Latency.Failed() || r.Codes.Failed()
}
