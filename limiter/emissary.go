package limiter

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
				a.reviseTicker(m.count)
			}
		default:
		}
		select {
		case m := <-a.emissary.C:
			switch m.Name() {
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
