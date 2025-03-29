package redirect

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/config"
	"net/http"
	"time"
)

// Namespace ID Namespace Specific String
// NID + NSS
// NamespaceName
const (
	NamespaceName = "resiliency:agent/behavioral-ai/traffic/redirect"
	minDuration   = time.Second * 10
	maxDuration   = time.Second * 15
	version       = 1
)

var (
	okResponse = httpx.NewResponse(http.StatusOK, nil, nil)
)

type agentT struct {
	running  bool
	traffic  string
	hostName string
	timeout  time.Duration

	exchange httpx.Exchange
	handler  messaging.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

// New - create a new agent
func New(handler messaging.Agent) messaging.Agent {
	return newAgent(handler)
}

func newAgent(handler messaging.Agent) *agentT {
	a := new(agentT)
	a.exchange = httpx.Do
	a.handler = handler
	a.ticker = messaging.NewTicker(messaging.Emissary, maxDuration)
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
	go emissaryAttend(a, content.Resolver, nil)
	a.running = true
}

func (a *agentT) enabled() bool {
	return false
}

// Link - chainable exchange
func (a *agentT) Link(next httpx.Exchange) httpx.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		// TODO: if a redirect is configured, then process and ignore rest of pipeline
		if a.enabled() {
			return okResponse, nil
		}
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		return
	}
}

func (a *agentT) dispatch(channel any, event1 string) {
	a.handler.Message(eventing.NewDispatchMessage(a, channel, event1))
}

func (a *agentT) reviseTicker(resolver *content.Resolution, s messaging.Spanner) {
}

func (a *agentT) emissaryShutdown() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterShutdown() {
	a.master.Close()
}

func (a *agentT) configure(m *messaging.Message) {
	var ok bool

	if a.hostName, ok = config.AppHostName(a, m); !ok {
		return
	}
	if a.timeout, ok = config.Timeout(a, m); !ok {
		return
	}
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}
