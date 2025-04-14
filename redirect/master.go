package redirect

import (
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
)

// master attention
func masterAttend(agent *agentT, ts *timeseries.Interface) {
	agent.dispatch(agent.master, messaging.StartupEvent)
	paused := false

	for {
		select {
		case msg := <-agent.master.C:
			agent.dispatch(agent.master, msg.Event())
			switch msg.Event() {
			case metricsEvent:
				if !paused {
					if m, ok := metricsContent(msg); ok {
						updateRedirect(agent, ts, m)
						//history = append(history, s)
					}
				}
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.masterShutdown()
				return
			default:
			}
		default:
		}
	}
}

func updateRedirect(agent *agentT, ts *timeseries.Interface, m metrics) {

	ts.Percentile(m.x, m.weights, false, float64(agent.redirect.Latency.Score))
	// TODO : calculate timeToLive, intervals.
	//return s
}
