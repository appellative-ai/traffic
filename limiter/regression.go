package limiter

type RegressionSample struct {
	X       []float64
	Y       []float64
	Weights []float64
	Origin  bool
}

func (p *RegressionSample) Update(event *event) {
	p.Y = append(p.Y, float64(event.Duration))
	p.X = append(p.Y, float64(event.Start.UnixMilli()))
}
