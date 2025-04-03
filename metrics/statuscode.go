package metrics

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	StatusCodeProfileName = "resiliency:type/traffic/metrics/status-code"
)

type StatusCodeProfile struct {
	Week [7][24]int
}

func NewStatusCodeProfile() (*StatusCodeProfile, *messaging.Status) {
	p, status := content.Resolve[StatusCodeProfile](StatusCodeProfileName, 1, content.Resolver)
	return &p, status
}

func (t *StatusCodeProfile) Now() int {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}
