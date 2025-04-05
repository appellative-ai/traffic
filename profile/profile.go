package profile

import "time"

type profile[T any] struct {
	Week [7][24]T
}

func newProfile[T any]() *profile[T] {
	p := new(profile[T])
	return p
}

func (p *profile[T]) Now() T {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return p.Week[day][hour]
}
