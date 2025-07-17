package access

import (
	"github.com/appellative-ai/core/fmtx"
	"net/http"
	"strconv"
	"time"
)

type Threshold struct {
	Timeout   any
	RateLimit any
	Redirect  any
	Cached    any
}

func newThreshold(resp *http.Response) Threshold {
	limit := resp.Header.Get(RateLimitName)
	timeout := resp.Header.Get(TimeoutName)
	redirect := resp.Header.Get(RedirectName)
	cached := resp.Header.Get(CachedName)
	resp.Header.Del(ThresholdName)
	return Threshold{Timeout: timeout, RateLimit: limit, Redirect: redirect, Cached: cached}
}

func (t Threshold) timeout() time.Duration {
	var dur time.Duration = -1

	if t.Timeout == nil {
		return -1
	}
	if s, ok := t.Timeout.(string); ok {
		if s == "" {
			return dur
		}
		i, err := fmtx.ParseDuration(s)
		if err == nil {
			dur = i
		}
	} else {
		if d, ok1 := t.Timeout.(time.Duration); ok1 {
			dur = d
		}
	}
	return dur
}

func (t Threshold) rateLimit() float64 {
	var limit float64 = -1

	if t.RateLimit == nil {
		return limit
	}
	if s, ok := t.RateLimit.(string); ok {
		if s == "" {
			return limit
		}
		i, _ := strconv.Atoi(s)
		return float64(i)
	}
	if l, ok1 := t.RateLimit.(float64); ok1 {
		return l
	}
	if l, ok1 := t.RateLimit.(int); ok1 {
		return float64(l)
	}
	return limit
}

/*
func (t Threshold) rateLimit() float64 {
	return t.RateLimitT()
}


*/

func (t Threshold) redirect() int {
	pct := -1
	if t.Redirect == nil {
		return pct
	}
	if s, ok := t.Redirect.(string); ok {
		if s == "" {
			return pct
		}
		i, _ := strconv.Atoi(s)
		pct = i
	} else {
		if d, ok1 := t.Redirect.(int); ok1 {
			pct = d
		}
	}
	return pct
}

/*
func (t Threshold) redirect() int {
	return t.RedirectT()
}


*/

func (t Threshold) cached() string {
	s := "false"
	if t.Cached == nil {
		return s
	}
	if s1, ok := t.Cached.(string); ok {
		return s1
	}
	return s
}
