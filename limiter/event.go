package limiter

const (
	metricsEvent = "event:profile"
)

type metrics struct {
	Duration   int //milliseconds
	StatusCode int
}

//func NewEvent(resp *http.)
