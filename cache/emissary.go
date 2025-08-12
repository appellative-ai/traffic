package cache

import (
	"github.com/appellative-ai/core/messaging"
)

// emissary attention
func emissaryAttend(a *agentT) {
	paused := false

	for {
		select {
		case <-a.ticker.T.C:
			if !paused {
				a.enabled.Store(a.state.Load().Now())
			}
		default:
		}
		select {
		case msg := <-a.emissary.C:
			switch msg.Name {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				a.emissaryShutdown()
				return
			default:
			}
		default:
		}
	}
}
