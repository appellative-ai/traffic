package analytics

import (
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
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
	catalog *messaging.Catalog

	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
	handler  eventing.Agent
}

// New - create a new analytics agent
func init() {
	a := newAgent(eventing.Handler)
	exchange.Register(a)
}

func newAgent(handler eventing.Agent) *agentT {
	a := new(agentT)
	a.enabled = true
	a.handler = handler
	a.events = newList()
	a.catalog = new(messaging.Catalog)

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
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
	if !a.running {
		if m.Event() == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Event() == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Event() == messaging.ShutdownEvent {
		a.running = false
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
	go masterAttend(a)
	go emissaryAttend(a, timeseries.Functions)
}

// Link - chainable exchange
func (a *agentT) Link(next httpx.Exchange) httpx.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()

		resp, err = next(req)
		if a.enabled {
			a.events.Enqueue(&timeseries.Event{Origin: a.origin, Start: start,
				Duration: time.Since(start), StatusCode: resp.StatusCode,
			})
		}
		return
	}
}

func (a *agentT) emissaryShutdown() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		if o, ok := timeseries.NewOriginFromMessage(a, m); ok {
			a.origin = o
		}
	}
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}

/*
func milliseconds(duration time.Duration) int {
	if duration <= 0 {
		return -1
	}
	return int(duration / time.Duration(1e6))
}

*/

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
