package metrics

type RegressionSample struct {
	X       []float64
	Y       []float64
	Weights []float64
	Origin  bool
}

func (p *RegressionSample) Update(event *Event) {
	p.Y = append(p.Y, float64(event.Duration))
}
