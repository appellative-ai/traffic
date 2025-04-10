package metrics

import "github.com/behavioral-ai/collective/timeseries"

type Redirect struct {
	Percentile timeseries.Percentile
	StatusCode StatusCodeSample `json:"status-code"`
}
