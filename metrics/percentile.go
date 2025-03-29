package metrics

type Percentile struct {
	Value float64
}

func NewPercentileProfile() *Profile[Percentile] {
	p := new(Profile[Percentile])
	return p
}
