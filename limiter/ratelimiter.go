package limiter

import (
	"fmt"
	"golang.org/x/time/rate"
)

const (
	InfValue     = "-1"
	DefaultBurst = 1
)

type rateLimiter struct {
	limit   rate.Limit
	burst   int
	limiter *rate.Limiter
}

func NewRateLimiter(limit rate.Limit, burst int) *rateLimiter {
	r := new(rateLimiter)
	r.limit = limit
	r.burst = burst
	r.limiter = rate.NewLimiter(r.limit, r.burst)
	return r
}

func (r *rateLimiter) String() string {
	return fmt.Sprintf("rate: %v burst: %v", r.limit, r.burst)
}

func (r *rateLimiter) Limit() rate.Limit {
	return r.limit
}

func (r *rateLimiter) Burst() int {
	return r.burst
}

func (r *rateLimiter) SetValues(limit rate.Limit, burst int) {
	r.limit = limit
	r.burst = burst
	r.limiter = rate.NewLimiter(limit, burst)
}

func (r *rateLimiter) Allow() bool {
	if r.limit == rate.Inf {
		return true
	}
	return r.limiter.Allow()
}

/*
// RateLimiter - interface for rate limiting
type RateLimiter interface {
	State
	Actuator
	Allow() bool
	StatusCode() int
	Limit() rate.Limit
	Burst() int
}

type RateLimiterConfig struct {
	Enabled    bool
	StatusCode int
	Limit      rate.Limit
	Burst      int
}

var nilRateLimiter = newRateLimiter(NilBehaviorName, nil, NewRateLimiterConfig(false, 0, 1, 1))

func NewRateLimiterConfig(enabled bool, statusCode int, limiter rate.Limit, burst int) *RateLimiterConfig {
	c := new(RateLimiterConfig)
	c.Limit = limiter
	c.Burst = burst
	if statusCode <= 0 {
		statusCode = http.StatusTooManyRequests
	}
	c.StatusCode = statusCode
	c.Enabled = enabled
	return c
}



func cloneRateLimiter(curr *rateLimiter) *rateLimiter {
	newLimiter := new(rateLimiter)
	*newLimiter = *curr
	return newLimiter
}

func newRateLimiter(name string, table *table, config *RateLimiterConfig) *rateLimiter {
	t := new(rateLimiter)
	t.name = name
	t.table = table
	t.config = RateLimiterConfig{Limit: rate.Inf, Burst: DefaultBurst}
	if config != nil {
		t.config = *config
	}
	t.rateLimiter = rate.NewLimiter(t.config.Limit, t.config.Burst)
	return t
}

func (r *rateLimiter) validate() error {
	if r.config.Limit < 0 {
		return errors.New(fmt.Sprintf("invalid configuration: RateLimiter limiter is < 0 [%v]", r.name))
	}
	if r.config.Burst < 0 {
		return errors.New(fmt.Sprintf("invalid configuration: RateLimiter burst is < 0 [%v]", r.name))
	}
	return nil
}

func (r *rateLimiter) state() (rate.Limit, int) {
	var limiter rate.Limit = -1
	var burst = -1

	if r != nil && r.IsEnabled() {
		limiter = r.config.Limit
		if limiter == rate.Inf {
			limiter = RateLimitInfValue
		}
		burst = r.config.Burst
	}
	return limiter, burst
}

func (r *rateLimiter) IsEnabled() bool { return r.config.Enabled }

func (r *rateLimiter) IsNil() bool { return r.name == NilBehaviorName }

func (r *rateLimiter) Enable() {
	if r.IsEnabled() {
		return
	}
	r.enableRateLimiter(true)
}

func (r *rateLimiter) Disable() {
	if !r.IsEnabled() {
		return
	}
	r.enableRateLimiter(false)
}

func (r *rateLimiter) Signal(values url.Values) error {
	if r.IsNil() {
		return errors.New("invalid signal: rate limiter is not configured")
	}
	if values == nil {
		return errors.New("invalid argument: values are nil for rate limiter signal")
	}
	UpdateEnable(r, values)
	limiter, burst, err := ParseLimitAndBurst(values)
	if err != nil {
		return err
	}
	if limiter != -1 || burst != -1 {
		if limiter == -1 {
			limiter = r.config.Limit
		}
		if burst == -1 {
			burst = r.config.Burst
		}
		if r.config.Limit != limiter || r.config.Burst != burst {
			r.setRateLimiter(limiter, burst)
		}
	}
	return nil
}


*/

