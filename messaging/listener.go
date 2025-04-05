package messaging

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	ListenerEvent    = "event:listener"
	ContentTypeAgent = "content-type/agent"
)

func NewConfigListenerMessage(agent messaging.Agent) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, ListenerEvent)
	m.SetContent(ContentTypeAgent, agent)
	return m
}

func ConfigListenerContent(m *messaging.Message) messaging.Agent {
	if m.Event() != ListenerEvent || m.ContentType() != ContentTypeAgent {
		return nil
	}
	if v, ok := m.Body.(messaging.Agent); ok {
		return v
	}
	return nil
}
