package limiter

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/timeseries"
	"time"
)

const (
	defaultScore = float64(99.0)
)

// TODO : need to create a history of metrics + actions.
// Q: Do we need percentage of status code 429?
// A: No, any status code 429, given a stable service, needs to lead to an increase in the rate
type stats struct {
	unixMS      int64
	gradiant    float64
	timeToLive  int     // milliseconds
	intervals   int     // number of intervals until reaching threshold
	latency     float64 // 99th percentile in milliseconds
	status429   int     // count of status code 429.
	limitChange int     // + or - percentage change
}

// master attention
func masterAttend(a *agentT, ts *timeseries.Interface) {
	a.dispatch(a.master, messaging.StartupEvent)
	paused := false
	var history []stats

	for {
		select {
		case m := <-a.master.C:
			a.dispatch(a.master, m.Name())
			switch m.Name() {
			case metricsEvent:
				if !paused {
					if ms, ok := metricsContent(m); ok {
						s := newStats(a, ts, ms)
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
				a.masterShutdown()
				return
			default:
			}
		default:
		}
	}
}

func newStats(agent *agentT, ts *timeseries.Interface, m metrics) stats {
	s := stats{unixMS: time.Now().UTC().UnixMilli(), status429: m.status429}

	// run statics calculations
	alpha, _ := ts.LinearRegression(m.regression.x, m.regression.y, m.regression.weights, m.regression.origin)
	s.gradiant = alpha
	s.latency = ts.Percentile(m.regression.x, m.regression.weights, false, defaultScore)
	// TODO : calculate timeToLive, intervals.
	return s
}
