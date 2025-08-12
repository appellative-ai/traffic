package routing

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/traffic/timeseries"
)

// master attention
func masterAttend(a *agentT, ts *timeseries.Interface) {
	//a.dispatch(a.master, messaging.StartupEvent)
	paused := false

	for {
		select {
		case msg := <-a.master.C:
			//a.dispatch(a.master, msg.Name)
			switch msg.Name {
			case metricsEvent:
				if !paused {
					if _, status := metricsContent(msg); status.OK() {
						//updateRedirect(a, ts, m)
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
