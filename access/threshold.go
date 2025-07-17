package access

import (
	"github.com/appellative-ai/core/fmtx"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type threshold struct {
	Timeout   string
	RateLimit string
	Redirect  string
	Cached    string
}

func newThreshold(v any) (t threshold) {
	var values []string
	var h http.Header

	if v == nil {
		return
	}
	if h1, ok1 := v.(http.Header); ok1 {
		h = h1
	} else {
		if resp, ok2 := v.(*http.Response); ok2 {
			h = resp.Header
		}
	}
	if h == nil {
		return
	}
	values = h.Values(ThresholdName)

	for _, s := range values {
		tokens := strings.Split(s, "=")
		if len(tokens) < 2 {
			continue
		}
		switch tokens[0] {
		case TimeoutName:
			t.Timeout = tokens[1]
		case RateLimitName:
			t.RateLimit = tokens[1]
		case RedirectName:
			t.Redirect = tokens[1]
		case CachedName:
			t.Cached = tokens[1]
		}
	}
	h.Del(ThresholdName)
	return
}

func (t threshold) timeout() time.Duration {
	var dur time.Duration = -1

	if t.Timeout == "" {
		return dur
	}
	i, err := fmtx.ParseDuration(t.Timeout)
	if err == nil {
		dur = i
	}
	return dur
}

func (t threshold) rateLimit() float64 {
	var limit float64 = -1

	if t.RateLimit == "" {
		return limit
	}
	i, _ := strconv.Atoi(t.RateLimit)
	return float64(i)
}

func (t threshold) redirect() int {
	pct := -1
	if t.Redirect == "" {
		return pct
	}
	i, _ := strconv.Atoi(t.Redirect)
	return i
}

func (t threshold) cached() string {
	if t.Cached == "" {
		return "false"
	}
	return t.Cached
}
