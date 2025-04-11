package limiter

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	metricsEvent       = "event:metrics"
	contentTypeMetrics = "application/x-metrics"
)

type event struct {
	Internal   bool          `json:"internal"`
	UnixMS     int64         `json:"unix-ms"`
	Duration   time.Duration `json:"duration"`
	StatusCode int           `json:"status-code"`
}

type metrics struct {
	Count      int
	Status429  int
	Regression *RegressionSample
}

func newMetrics() *metrics {
	m := new(metrics)
	m.Regression = new(RegressionSample)
	return m
}

func (m *metrics) Update(event *event) {
	m.Count++
	if event.Internal && event.StatusCode == http.StatusTooManyRequests {
		m.Status429++
	}
	m.Regression.Update(event)

}

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
