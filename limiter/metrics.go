package limiter

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	metricsEvent       = "event:metrics"
	contentTypeMetrics = "application/x-metrics"
)

type event struct {
	UnixMS     int64         `json:"unix-ms"`
	Duration   time.Duration `json:"duration"`
	StatusCode int           `json:"status-code"`
}

type metrics struct {
	Count      int
	Regression *RegressionSample
	StatusCode *StatusCodeSample
}

func newMetrics() *metrics {
	m := new(metrics)
	m.Regression = new(RegressionSample)
	m.StatusCode = new(StatusCodeSample)
	return m
}

func (m *metrics) Update(event *event) {
	m.Count++
	m.Regression.Update(event)
	m.StatusCode.Update(event)
}

/*
func (m *metrics) RPS() int {
	return requestsSecond(m.Duration, m.Count)
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

*/

func newMetricsMessage(metrics metrics) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelMaster, metricsEvent)
	m.SetContent(contentTypeMetrics, metrics)
	return m
}

func metricsContent(m *messaging.Message) (metrics, bool) {
	if m != nil || m.Event() != metricsEvent || m.ContentType() != contentTypeMetrics {
		return metrics{}, false
	}
	if v, ok := m.Body.(metrics); ok {
		return v, true
	}
	return metrics{}, false
}
