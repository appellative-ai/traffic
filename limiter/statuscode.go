package limiter

import (
	"net/http"
)

type StatusCodeSample struct {
	Status2xx int `json:"status-2xx"`
	Status4xx int `json:"status-4xx"`
	Status429 int `json:"status-429"` // Too Many Requests
	Status504 int `json:"status-504"` // Gateway Timeout
	Status5xx int `json:"status-5xx"`
}

func (s *StatusCodeSample) Update(event *event) {
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
