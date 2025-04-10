package limiter

import (
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
)

const (
	score = float64(0.95)
)

// master attention
func masterAttend(agent *agentT, ts *timeseries.Interface) {
	agent.dispatch(agent.master, messaging.StartupEvent)
	paused := false

	for {
		select {
		case m := <-agent.master.C:
			agent.dispatch(agent.master, m.Event())
			switch m.Event() {
			case metricsEvent:
				if !paused {
					if ms, ok := metricsContent(m); ok {
						alpha, beta := ts.LinearRegression(ms.Regression.X, ms.Regression.Y, ms.Regression.Weights, ms.Regression.Origin)
						if alpha > 0.0 && beta > 0.0 {
						}
						p := ts.Percentile(ms.Percentile.X, ms.Percentile.Weights, false, float64(agent.threshold.Score))
						if p.Latency > 0.0 {
						}
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
