package metrics

import "github.com/behavioral-ai/collective/timeseries"

type RegressionSample struct {
	X       []float64
	Y       []float64
	Weights []float64
	Origin  bool
}

func (p *RegressionSample) Update(event *timeseries.Event) {
	p.Y = append(p.Y, float64(event.Duration))
}
