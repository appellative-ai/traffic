package routing

import (
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(a *agentT) {
	paused := false

	for {
		select {
		case <-a.ticker.C():
			if !paused {
				m := newMetrics()
				for e := a.events.Dequeue(); e != nil; {
					m.update(e)
				}
				a.Message(newMetricsMessage(*m))
			}
		default:
		}
		select {
		case msg := <-a.emissary.C:
			//a.dispatch(a.emissary, msg.Name)
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
