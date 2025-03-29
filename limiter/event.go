package limiter

const (
	metricsEvent = "event:metrics"
)

type metrics struct {
	Duration   int //milliseconds
	StatusCode int
}

//func NewEvent(resp *http.)
