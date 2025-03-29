package metrics

import "time"

type Profile[T any] struct {
	Week [7][24]T
}

func (p *Profile[T]) Now() T {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return p.Week[day][hour]
}
