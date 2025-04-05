package analytics

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
	messaging2 "github.com/behavioral-ai/traffic/messaging"
)

// master attention
func masterAttend(a *agentT, resolver *content.Resolution) {
	paused := false
	var metrics messaging2.Metrics

	for {
		select {
		case m := <-a.master.C:
			switch m.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				a.masterShutdown()
				return
			case timeseries.LoadEvent:
				if paused {
					continue
				}
				metrics = createMetrics(a, timeseries.LoadContent(m))
				if a.listener != nil {
					a.listener.Message(messaging2.NewMetricsMessage(metrics))
				}
			default:
			}
		default:
		}
	}
}

func createMetrics(agent *agentT, events []*timeseries.Event) messaging2.Metrics {

	return messaging2.Metrics{}
}
