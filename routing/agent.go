package routing

import (
	"fmt"
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/core/access2"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"github.com/behavioral-ai/core/uri"
	"github.com/behavioral-ai/traffic/routing/representation1"
	"github.com/behavioral-ai/traffic/timeseries"
	"net/http"
)

const (
	NamespaceName = "test:resiliency:agent/routing/request/http"
	defaultRoute  = "test:core:routing/default"
)

var (
	serverErrorResponse = httpx.NewResponse(http.StatusInternalServerError, nil, nil)
)

type agentT struct {
	events   *list
	state    *representation1.Routing
	exchange rest.Exchange
	service  *operations.Service

	review   *messaging.Review
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent(representation1.Initialize(nil), nil, operations.Serve)
	})
}

func ConstructorOverride(m map[string]string, ex rest.Exchange, service *operations.Service) {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent(representation1.Initialize(m), ex, service)
	})
}

func newAgent(state *representation1.Routing, ex rest.Exchange, service *operations.Service) *agentT {
	a := new(agentT)
	a.state = state
	a.service = service
	if ex == nil {
		a.exchange = httpx.Do
	} else {
		a.exchange = ex
	}
	a.events = newList()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.state.Interval)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Name == messaging.ConfigEvent {
		a.configure(m)
	}
}

// Run - run the agent
func (a *agentT) run() {
	go masterAttend(a, timeseries.Functions)
	go emissaryAttend(a)
}

// Link  - implementation for rest.Exchangeable interface
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		var status *messaging.Status

		url := uri.BuildURL(a.state.AppHost, r.URL.Path, r.URL.Query())
		// TODO : need to check and remove Caching header.
		resp, status = do(a, r.Method, url, httpx.CloneHeaderWithEncoding(r), r.Body)
		if status.Err != nil {
			a.service.Message(messaging.NewStatusMessage(status.WithLocation(a.Name()), a.Name()))
		}
		if resp.StatusCode == http.StatusGatewayTimeout {
			resp.Header.Add(access2.XTimeout, fmt.Sprintf("%v", a.state.Timeout))
		}
		return resp, status.Err
	}
}

func (a *agentT) trace(task, observation, action string) {
	if a.review == nil {
		return
	}
	if !a.review.Started() {
		a.review.Start()
	}
	if a.review.Expired() {
		return
	}
	a.service.Trace(a.Name(), task, observation, action)
}

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		cfg, status := messaging.MapContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.state.Update(cfg)
	case messaging.ContentTypeReview:
		r, status := messaging.ReviewContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.review = r
	case rest.ContentTypeExchange:
		ex, status := rest.ExchangeContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.exchange = ex
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}

func (a *agentT) enabled() bool {
	//if !a.state.Enabled() {
	//	return false
	//}
	//if a.state.Failed() {
	//	return false
	//}
	//if !a.limiter.Allow() {
	//	return false
	//}
	return true
}

func (a *agentT) emissaryShutdown() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

// Link - chainable exchange
/*
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		if !a.enabled() {
			return next(req)
		}
		var (
			start  = time.Now().UTC()
			newReq = req
		)
		resp, err = next(newReq)
		a.events.Enqueue(&routing.event{duration: time.Since(start), statusCode: resp.StatusCode})
		return
	}
}


*/
