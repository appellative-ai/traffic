package profile

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	TrafficOffPeak   = "off-peak"
	TrafficPeak      = "peak"
	TrafficScaleUp   = "scale-up"
	TrafficScaleDown = "scale-down"
	trafficName      = "resiliency:type/traffic/profile/traffic"
)

type Traffic struct {
	Week [7][24]string
}

func NewTraffic(curr *Traffic, resolver *content.Resolution) *Traffic {
	p, status := content.Resolve[Traffic](trafficName, 1, resolver)
	if !status.OK() {
		agent.Message(messaging.NewStatusMessage(status, ""))
		if curr == nil {
			curr = &Traffic{}
		}
		return curr
	}
	return &p
}

func (t *Traffic) Now() string {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}
