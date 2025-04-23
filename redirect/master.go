package redirect

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/timeseries"
)

// master attention
func masterAttend(a *agentT, ts *timeseries.Interface) {
	a.dispatch(a.master, messaging.StartupEvent)
	paused := false

	for {
		select {
		case msg := <-a.master.C:
			a.dispatch(a.master, msg.Event())
			switch msg.Event() {
			case metricsEvent:
				if !paused {
					if m, ok := metricsContent(msg); ok {
						updateRedirect(a, ts, m)
						//history = append(history, s)
					}
				}
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				a.masterShutdown()
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
