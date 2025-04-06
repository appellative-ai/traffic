package profile

import (
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "resiliency:agent/behavioral-ai/traffic/profile"
)

type agentT struct {
	handler eventing.Agent
}

var (
	agent messaging.Agent
)

// New - create a new agent
func init() {
	agent = newAgent(eventing.Handler)
	exchange.Register(agent)
}

func newAgent(handler eventing.Agent) *agentT {
	a := new(agentT)
	a.handler = handler
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
}

func (a *agentT) configure(m *messaging.Message) {
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}
