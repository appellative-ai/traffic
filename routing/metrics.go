package routing

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
	duration   time.Duration `json:"duration"`
	statusCode int           `json:"status-code"`
}

type metrics struct {
	count     int
	status2xx int
	status4xx int
	status5xx int
	x         []float64 `json:"x"`
	weights   []float64 `json:"weights"`
}

func newMetrics() *metrics {
	m := new(metrics)
	return m
}

func (m *metrics) update(event *event) {
	m.count++
	m.x = append(m.x, float64(event.duration))

	// update status codes
	if event.statusCode >= http.StatusOK && event.statusCode < http.StatusMultipleChoices {
		m.status2xx++
		return
	}
	/*
		if event.StatusCode == http.StatusTooManyRequests {
			s.Status429++
			return
		}
		if event.StatusCode == http.StatusGatewayTimeout {
			s.Status504++
			return
		}

	*/
	if event.statusCode >= http.StatusBadRequest && event.statusCode < http.StatusInternalServerError {
		m.status4xx++
		return
	}
	if event.statusCode >= http.StatusInternalServerError {
		m.status5xx++
		return
	}

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
