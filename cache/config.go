package cache

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
)

func (a *agentT) config(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if ex, ok := messaging.ConfigContent[rest.Exchange](m); ok && ex != nil {
		if !a.running.Load() {
			a.exchange = ex
			return
		}
	}
	if t, ok := messaging.ConfigContent[map[string]string](m); ok {
		state := a.state.Load()
		host := state.Host
		// Changed?
		if !state.Update(t) {
			return
		}
		// If the host has changed, then reset back to original if we are running
		if host != state.Host && a.running.Load() {
			state.Host = host
		}
	}
	return
}
