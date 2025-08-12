package limiter

import (
	"github.com/appellative-ai/core/messaging"
)

// emissary attention
func emissaryAttend(a *agentT) {
	a.dispatch(a.emissary, messaging.StartupEvent)
	paused := false

	for {
		select {
		case <-a.ticker.T.C:
			a.dispatch(a.emissary, a.ticker.Name)
			if !paused {
				m := newMetrics()
				for e := a.events.dequeue(); e != nil; {
					m.update(e)
				}
				a.Message(newMetricsMessage(*m))
				a.reviseTicker(m.count)
			}
		default:
		}
		select {
		case m := <-a.emissary.C:
			a.dispatch(a.emissary, m.Name)
			switch m.Name {
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
