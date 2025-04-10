package limiter

import (
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT) {
	paused := false

	for {
		select {
		case <-agent.ticker.C():
			if !paused {
				m := newMetrics()
				for e := agent.events.Dequeue(); e != nil; {
					m.Update(e)
				}
				agent.Message(newMetricsMessage(*m))
				agent.reviseTicker(m.Count)
			}
		default:
		}
		select {
		case m := <-agent.emissary.C:
			switch m.Event() {
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
