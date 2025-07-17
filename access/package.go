package access

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func SetTimeout(h http.Header, v time.Duration) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", TimeoutName, v))
}

func SetRateLimit(h http.Header, v float64) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", RateLimitName, v))
}

func SetRedirect(h http.Header, v int) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", RedirectName, v))
}

func SetCached(h http.Header, v bool) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", CachedName, v))
}

func RemoveThresholds(h http.Header) {
	if h == nil {
		return
	}
	h.Del(ThresholdName)
}

func Log(operators []Operator, traffic string, start time.Time, duration time.Duration, route string, req any, resp any) {
	if len(operators) == 0 {
		if agent == nil {
			operators = defaultOperators
		} else {
			operators = agent.operators
		}
	}
	e := newEvent(traffic, start, duration, route, req, resp)
	s := writeJson(operators, e)
	log.Printf("%v\n", s)
}
