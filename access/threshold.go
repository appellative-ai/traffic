package access

import (
	"github.com/appellative-ai/core/fmtx"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type threshold struct {
	to    string
	limit string
	red   string
	cache string
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
			t.to = tokens[1]
		case RateLimitName:
			t.limit = tokens[1]
		case RedirectName:
			t.red = tokens[1]
		case CachedName:
			t.cache = tokens[1]
		}
	}
	h.Del(ThresholdName)
	return
}

func (t threshold) timeout() time.Duration {
	var dur time.Duration = -1

	if t.to == "" {
		return dur
	}
	i, err := fmtx.ParseDuration(t.to)
	if err == nil {
		dur = i
	}
	return dur
}

func (t threshold) rateLimit() float64 {
	var limit float64 = -1

	if t.limit == "" {
		return limit
	}
	i, _ := strconv.Atoi(t.limit)
	return float64(i)
}

func (t threshold) redirect() int {
	pct := -1
	if t.red == "" {
		return pct
	}
	i, _ := strconv.Atoi(t.red)
	return i
}

func (t threshold) cached() string {
	if t.cache == "" {
		return "false"
	}
	return t.cache
}
