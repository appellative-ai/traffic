package messaging

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	MetricsEvent       = "event:metrics"
	ContentTypeMetrics = "application/metrics"
)

type StatusCodeDistribution struct {
	Status504 int // Gateway Timeout
	Status429 int // Too Many Requests
	Status5xx int
	Status2xx int
	Status4xx int
}

type Metrics struct {
	Latency time.Duration // For the 95th percentile
	Alpha   float64
	Beta    float64
	Code    StatusCodeDistribution
}

func NewMetricsMessage(metrics Metrics) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, MetricsEvent)
	m.SetContent(ContentTypeMetrics, metrics)
	return m
}

func MetricsContent(m *messaging.Message) *Metrics {
	if m.Event() != MetricsEvent || m.ContentType() != ContentTypeMetrics {
		return nil
	}
	if v, ok := m.Body.(Metrics); ok {
		return &v
	}
	return nil
}
