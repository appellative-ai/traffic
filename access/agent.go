package access

import (
	"encoding/json"
	"errors"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"log"
	"net/http"
	"time"
)

//XRateBurst      = "x-rate-burst"

const (
	NamespaceName  = "core:common:agent/log/access/http"
	Route          = "host"
	EgressTraffic  = "egress"
	IngressTraffic = "ingress"

	failsafeUri     = "https://invalid-uri.com"
	RequestIdName   = "x-request-id"
	RateLimitName   = "rate-limit"
	TimeoutName     = "timeout"
	RedirectName    = "redirect"
	CachedName      = "cached"
	ContentEncoding = "Content-Encoding"

	ThresholdRequest       = "x-threshold-request"
	ThresholdResponse      = "x-threshold-response"
	ThresholdCacheName     = "cache"
	ThresholdRateLimitName = "rate-limit"
	ThresholdTimeoutName   = "timeout"
	ThresholdRedirectName  = "redirect"
)

type LogAgent interface {
	messaging.Agent
	SetOrigin(region, zone, subZone, host, instanceId string)
	ConfigureOperators(read func() ([]byte, error)) error
	Log(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold)
}

var (
	Agent LogAgent
	agent *agentT
)

func init() {
	log.SetFlags(0)
	Agent = newLogAgent()
}

type agentT struct {
	name             string
	origin           originT
	originSet        bool
	defaultOperators []Operator
}

func newLogAgent() LogAgent {
	agent = newAgent()
	return agent
}

func newAgent() *agentT {
	a := new(agentT)
	a.name = NamespaceName
	a.defaultOperators, _ = createOperators([]string{TrafficOperator, StartTimeOperator,
		DurationOperator, RouteOperator,
		RequestMethodOperator, RequestUrlOperator,
		ResponseStatusCodeOperator, ResponseCachedOperator,
		ResponseContentEncodingOperator, ResponseBytesReceivedOperator,
		TimeoutDurationOperator, RateLimitOperator,
		RedirectOperator,
	})
	return a
}

func (a *agentT) Name() string { return a.name }
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Name == messaging.StartupEvent {
		//t.run()
		return
	}

}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		resp, err = next(r)
		a.Log(IngressTraffic, start, time.Since(start), Route, r, resp, newThreshold(resp))
		return
	}
}

// SetOrigin -
func (a *agentT) SetOrigin(region, zone, subZone, host, instanceId string) {
	a.origin.Region = region
	a.origin.Zone = zone
	a.origin.SubZone = subZone
	a.origin.Host = host
	a.origin.InstanceId = instanceId
	a.originSet = true
}

// ConfigureOperators - load operators from file
func (a *agentT) ConfigureOperators(read func() ([]byte, error)) error {
	if read == nil {
		return errors.New("invalid argument: ReadConfig function is nil")
	}
	buf, err0 := read()
	if err0 != nil {
		return err0
	}
	var ops []Operator

	err := json.Unmarshal(buf, &ops)
	if err != nil {
		return err
	}
	ops, err = initOperators(ops)
	if err == nil {
		a.defaultOperators = ops
	}
	return err
}

func (a *agentT) Log(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	LogWithOperators(a.defaultOperators, traffic, start, duration, route, req, resp, thresholds)
}

func LogWithOperators(operators []Operator, traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	if len(operators) == 0 {
		log.Printf("%v\n", "{ \"error\" : \"no operators configured\" }")
		return
	}
	e := newEvent(traffic, start, duration, route, req, resp, thresholds)
	s := writeJson(operators, e)
	log.Printf("%v\n", s)
}

func newThreshold(resp *http.Response) Threshold {
	limit := resp.Header.Get(RateLimitName)
	resp.Header.Del(RateLimitName)
	timeout := resp.Header.Get(TimeoutName)
	resp.Header.Del(TimeoutName)
	redirect := resp.Header.Get(RedirectName)
	resp.Header.Del(RedirectName)
	return Threshold{Timeout: timeout, RateLimit: limit, Redirect: redirect}
}

/*
func (t *agentT) run() {
	go func() {
		for {
			select {
			case msg := <-t.ch.C:
				fmt.Printf("test: agent.Message() -> %v", msg)
				switch msg.Name {
				case messaging.ShutdownEvent:
					t.ch.Close()
					return
				default:
				}
			default:
			}
		}
	}()
}


*/
