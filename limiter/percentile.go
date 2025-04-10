package limiter

type PercentileSample struct {
	X       []float64
	Weights []float64
	Sorted  bool
}

func (p *PercentileSample) Update(event *event) {
	p.X = append(p.X, float64(event.Duration))
}
