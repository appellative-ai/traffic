package analytics

import (
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT, ts *timeseries.Interface) {
	paused := false

	for {
		select {
		case <-agent.ticker.C():
			if !paused {
				events := make([]*timeseries.Event, loadSize)
				for e := agent.events.Dequeue(); e != nil; {
					events = append(events, e)
				}
				m := timeseries.NewLoadMessage(events)
				ts.Message(m)
				agent.Message(m.SetChannel(messaging.Master))
				agent.reviseTicker(len(events))
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
