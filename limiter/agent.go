package limiter

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/collective/notification"
	"github.com/appellative-ai/core/fmtx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/traffic/limiter/representation1"
	"github.com/appellative-ai/traffic/timeseries"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

const (
	NamespaceName     = "test:resiliency:agent/rate-limiting/request/http"
	NamespaceTaskName = "test:resiliency:task/analyze/traffic"
	rateLimitName     = "rate-limit" // Sync with core/access
	//rateBurstName     = "x-rate-burst" // Sync with core/access
)

type agentT struct {
	state    *representation1.Limiter
	limiter  *rate.Limiter
	events   *list
	notifier *notification.Interface

	review     *messaging.Review
	ticker     *messaging.Ticker
	master     *messaging.Channel
	emissary   *messaging.Channel
	dispatcher messaging.Dispatcher
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent()
	})
}

func newAgent() *agentT {
	a := new(agentT)
	a.state = representation1.Initialize(nil)
	a.state.Enabled = true
	a.notifier = notification.Notifier

	a.limiter = rate.NewLimiter(a.state.Limit, a.state.Burst)
	a.events = newList()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.state.PeakDuration)
	a.master = messaging.NewMasterChannel()
	a.emissary = messaging.NewEmissaryChannel()

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
	switch m.Name {
	case messaging.ConfigEvent:
		if a.state.Running {
			return
		}
		messaging.UpdateContent[*messaging.Review](m, &a.review)
		messaging.UpdateContent[messaging.Dispatcher](m, &a.dispatcher)
		messaging.UpdateMap(a.Name(), func(cfg map[string]string) {
			a.state.Update(cfg)
		}, m)
		return
	case messaging.StartupEvent:
		if a.state.Running {
			return
		}
		a.state.Running = true
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.state.Running {
			return
		}
		a.state.Running = false
	}
	switch m.Channel() {
	case messaging.ChannelMaster:
		a.master.C <- m
	case messaging.ChannelEmissary:
		a.emissary.C <- m
	default:
		a.master.C <- m
		a.emissary.C <- m
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
			h.Add(rateLimitName, fmt.Sprintf("%v", a.limiter.Limit()))
			//	h.Add(rateBurstName, fmt.Sprintf("%v", a.limiter.Burst()))
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
	a.notifier.Trace(a.Name(), task, observation, action)
}

func (a *agentT) emissaryShutdown() {
	a.ticker.Stop()
	a.emissary.Close()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

func (a *agentT) bucket() int {
	return fmtx.Milliseconds(a.ticker.Duration)
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
		a.ticker.Reset(newDuration)
	}
}

/*
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
	case messaging.ContentTypeDispatcher:
		if dispatcher, ok := messaging.DispatcherContent(m); ok {
			a.dispatcher = dispatcher
		}
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}


*/
