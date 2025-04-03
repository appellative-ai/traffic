package metrics

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	PercentileProfileName = "resiliency:type/traffic/metrics/percentile"
)

type PercentileProfile struct {
	Week [7][24]float64
}

func NewPercentileProfile() (*PercentileProfile, *messaging.Status) {
	p, status := content.Resolve[PercentileProfile](PercentileProfileName, 1, content.Resolver)
	return &p, status
}

func (t *PercentileProfile) Now() float64 {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}
