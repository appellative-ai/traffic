package metrics

import (
	"github.com/behavioral-ai/collective/timeseries"
	"net/http"
)

type StatusCodeSample struct {
	Status2xx int
	Status4xx int
	Status429 int // Too Many Requests
	Status504 int // Gateway Timeout
	Status5xx int
}

func (s *StatusCodeSample) Update(event *timeseries.Event) {
	if event.StatusCode >= http.StatusOK && event.StatusCode < http.StatusMultipleChoices {
		s.Status2xx++
		return
	}
	if event.StatusCode == http.StatusTooManyRequests {
		s.Status429++
		return
	}
	if event.StatusCode == http.StatusGatewayTimeout {
		s.Status504++
		return
	}
	if event.StatusCode >= http.StatusBadRequest && event.StatusCode < http.StatusInternalServerError {
		s.Status4xx++
		return
	}
	if event.StatusCode >= http.StatusInternalServerError {
		s.Status5xx++
		return
	}
}
