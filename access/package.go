package access

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"log"
	"net/http"
	"time"
)

const (
	ContentTypeOperators = "application/x-log-operators"
	DefaultRoute         = "host"
	EgressTraffic        = "egress"
	IngressTraffic       = "ingress"

	RequestIdName = "x-request-id"
	ThresholdName = "x-threshold"
	RateLimitName = "rate-limit"
	TimeoutName   = "timeout"
	RedirectName  = "redirect"
	CachedName    = "cached"
)

// Operator - configuration of logging entries
type Operator struct {
	Name  string
	Value string
}

// Request - request interface for non HTTP traffic
type Request interface {
	Url() string
	Header() http.Header
	Method() string
	Protocol() string
}

// Response - response interface for non HTTP traffic
type Response interface {
	StatusCode() int
	Header() http.Header
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

func LogEgress(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	var r *http.Response
	if duration > 0 {
		r = buildResponse(resp)
		SetTimeout(r.Header, timeout)
	}
	Log(nil, EgressTraffic, start, duration, route, req, r)
}

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

func NewOperatorsMessage(ops []Operator) *messaging.Message {
	return messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent).SetContent(ContentTypeOperators, ops)
}

func OperatorsContent(m *messaging.Message) ([]Operator, *messaging.Status) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeOperators) {
		return nil, messaging.NewStatus(messaging.StatusInvalidContent, "")
	}
	return messaging.New[[]Operator](m.Content)
}
