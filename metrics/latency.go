package metrics

import (
	"github.com/behavioral-ai/collective/timeseries"
	"time"
)

type LatencySample struct {
	First time.Time
	Last  time.Time
}

func (d *LatencySample) Update(event *timeseries.Event) {
	if d.First.IsZero() {
		d.First = event.Start
	} else {
		if event.Start.Compare(d.First) == -1 {
			d.First = event.Start
		}
	}
	if d.Last.IsZero() {
		d.Last = event.Start.Add(event.Duration)
	} else {
		if event.Start.Compare(d.Last) == 1 {
			d.Last = event.Start.Add(event.Duration)
		}
	}
}

func (l *LatencySample) Elapsed() time.Duration {
	return l.Last.Sub(l.First)
}
