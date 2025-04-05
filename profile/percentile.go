package profile

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	percentileName = "resiliency:type/traffic/profile/percentile"
)

type Percentile struct {
	Week [7][24]float64
}

func NewPercentile(curr *Percentile, resolver *content.Resolution) *Percentile {
	p, status := content.Resolve[Percentile](percentileName, 1, resolver)
	if !status.OK() {
		agent.Message(messaging.NewStatusMessage(status, ""))
		return curr
	}
	return &p
}

func (t *Percentile) Now() float64 {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}
