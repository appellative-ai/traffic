package redirect

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT, resolver *content.Resolution, s messaging.Spanner) {
	agent.dispatch(agent.emissary, messaging.StartupEvent)
	paused := false
	agent.reviseTicker(resolver, s)

	for {
		select {
		case <-agent.ticker.C():
			agent.dispatch(agent.ticker, messaging.ObservationEvent)
			if !paused {
				agent.reviseTicker(resolver, s)
			}
		default:
		}
		select {
		case msg := <-agent.emissary.C:
			agent.dispatch(agent.emissary, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.emissaryShutdown()
				return
			default:
			}
		default:
		}
	}
}
