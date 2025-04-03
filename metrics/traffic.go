package metrics

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	TrafficOffPeak     = "off-peak"
	TrafficPeak        = "peak"
	TrafficScaleUp     = "scale-up"
	TrafficScaleDown   = "scale-down"
	TrafficProfileName = "resiliency:type/traffic/metrics/traffic"
)

type TrafficProfile struct {
	Week [7][24]string
}

func NewTrafficProfile() (*TrafficProfile, *messaging.Status) {
	p, status := content.Resolve[TrafficProfile](TrafficProfileName, 1, content.Resolver)
	return &p, status
}

func (t *TrafficProfile) Now() string {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}
