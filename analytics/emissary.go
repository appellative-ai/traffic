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
				// TODO : do wee need collective timeseries loading?
				// Message based on subscriptions
				if subs, ok := agent.catalog.Lookup(metrics.Event); ok {
					for _, item := range subs {
						msg := metrics.NewMetricsMessage(item, NamespaceName, *m)
						exchange.Message(msg)
					}
				}
				agent.reviseTicker(m.Count)
			}
		default:
		}
		select {
		case m := <-agent.emissary.C:
			switch m.Event() {
			case messaging.SubscriptionCreateEvent:
				// ignore error, message creation has error handling
				_ = agent.catalog.CreateWithMessage(m)
				if m.Reply != nil {
					go func() {
						messaging.Reply(messaging.NewMessage(messaging.ChannelControl, m.Event()), messaging.StatusOK(), NamespaceName)
					}()
				}
			case messaging.SubscriptionCancelEvent:
				agent.catalog.CancelWithMessage(m)
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
