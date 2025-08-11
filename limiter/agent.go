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
	"sync/atomic"
	"time"
)

const (
	AgentName     = "common:resiliency:agent/rate-limiting/request/http"
	TaskName      = "common:resiliency:task/analyze/traffic"
	rateLimitName = "rate-limit" // Sync with core/access
	//rateBurstName     = "x-rate-burst" // Sync with core/access
)

type agentT struct {
	running  bool
	enabled  atomic.Bool
	state    atomic.Pointer[representation1.Limiter]
	limiter  *rate.Limiter
	events   *list
	notifier *notification.Interface

	review     atomic.Pointer[messaging.Review]
	ticker     *messaging.Ticker
	master     *messaging.Channel
	emissary   *messaging.Channel
	dispatcher messaging.Dispatcher
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(AgentName, func() messaging.Agent {
		return newAgent(notification.Notifier)
	})
}

func newAgent(notifier *notification.Interface) *agentT {
	a := new(agentT)
	a.enabled.Store(true)

	state := representation1.Initialize(nil)
	a.state.Store(state)
	a.notifier = notifier
	a.review.Store(messaging.NewReview(0))

	a.limiter = rate.NewLimiter(state.Limit, state.Burst)
	a.events = newList()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, state.PeakDuration)
	a.master = messaging.NewMasterChannel()
	a.emissary = messaging.NewEmissaryChannel()

	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return AgentName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		if a.running {
			return
		}
		if t, ok := messaging.ConfigContent[map[string]string](m); ok {
			a.state.Load().Update(t)
			// Update review
			dur := a.state.Load().ReviewDuration
			if dur > 0 {
				review := a.review.Load()
				if !review.Started() {
					a.review.Start(dur)
					//a.review.Store(review)
				}
			}
		}
		return
	case messaging.StartupEvent:
		if a.running {
			return
		}
		a.running = true
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.running {
			return
		}
		a.running = false
	case messaging.PauseEvent:
		// TODO : remove enqueued events
		a.enabled.Store(false)
	case messaging.ResumeEvent:
		a.enabled.Store(true)
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
		if !a.enabled.Load() {
			return next(req)
		}
		start := time.Now().UTC()
		if !a.limiter.Allow() {
			h := make(http.Header)
			h.Add(rateLimitName, fmt.Sprintf("%v", a.limiter.Limit()))
			a.events.Enqueue(&event{internal: true, unixMS: start.UnixMilli(), duration: time.Since(start), statusCode: resp.StatusCode})
			return &http.Response{StatusCode: http.StatusTooManyRequests, Header: h}, nil
		}
		resp, err = next(req)
		a.events.Enqueue(&event{unixMS: start.UnixMilli(), duration: time.Since(start), statusCode: resp.StatusCode})
		return
	}
}

func (a *agentT) dispatch(channel any, event string) {
	if a.dispatcher != nil {
		a.dispatcher.Dispatch(a, channel, event)
	}
}

func (a *agentT) trace(task, observation, action string) {
	if a.review.Load().Expired() {
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

	if cnt == a.state.Load().LoadSize {
		return
	}
	if cnt > 2*a.state.Load().LoadSize {
		newDuration = a.state.Load().PeakDuration
	} else {
		if cnt < a.state.Load().LoadSize/2 {
			newDuration = a.state.Load().OffPeakDuration
		}
	}
	if newDuration != 0 {
		a.ticker.Reset(newDuration)
	}
}
