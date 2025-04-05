package profile

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	statusCodeName = "resiliency:type/traffic/profile/status-code"
)

type StatusCode struct {
	Week [7][24]int
}

func NewStatusCode(curr *StatusCode, resolver *content.Resolution) *StatusCode {
	p, status := content.Resolve[StatusCode](statusCodeName, 1, content.Resolver)
	if !status.OK() {
		Agent.Message(messaging.NewStatusMessage(status, ""))
		return curr
	}
	return &p
}

func (t *StatusCode) Now() int {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}
