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
	internal   bool          `json:"internal"`
	unixMS     int64         `json:"unix-ms"`
	duration   time.Duration `json:"duration"`
	statusCode int           `json:"status-code"`
}

type metrics struct {
	count      int
	status429  int
	regression *regressionSample
}

func newMetrics() *metrics {
	m := new(metrics)
	m.regression = new(regressionSample)
	return m
}

func (m *metrics) update(event *event) {
	m.count++
	if event.internal && event.statusCode == http.StatusTooManyRequests {
		m.status429++
	}
	m.regression.update(event)

}

func newMetricsMessage(metrics metrics) *messaging.Message {
	return messaging.NewMessage(messaging.ChannelMaster, metricsEvent).SetContent(contentTypeMetrics, metrics)
}

func metricsContent(m *messaging.Message) (metrics, *messaging.Status) {
	if !messaging.ValidContent(m, metricsEvent, contentTypeMetrics) {
		return metrics{}, messaging.NewStatus(messaging.StatusInvalidContent, "")
	}
	return messaging.New[metrics](m.Content)
}