/*





 */

/*
func (r *rateLimiter) SetRateLimiter(limiter rate.Limit, burst int) {
	validateLimiter(&limiter, &burst)
	if r.config.Limit == limiter && r.config.Burst == burst {
		return
	}
	r.table.setRateLimiter(r.name, RateLimiterConfig{Limit: limiter, Burst: burst})
}

func (r *rateLimiter) AdjustRateLimiter(percentage int) bool {
	newLimit, ok := limitAdjust(float64(r.config.Limit), percentage)
	if !ok {
		return false
	}
	newBurst, ok1 := burstAdjust(r.config.Burst, percentage)
	if !ok1 {
		return false
	}
	r.table.setRateLimiter(r.name, RateLimiterConfig{Limit: rate.Limit(newLimit), Burst: newBurst})
	return true
}

func limitAdjust(val float64, percentage int) (float64, bool) {
	change := (math.Abs(float64(percentage)) / 100.0) * val
	if change >= val {
		return val, false
	}
	if percentage > 0 {
		return val + change, true
	}
	return val - change, true
}

func burstAdjust(val int, percentage int) (int, bool) {
	floatChange := (math.Abs(float64(percentage)) / 100.0) * float64(val)
	change := int(math.Round(floatChange))
	if change == 0 || change >= val {
		return val, false
	}
	if percentage > 0 {
		return val + change, true
	}
	return val - change, true
}

*/

/*
func (r *rateLimiter) enableRateLimiter(enabled bool) {
	if r.table == nil || r.IsNil() {
		return
	}
	r.table.mu.Lock()
	defer r.table.mu.Unlock()
	if ctrl, ok := r.table.controllers[r.name]; ok {
		c := cloneRateLimiter(ctrl.rateLimiter)
		c.config.Enabled = enabled
		r.table.update(r.name, cloneController[*rateLimiter](ctrl, c))
	}
}

func (r *rateLimiter) setRateLimiter(limiter rate.Limit, burst int) {
	if r.table == nil || r.IsNil() {
		return
	}
	r.table.mu.Lock()
	defer r.table.mu.Unlock()
	if ctrl, ok := r.table.controllers[r.name]; ok {
		c := cloneRateLimiter(ctrl.rateLimiter)
		c.config.Limit = limiter
		c.config.Burst = burst
		// Not cloning the limiter as an old reference will not cause stale data when logging
		c.rateLimiter = rate.NewLimiter(limiter, burst)
		r.table.update(r.name, cloneController[*rateLimiter](ctrl, c))
	}
}


*/

/*
func (r *rateLimiter) setRateBurst(burst int) {
	if r.table == nil {
		return
	}
	r.table.mu.Lock()
	defer r.table.mu.Unlock()
	if ctrl, ok := r.table.controllers[r.name]; ok {
		c := cloneRateLimiter(ctrl.rateLimiter)
		c.config.Burst = burst
		// Not cloning the limiter as an old reference will not cause stale data when logging
		c.rateLimiter.SetBurst(burst)
		r.table.update(r.name, cloneController[*rateLimiter](ctrl, c))
	}
}


*/
/*
func (t *table) setRateLimiter(name string, config RateLimiterConfig) {
	if name == "" {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if ctrl, ok := t.controllers[name]; ok {
		c := cloneRateLimiter(ctrl.rateLimiter)
		c.config.Limit = config.Limit
		c.config.Burst = config.Burst
		c.rateLimiter = rate.NewLimiter(c.config.Limit, c.config.Burst)
		t.update(name, cloneController[*rateLimiter](ctrl, c))
	}
}

*/
