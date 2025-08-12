package representation1

import (
	"encoding/json"
	"github.com/appellative-ai/core/fmtx"
	"net/http"
	"strings"
	"time"
)

//Fragment     = "v1"

const (
	AppHostKey     = "app-host"
	CacheHostKey   = "cache-host"
	TimeoutKey     = "timeout"
	IntervalKey    = "interval"
	defaultTimeout = time.Millisecond * 2500
)

type Redirect struct {
	Template            bool
	Path                string   `json:"path"`          // Redirected path
	LoadSchedule        string   `json:"load-schedule"` //"load-schedule": "10,20,40,100",
	StatusCodes         []string `json:"status-codes"`  // Status Codes to monitor : "200", "2xx", "5xx"
	StatusCodeThreshold int      `json:"status-code-threshold"`
	Percentile          string   `json:"percentile"`
	PercentileThreshold int      `json:"percentile-threshold"`
	Codes               *StatusCodeMetrics
	Latency             *PercentileMetrics
}

func (r Redirect) Failed() bool {
	return r.Latency.Failed() || r.Codes.Failed()
}

type Route struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"` // ??Needs to allow templates, basically an '*' to match ranges
	Redirect Redirect `json:"redirect"`
}

type RoutingTable struct {
	Version string
	StartTS string
	StopTS  string
	Routes  []*Route
}

func NewRoutingTable(buf []byte) (*RoutingTable, error) {
	t := new(RoutingTable)
	err := json.Unmarshal(buf, &t)
	if err == nil {
		for _, r := range t.Routes {
			if strings.Index(r.Path, "*") != -1 {
				r.Redirect.Template = true
			}
			r.Redirect.Codes = new(StatusCodeMetrics)
			r.Redirect.Latency = new(PercentileMetrics)
		}
	}
	return t, err
}

func (t *RoutingTable) Enabled() bool {
	return true
}

func (t *RoutingTable) Route(r *http.Request) *Route {
	if r == nil {
		return nil
	}
	for _, route := range t.Routes {
		if route.Name == r.URL.Path {
			return route
		}
	}
	return nil
}

type Routing struct {
	//EnabledT  bool
	//FailedT   bool
	Running   bool
	AppHost   string        `json:"app-host"` // User requirement
	CacheHost string        `json:"cache-host"`
	Interval  time.Duration `json:"interval"`
	Timeout   time.Duration `json:"timeout"`
}

func Initialize(m map[string]string) *Routing {
	r := new(Routing)
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
	s := m[AppHostKey]
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
