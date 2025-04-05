package metrics

import "github.com/behavioral-ai/collective/timeseries"

type PercentileSample struct {
	X       []float64
	Weights []float64
	Sorted  bool
}

func (p *PercentileSample) Update(event *timeseries.Event) {
	p.X = append(p.X, float64(event.Duration))
}
