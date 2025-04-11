package limiter

import (
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
)

const (
	defaultScore = 99
)

// TODO : need to create a history of metrics + actions.
// Q: Do we need percentage of status code 429?
// A: No, any status code 429, given a stable service, needs to lead to an increase in the rate
type stats struct {
	gradiant    float64
	timeToLive  int                   // milliseconds
	intervals   int                   // number of intervals until reaching threshold
	centile     timeseries.Percentile // 99th percentile in milliseconds
	status429   int                   // count of status code 429.
	limitChange int                   // + or - percentage change
}

// master attention
func masterAttend(agent *agentT, ts *timeseries.Interface) {
	agent.dispatch(agent.master, messaging.StartupEvent)
	paused := false
	var history []stats

	for {
		select {
		case m := <-agent.master.C:
			agent.dispatch(agent.master, m.Event())
			switch m.Event() {
			case metricsEvent:
				if !paused {
					if ms, ok := metricsContent(m); ok {
						s := newStats(agent, ts, ms)
						// TODO : determine action
						// If the alpha is less than one, then determine if we need to increase the rate limiting
						// based on the internal generated 429.
						if s.gradiant > 1.0 {
						}
						history = append(history, s)

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

func newStats(agent *agentT, ts *timeseries.Interface, m metrics) stats {
	s := stats{centile: timeseries.Percentile{Score: defaultScore}, status429: m.Status429}

	// run statics calculations
	alpha, _ := ts.LinearRegression(m.Regression.X, m.Regression.Y, m.Regression.Weights, m.Regression.Origin)
	s.gradiant = alpha
	ts.Percentile(m.Regression.X, m.Regression.Weights, false, &s.centile)
	// TODO : calculate timeToLive, intervals.
	return s
}
