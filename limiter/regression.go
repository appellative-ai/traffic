package limiter

type regressionSample struct {
	x       []float64 `json:"x"`
	y       []float64 `json:"y"`
	weights []float64 `json:"weights"`
	origin  bool      `json:"origin"`
}

func (p *regressionSample) update(event *event) {
	p.y = append(p.y, float64(event.duration))
	p.x = append(p.x, float64(event.unixMS))
}
