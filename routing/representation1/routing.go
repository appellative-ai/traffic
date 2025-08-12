package representation1

import (
	"github.com/appellative-ai/core/fmtx"
	"time"
)

const (
	AppHostKey         = "app-host"
	CacheHostKey       = "cache-host"
	TimeoutDurationKey = "timeout-duration"
	ReviewDurationKey  = "review-duration"
	//IntervalKey        = "interval"
	defaultTimeout = time.Millisecond * 2500
)

type Routing struct {
	AppHost        string        `json:"app-host"`   // User requirement, not modifiable when running
	CacheHost      string        `json:"cache-host"` // User requirement, not modifiable when running
	Timeout        time.Duration `json:"timeout"`
	ReviewDuration time.Duration `json:"review-duration"`
}

func Initialize(m map[string]string) *Routing {
	r := new(Routing)
	r.Timeout = defaultTimeout
	parseRouting(r, m)
	return r
}

func (r *Routing) Update(m map[string]string) bool {
	return parseRouting(r, m)
}

func parseRouting(r *Routing, m map[string]string) (changed bool) {
	if r == nil || m == nil {
		return
	}
	s := m[AppHostKey]
	if s != "" {
		if r.AppHost != s {
			r.AppHost = s
			changed = true
		}
	}
	s = m[CacheHostKey]
	if s != "" {
		if r.CacheHost != s {
			r.CacheHost = s
			changed = true
		}
	}
	s = m[TimeoutDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if dur != r.Timeout {
				r.Timeout = dur
				changed = true
			}
		}
	}
	s = m[ReviewDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if r.ReviewDuration != dur {
				r.ReviewDuration = dur
				changed = true
			}
		}
	}
	return
}
