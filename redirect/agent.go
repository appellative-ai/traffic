package redirect

import (
	"github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/collective/repository"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"github.com/behavioral-ai/traffic/redirect/representation1"
	"github.com/behavioral-ai/traffic/timeseries"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

const (
	NamespaceName = "test:resiliency:agent/redirect/request/http"
)

type agentT struct {
	events  *list
	limiter *rate.Limiter
	state   *representation1.Redirect
	service *operations.Service

	ticker     *messaging.Ticker
	emissary   *messaging.Channel
	master     *messaging.Channel
	dispatcher messaging.Dispatcher
}

// init - register an agent constructor
func init() {
	repository.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent(representation1.Initialize(nil), operations.Serve)
	})
}

func ConstructorOverride(m map[string]string, service *operations.Service) {
	repository.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent(representation1.Initialize(m), service)
	})
}

func newAgent(state *representation1.Redirect, service *operations.Service) *agentT {
	a := new(agentT)
	a.state = state
	a.service = service

	a.limiter = rate.NewLimiter(a.state.Limit, a.state.Burst)
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
	if !a.state.Running {
		if m.Name == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Name == messaging.StartupEvent {
			a.run()
			a.state.Running = true
			return
		}
		return
	}
	if m.Name == messaging.ShutdownEvent {
		a.state.Running = false
	}
	switch m.Channel() {
	case messaging.ChannelEmissary:
		a.emissary.C <- m
	case messaging.ChannelMaster:
		a.master.C <- m
	case messaging.ChannelControl:
		a.emissary.C <- m
		a.master.C <- m
	default:
		a.emissary.C <- m
	}
}

// Run - run the agent
func (a *agentT) run() {
	go masterAttend(a, timeseries.Functions)
	go emissaryAttend(a)
}

func (a *agentT) enabled() bool {
	if !a.state.Enabled() {
		return false
	}
	if a.state.Failed() {
		return false
	}
	if !a.limiter.Allow() {
		return false
	}
	return true
}

// Link - chainable exchange
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
		a.events.Enqueue(&event{duration: time.Since(start), statusCode: resp.StatusCode})
		return
	}
}

func (a *agentT) dispatch(channel any, event string) {
	if a.dispatcher != nil {
		a.dispatcher.Dispatch(a, channel, event)
	}
}

func (a *agentT) trace(task, observation, action string) {
	a.service.Trace(a.Name(), task, observation, action)
}

func (a *agentT) emissaryShutdown() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

// TODO : need to configure current and redirect URL's
func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		cfg, status := messaging.MapContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.state.Update(cfg)
	case messaging.ContentTypeDispatcher:
		if dispatcher, ok := messaging.DispatcherContent(m); ok {
			a.dispatcher = dispatcher
		}
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}
