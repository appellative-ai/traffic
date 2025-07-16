package access

import (
	"github.com/appellative-ai/core/fmtx"
	"strconv"
	"time"
)

type Threshold struct {
	Timeout   any
	RateLimit any
	Redirect  any
}

func (t Threshold) TimeoutT() time.Duration {
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

func (t Threshold) RateLimitT() float64 {
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

func (t Threshold) rateLimit() float64 {
	return t.RateLimitT()
}

func (t Threshold) RedirectT() int {
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

func (t Threshold) redirect() int {
	return t.RedirectT()
}
