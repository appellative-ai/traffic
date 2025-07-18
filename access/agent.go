package access

import (
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"log"
	"net/http"
	"time"
)

const (
	NamespaceName   = "test:resiliency:agent/log/access/http"
	contentEncoding = "Content-Encoding"
	failsafeUri     = "https://invalid-uri.com"
)

var (
	agent *agentT
)

type agentT struct {
	name      string
	operators []Operator
}

// init - register an agent constructor
func init() {
	// initialize Golang logging
	log.SetFlags(0)
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		agent = newAgent()
		return agent
	})
}

func newAgent() *agentT {
	a := new(agentT)
	a.name = NamespaceName
	a.operators = defaultOperators
	return a
}

func (a *agentT) Name() string { return a.name }
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		ops, status := OperatorsContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.name)
			return
		}
		if len(ops) > 0 {
			var err error
			a.operators, err = initOperators(ops)
			if err != nil {
				messaging.Reply(m, messaging.NewStatus(messaging.StatusInvalidArgument, err), a.name)
			}
		}
		messaging.Reply(m, messaging.StatusOK(), a.name)
		return
	}
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		resp, err = next(r)
		Log(a.operators, IngressTraffic, start, time.Since(start), DefaultRoute, r, resp)
		return
	}
}

func (a *agentT) log(start time.Time, duration time.Duration, req any, resp any) {
	Log(a.operators, IngressTraffic, start, duration, DefaultRoute, req, resp)

}
