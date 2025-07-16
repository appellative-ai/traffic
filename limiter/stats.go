package limiter

import (
	"github.com/appellative-ai/traffic/timeseries"
	"time"
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

func newStats(agent *agentT, ts *timeseries.Interface, m metrics) stats {
	s := stats{unixMS: time.Now().UTC().UnixMilli(), status429: m.status429}

	// run statics calculations
	alpha, _ := ts.LinearRegression(m.regression.x, m.regression.y, m.regression.weights, m.regression.origin)
	s.gradiant = alpha
	s.latency = ts.Percentile(m.regression.x, m.regression.weights, false, defaultScore)
	// TODO : calculate timeToLive, intervals.
	return s
}

func (s stats) observation() string {
	return ""
}

func (s stats) action() string {
	return ""
}
