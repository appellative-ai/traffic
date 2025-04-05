package limiter

import (
	"net/http"
	"time"
)

const (
	contentTypeRequest = "content-type/request-profile"
)

type requestMetrics struct {
	start      time.Time
	latency    time.Duration
	statusCode int
}

type metric struct {
	start      time.Time
	latency    time.Duration
	statusCode int
}

func newMetric(start time.Time, duration time.Duration, statusCode int) metrics {
	return metrics{}
}

type observation struct {
	duration    time.Duration
	start       time.Time
	rateLimited int
	timeout     int
	metrics     []metric
}

func (o *observation) reset() {
	o.duration = 0
	o.timeout = 0
	o.rateLimited = 0
	o.metrics = nil
	o.start = time.Now().UTC()
}

func (o *observation) update(duration time.Duration, statusCode int) {
	switch statusCode {
	case http.StatusGatewayTimeout:
		o.timeout++
	case http.StatusTooManyRequests:
		o.rateLimited++
	}
}
