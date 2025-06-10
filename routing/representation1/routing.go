package representation1

import (
	"github.com/behavioral-ai/core/fmtx"
	"time"
)

const (
	Fragment     = "v1"
	logRouteName = "app"

	AppHostKey   = "app-host"
	CacheHostKey = "cache-host"
	LogKey       = "log"
	LogRouteKey  = "route-name"
	TimeoutKey   = "timeout"

	defaultTimeout = time.Millisecond * 2500
)

type Redirect struct {
	//Name                string   `json:"name"`
	Path                string   `json:"path"`         // Redirected path
	StatusCodes         []string `json:"status-codes"` // Status Codes to monitor : "200", "2xx", "5xx"
	StatusCodeThreshold int      `json:"status-code-threshold"`
	Percentile          string   `json:"percentile"`
	PercentileThreshold int      `json:"percentile-threshold"`
	Codes               *StatusCodeMetrics
	Latency             *PercentileMetrics
}

type Route struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"` // ??Needs to allow templates, basically an '*' to match ranges
	Redirect Redirect `json:"redirect"`
}

type RoutingTable struct {
	Version     string
	EffectiveTS string
	Routes      []Route
}

type Routing struct {
	EnabledT     bool
	FailedT      bool
	Log          bool          `json:"log"`
	AppHost      string        `json:"app-host"` // User requirement
	CacheHost    string        `json:"cache-host"`
	LogRouteName string        `json:"route-name"`
	Interval     time.Duration `json:"interval"`
	Timeout      time.Duration `json:"timeout"`
	//Latency      *PercentileMetrics
	//Codes        *StatusCodeMetrics
}

func (r *Routing) Enabled() bool {
	return false
}

func (r *Routing) Failed() bool {
	return true //r.Latency.Failed() || r.Codes.Failed()
}

func Initialize(m map[string]string) *Routing {
	r := new(Routing)
	r.Log = true
	r.LogRouteName = logRouteName
	r.Timeout = defaultTimeout
	r.Interval = time.Millisecond * 2000
	parseRouting(r, m)
	return r
}

/*
func NewRouting(name string) *Routing {
	//m, _ := resource.Resolve[map[string]string](name, Fragment, resource.Resolver)
	return newRouting(nil)
}


*/

func newRouting(m map[string]string) *Routing {
	c := Initialize(m)
	return c
}

func (r *Routing) Update(m map[string]string) {
	parseRouting(r, m)
}

func parseRouting(r *Routing, m map[string]string) {
	if r == nil || m == nil {
		return
	}
	s := m[LogKey]
	if s != "" {
		if s == "true" {
			r.Log = true
		} else {
			r.Log = false
		}
	}
	s = m[LogRouteKey]
	if s != "" {
		r.LogRouteName = s
	}
	s = m[AppHostKey]
	if s != "" {
		r.AppHost = s
	}
	s = m[CacheHostKey]
	if s != "" {
		r.CacheHost = s
	}
	s = m[TimeoutKey]
	if s != "" {
		dur, err := fmtx.ParseDuration(s)
		if err != nil {
			//messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
			return
		}
		r.Timeout = dur
	}
	s = m[IntervalKey]
	if s != "" {
		dur, err := fmtx.ParseDuration(s)
		if err != nil {
			//messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
			return
		}
		r.Interval = dur
	}
}
