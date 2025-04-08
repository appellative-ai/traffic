package limiter

import (
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/metrics"
	"github.com/behavioral-ai/traffic/urn"
)

const (
	percentile = float64(0.95)
)

// master attention
func masterAttend(agent *agentT, ts *timeseries.Interface) {
	agent.dispatch(agent.master, messaging.StartupEvent)
	paused := false
	exchange.Message(messaging.NewSubscriptionCreateMessage(urn.AnalyticsAgent, NamespaceName, metrics.Event))

	for {
		select {
		case m := <-agent.master.C:
			agent.dispatch(agent.master, m.Event())
			switch m.Event() {
			case metrics.Event:
				if !paused {
					if ms, ok := metrics.MetricsContent(m); ok {
						alpha, beta := ts.LinearRegression(ms.Regression.X, ms.Regression.Y, ms.Regression.Weights, ms.Regression.Origin)
						if alpha > 0.0 && beta > 0.0 {
						}
						latency := ts.Percentile(ms.Percentile.X, ms.Percentile.Weights, false, percentile)
						if latency > 0.0 {
						}
					}
				}
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				exchange.Message(messaging.NewSubscriptionCancelMessage(urn.AnalyticsAgent, NamespaceName, metrics.Event))
				agent.masterShutdown()
				return
			default:
			}
		default:
		}
	}
}
