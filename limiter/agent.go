package limiter

import (
	"fmt"
	"github.com/behavioral-ai/collective/repository"
	"github.com/behavioral-ai/core/access2"
	"github.com/behavioral-ai/core/eventing"
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"github.com/behavioral-ai/traffic/limiter/representation1"
	"github.com/behavioral-ai/traffic/timeseries"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

const (
	NamespaceName = "resiliency:agent/rate-limiting/request/http"
)

type agentT struct {
	state   *representation1.Limiter
	limiter *rate.Limiter
	events  *list

	ticker     *messaging.Ticker
	master     *messaging.Channel
	emissary   *messaging.Channel
	handler    eventing.Agent
	dispatcher messaging.Dispatcher
}

// New - create a new agent
func init() {
	repository.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent(eventing.Handler, representation1.NewLimiter(NamespaceName))
	})
}

func newAgent(handler eventing.Agent, state *representation1.Limiter) *agentT {
	a := new(agentT)
	if state == nil {
		a.state = representation1.Initialize()
	} else {
		a.state = state
	}
	a.state.Enabled = true
	a.limiter = rate.NewLimiter(a.state.Limit, a.state.Burst)
	a.events = newList()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.state.PeakDuration)
	a.master = messaging.NewMasterChannel()
	a.emissary = messaging.NewEmissaryChannel()
	a.handler = handler
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
		if m.Name() == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Name() == messaging.StartupEvent {
			a.run()
			a.state.Running = true
			return
		}
		return
	}
	if m.Name() == messaging.ShutdownEvent {
		a.state.Running = false
	}
	switch m.Channel() {
	case messaging.ChannelMaster:
		a.master.C <- m
	case messaging.ChannelControl:
		a.master.C <- m
	default:
		fmt.Printf("limiter - invalid channel %v\n", m)
	}
}

// Run - run the agent
func (a *agentT) run() {
	go emissaryAttend(a)
	go masterAttend(a, timeseries.Functions)
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		if !a.limiter.Allow() {
			h := make(http.Header)
			h.Add(access2.XRateLimit, fmt.Sprintf("%v", a.limiter.Limit()))
			h.Add(access2.XRateBurst, fmt.Sprintf("%v", a.limiter.Burst()))
			if a.state.Enabled {
				a.events.Enqueue(&event{internal: true, unixMS: start.UnixMilli(), duration: time.Since(start), statusCode: resp.StatusCode})
			}
			return &http.Response{StatusCode: http.StatusTooManyRequests, Header: h}, nil
		}
		resp, err = next(req)
		if a.state.Enabled {
			a.events.Enqueue(&event{unixMS: start.UnixMilli(), duration: time.Since(start), statusCode: resp.StatusCode})
		}
		return
	}
}

func (a *agentT) dispatch(channel any, event string) {
	if a.dispatcher != nil {
		a.dispatcher.Dispatch(a, channel, event)
	}
}

func (a *agentT) emissaryShutdown() {
	a.ticker.Stop()
	a.emissary.Close()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

func (a *agentT) bucket() int {
	return fmtx.Milliseconds(a.ticker.Duration())
}

func (a *agentT) reviseTicker(cnt int) {
	var newDuration time.Duration

	if cnt == a.state.LoadSize {
		return
	}
	if cnt > 2*a.state.LoadSize {
		newDuration = a.state.PeakDuration
	} else {
		if cnt < a.state.LoadSize/2 {
			newDuration = a.state.OffPeakDuration
		}
	}
	if newDuration != 0 {
		a.ticker.Start(newDuration)
	}
}

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		cfg := messaging.ConfigMapContent(m)
		if cfg == nil {
			messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Name())
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
