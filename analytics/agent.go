package analytics

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	messaging2 "github.com/behavioral-ai/traffic/messaging"
	"net/http"
	"time"
)

// Namespace ID Namespace Specific String
// NID + NSS
// NamespaceName
const (
	NamespaceName = "resiliency:agent/behavioral-ai/traffic/analytics"
	//minDuration   = time.Second * 10
	//dmaxDuration   = time.Second * 15
	duration = time.Minute * 5
	loadSize = 200
)

type agentT struct {
	running bool
	enabled bool
	traffic string
	origin  timeseries.Origin
	events  *list

	listener messaging.Agent
	handler  messaging.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

// New - create a new analytics agent
func New(handler messaging.Agent) messaging.Agent {
	return newAgent(handler)
}

func newAgent(handler messaging.Agent) *agentT {
	a := new(agentT)
	a.enabled = true
	a.handler = handler
	a.events = newList()

	a.ticker = messaging.NewTicker(messaging.Emissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Event() == messaging.ConfigEvent {
		a.configure(m)
		return
	}
	if m.Event() == messaging.StartupEvent {
		a.run()
		return
	}
	if !a.running {
		return
	}
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.C <- m
	case messaging.Master:
		a.master.C <- m
	case messaging.Control:
		a.emissary.C <- m
		a.master.C <- m
	default:
		a.emissary.C <- m
	}
}

// Run - run the agent
func (a *agentT) run() {
	if a.running {
		return
	}
	go masterAttend(a, content.Resolver)
	go emissaryAttend(a, timeseries.Functions)
	// Need to start worker go routine, but need to wait for the traffic profile to be initialized
	a.running = true
}

// Link - chainable exchange
func (a *agentT) Link(next httpx.Exchange) httpx.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()

		resp, err = next(req)
		if a.enabled {
			a.events.Enqueue(&timeseries.Event{Origin: a.origin, Start: start,
				Duration: milliseconds(time.Since(start)), StatusCode: resp.StatusCode,
			})
		}
		return
	}
}

/*
func (a *agentT) dispatch(channel any, event1 string) {
	a.handler.Message(eventing.NewDispatchMessage(a, channel, event1))
}
*/

func (a *agentT) emissaryShutdown() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

func (a *agentT) configure(m *messaging.Message) {
	if o, ok := timeseries.NewOriginFromMessage(a, m); ok {
		a.origin = o
	}
	a.listener = messaging2.ConfigListenerContent(m)
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}

func milliseconds(duration time.Duration) int {
	if duration <= 0 {
		return -1
	}
	return int(duration / time.Duration(1e6))
}

func (a *agentT) reviseTicker(cnt int) {
	var (
		newDuration  time.Duration
		currDuration = a.ticker.Duration()
	)

	if cnt == loadSize {
		return
	}
	if cnt > loadSize {
		newDuration = currDuration / 2
	} else {
		newDuration = currDuration + time.Second*20
	}
	if newDuration != 0 {
		a.ticker.Start(newDuration)
	}
}
