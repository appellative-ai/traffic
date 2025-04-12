package metrics

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	MetricsEvent = "event:metrics"
	ContentType  = "application/x-metrics"
)

type Metrics struct {
	Count      int
	Regression *RegressionSample
	StatusCode *StatusCodeSample
}

func NewMetrics() *Metrics {
	m := new(Metrics)
	m.Regression = new(RegressionSample)
	m.StatusCode = new(StatusCodeSample)
	return m
}

func (m *Metrics) Update(event *Event) {
	m.Count++
	m.Regression.Update(event)
	m.StatusCode.Update(event)

}

func (m *Metrics) RPS() int {
	return 0 //requestsSecond(m.Latency.Elapsed(), m.Count)
}

func requestsSecond(latency time.Duration, count int) int {
	if latency <= 0 {
		return -1
	}
	if count <= 0 {
		return 0
	}
	secs := int(latency / time.Duration(1e9))
	if secs == 0 {
		return count * 1e3
	}
	return count / secs
}

func NewMetricsMessage(s messaging.Subscription, from string, metrics Metrics) *messaging.Message {
	m := messaging.NewMessage(s.Channel, MetricsEvent)
	m.SetTo(s.From)
	m.SetFrom(from)
	m.SetContent(ContentType, metrics)
	return m
}

func MetricsContent(m *messaging.Message) (Metrics, bool) {
	if m != nil || m.Event() != MetricsEvent || m.ContentType() != ContentType {
		return Metrics{}, false
	}
	if v, ok := m.Body.(Metrics); ok {
		return v, true
	}
	return Metrics{}, false
}
