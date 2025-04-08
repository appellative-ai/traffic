package metrics

import (
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	Event       = "event:metrics"
	ContentType = "application/x-metrics"
)

type Metrics struct {
	Count      int
	Latency    *LatencySample
	Regression *RegressionSample
	Percentile *PercentileSample
	StatusCode *StatusCodeSample
}

func NewMetrics() *Metrics {
	m := new(Metrics)
	m.Latency = new(LatencySample)
	m.Regression = new(RegressionSample)
	m.Percentile = new(PercentileSample)
	m.StatusCode = new(StatusCodeSample)
	return m
}

func (m *Metrics) Update(event *timeseries.Event) {
	m.Count++
	m.Percentile.Update(event)
	m.Regression.Update(event)
	m.StatusCode.Update(event)
	m.Latency.Update(event)
}

func (m *Metrics) RPS() int {
	return requestsSecond(m.Latency.Elapsed(), m.Count)
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

func NewMetricsMessage(metrics Metrics) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, Event)
	m.SetContent(ContentType, metrics)
	return m
}

func MetricsContent(m *messaging.Message) (Metrics, bool) {
	if m.Event() != Event || m.ContentType() != ContentType {
		return Metrics{}, false
	}
	if v, ok := m.Body.(Metrics); ok {
		return v, true
	}
	return Metrics{}, false
}
