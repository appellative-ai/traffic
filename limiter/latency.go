package limiter

import (
	"time"
)

type LatencySample struct {
	First time.Time
	Last  time.Time
}

func (l *LatencySample) Update(event *event) {
	if l.First.IsZero() {
		l.First = event.Start
	} else {
		if event.Start.Compare(l.First) == -1 {
			l.First = event.Start
		}
	}
	if l.Last.IsZero() {
		l.Last = event.Start.Add(event.Duration)
	} else {
		if event.Start.Compare(l.Last) == 1 {
			l.Last = event.Start.Add(event.Duration)
		}
	}
}

func (l *LatencySample) Elapsed() time.Duration {
	return l.Last.Sub(l.First)
}
