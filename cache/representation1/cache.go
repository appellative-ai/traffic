package representation1

import (
	"github.com/appellative-ai/core/fmtx"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	Fragment            = "v1"
	HostKey             = "host"
	CacheControlKey     = "cache-control"
	TimeoutDurationKey  = "timeout-duration"
	IntervalDurationKey = "interval-duration"

	SundayKey      = "sun"
	MondayKey      = "mon"
	TuesdayKey     = "tue"
	WednesdayKey   = "wed"
	ThursdayKey    = "thu"
	FridayKey      = "fri"
	SaturdayKey    = "sat"
	rangeSeparator = "-"

	defaultInterval = time.Minute * 30
	defaultTimeout  = time.Millisecond * 5000
)

type Cache struct {
	Timeout  time.Duration
	Interval time.Duration    // Not changed if running
	Host     string           // User requirement, not changed if running
	Policy   http.Header      // User requirement
	Days     map[string]Range // User requirement
}

// Initialize - add a default policy
func Initialize(m map[string]string) *Cache {
	c := new(Cache)
	c.Timeout = defaultTimeout
	c.Interval = defaultInterval
	c.Policy = make(http.Header)
	c.Days = make(map[string]Range)
	parseCache(c, m)
	return c
}

func (c *Cache) Now() bool {
	ts := time.Now()
	day := ts.Weekday()
	s := ""
	switch day {
	case 0:
		s = SundayKey
	case 1:
		s = MondayKey
	case 2:
		s = TuesdayKey
	case 3:
		s = WednesdayKey
	case 4:
		s = ThursdayKey
	case 5:
		s = FridayKey
	case 6:
		s = SaturdayKey
	}
	return c.Days[s].In(ts)
}

func (c *Cache) Update(m map[string]string) bool {
	return parseCache(c, m)
}

func parseCache(c *Cache, m map[string]string) (changed bool) {
	if c == nil || m == nil {
		return
	}
	if c.Policy == nil {
		c.Policy = make(http.Header)
	}
	if c.Days == nil {
		c.Days = make(map[string]Range)
	}
	s := m[HostKey]
	if s != "" {
		if c.Host != s {
			c.Host = s
			changed = true
		}
	}
	s = m[TimeoutDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if c.Timeout != dur {
				c.Timeout = dur
				changed = true
			}
		}
	}
	s = m[IntervalDurationKey]
	if s != "" {
		if dur, err := fmtx.ParseDuration(s); err == nil && dur > 0 {
			if c.Interval != dur {
				c.Interval = dur
				changed = true
			}
		}
	}
	s = m[CacheControlKey]
	if s != "" {
		if c.Policy.Get(CacheControlKey) != s {
			c.Policy.Set(CacheControlKey, s)
			changed = true
		}
	}
	return parseDays(c, m)
}

func parseDays(c *Cache, m map[string]string) (changed bool) {
	if parseDay(c, SundayKey, m) {
		changed = true
	}
	if parseDay(c, MondayKey, m) {
		changed = true
	}
	if parseDay(c, TuesdayKey, m) {
		changed = true
	}
	if parseDay(c, WednesdayKey, m) {
		changed = true
	}
	if parseDay(c, ThursdayKey, m) {
		changed = true
	}
	if parseDay(c, FridayKey, m) {
		changed = true
	}
	if parseDay(c, SaturdayKey, m) {
		changed = true
	}
	return
}

func parseDay(c *Cache, key string, m map[string]string) (changed bool) {
	s := m[key]
	if s == "" {
		return
	}
	r := NewRange(s)
	if !r.Empty() {
		c.Days[key] = r
	}
	return true
}

// Range - hour range
type Range struct {
	From int
	To   int
}

func NewRange(s string) Range {
	if s == "" {
		return Range{}
	}
	tokens := strings.Split(strings.Trim(s, " "), rangeSeparator)
	if len(tokens) != 2 {
		return Range{}
	}
	r := Range{}
	if i, err := strconv.Atoi(tokens[0]); err == nil {
		r.From = i
	}
	if i, err := strconv.Atoi(tokens[1]); err == nil {
		r.To = i
	}
	return r
}

func (r Range) Empty() bool {
	if r.From < 0 || r.To <= 0 {
		return true
	}
	if r.From > 23 || r.To > 23 {
		return false
	}
	return r.From > r.To
}

func (r Range) In(ts time.Time) bool {
	hour := ts.Hour()
	return r.From <= hour && hour <= r.To
}
