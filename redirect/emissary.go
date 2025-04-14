package redirect

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT, resolver *content.Resolution, s messaging.Spanner) {
	//agent.dispatch(agent.emissary, messaging.StartupEvent)
	paused := false

	for {
		select {
		case <-agent.ticker.C():
			if !paused {
				m := newMetrics()
				for e := agent.events.Dequeue(); e != nil; {
					m.update(e)
				}
				agent.Message(newMetricsMessage(*m))
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
