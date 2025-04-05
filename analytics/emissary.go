package analytics

import (
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/metrics"
)

// emissary attention
func emissaryAttend(agent *agentT, ts *timeseries.Interface) {
	paused := false

	for {
		select {
		case <-agent.ticker.C():
			if !paused {
				m := metrics.NewMetrics()
				for e := agent.events.Dequeue(); e != nil; {
					m.Update(e)
				}
				// TODO : do wee need collective timeseries loading??
				//m := timeseries.NewLoadMessage(events)
				//ts.Message(m)

				// Broadcast metrics to all agents
				exchange.Broadcast(metrics.NewMetricsMessage(m))
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
