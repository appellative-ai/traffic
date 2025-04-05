package profile

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "resiliency:agent/behavioral-ai/traffic/profile"
)

type agentT struct {
	handler messaging.Agent
}

// New - create a new agent
func New(handler messaging.Agent) messaging.Agent {
	return newAgent(handler)
}

func newAgent(handler messaging.Agent) *agentT {
	a := new(agentT)
	a.handler = handler
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {}
