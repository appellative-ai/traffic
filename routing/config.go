package routing

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
		appHost := state.AppHost
		cacheHost := state.CacheHost
		reviewDuration := state.ReviewDuration
		// Changed?
		if !state.Update(t) {
			return
		}
		// If the host has changed, then reset back to original if we are running
		if appHost != state.AppHost && a.running.Load() {
			state.AppHost = appHost
		}
		if cacheHost != state.CacheHost && a.running.Load() {
			state.CacheHost = cacheHost
		}
		dur := a.state.Load().ReviewDuration
		if reviewDuration != dur && dur > 0 {
			review := a.review.Load()
			if !review.Started() {
				review.Start(dur)
			}
		}
	}
}
