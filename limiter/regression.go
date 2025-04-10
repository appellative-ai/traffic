package limiter

type RegressionSample struct {
	X       []float64 `json:"x"`
	Y       []float64 `json:"y"`
	Weights []float64 `json:"weights"`
	Origin  bool      `json:"origin"`
}

func (p *RegressionSample) Update(event *event) {
	p.Y = append(p.Y, float64(event.Duration))
	p.X = append(p.Y, float64(event.UnixMS))
}
