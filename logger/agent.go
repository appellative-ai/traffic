package logger

import (
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/logx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"net/http"
	"time"
)

const (
	NamespaceName = "test:resiliency:agent/log/access/http"
	defaultRoute  = "host"
)

// Agent - agent
type Agent interface {
	messaging.Agent
	LogEgress(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)
	LogStatus(status *std.Status)
}

var (
	agent *agentT
)

type agentT struct {
	name      string
	operators []logx.Operator
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		agent = newAgent()
		return agent
	})
}

func newAgent() *agentT {
	a := new(agentT)
	a.name = NamespaceName
	return a
}

func (a *agentT) Name() string { return a.name }
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		if ops, ok := messaging.ConfigContent[[]logx.Operator](m); ok {
			if len(ops) > 0 {
				var err error
				a.operators, err = logx.InitOperators(ops)
				if err != nil {
					messaging.Reply(m, std.NewStatus(std.StatusInvalidArgument, "", err), a.name)
				}
			}
		}
		messaging.Reply(m, std.StatusOK, a.name)
		return
	}
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		resp, err = next(r)
		logx.LogAccess(a.operators, logx.IngressTraffic, start, time.Since(start), defaultRoute, r, resp)
		return
	}
}

func (a *agentT) LogEgress(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	logx.LogEgress(a.operators, start, duration, route, req, resp, timeout)
}

func (a *agentT) LogStatus(status *std.Status) {
	logx.LogStatus(status)
}
