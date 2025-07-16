package limiter

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/traffic/timeseries"
)

const (
	defaultScore = float64(99.0)
)

// master attention
func masterAttend(a *agentT, ts *timeseries.Interface) {
	a.dispatch(a.master, messaging.StartupEvent)
	paused := false
	var history []stats

	for {
		select {
		case m := <-a.master.C:
			a.dispatch(a.master, m.Name)
			switch m.Name {
			case metricsEvent:
				if !paused {
					if ms, status := metricsContent(m); status.OK() {
						s := newStats(a, ts, ms)
						// TODO : determine action
						// If the alpha is less than one, then determine if we need to increase the rate limiting
						// based on the internal generated 429.
						if s.gradiant > 1.0 {
						}
						history = append(history, s)
						a.trace(NamespaceTaskName, s.observation(), s.action())
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
