package timeseries

/*
type percentile struct {
	Score   int // k-th percentile
	Latency int // milliseconds
}

*/

// Interface -
type Interface struct {
	LinearRegression func(x, y, weights []float64, origin bool) (alpha, beta float64)
	Percentile       func(x, weights []float64, sorted bool, pctile float64) float64
}

// Functions -
var Functions = func() *Interface {
	return &Interface{
		LinearRegression: func(x, y, weights []float64, origin bool) (alpha, beta float64) {
			return
		},
		Percentile: func(x, weights []float64, sorted bool, pctile float64) float64 {

			return 0.0
		},
	}
}()

//
// https://thelinuxcode.com/introduction-to-gonum-enabling-scientific-computing-in-go/
//
// Linear regression
//model := stat.LinearRegression(spend, revenue, nil, false)
//
//fmt.Println(model.Coeff(1)) // Slope = 9.8
//fmt.Print("%0.2f", model.Predicted([]float64{100})[0]) // Predicts $968.32 revenue
// for $100 ad spend

type Sample struct {
	// Xs is the slice of sample values.
	Xs []float64

	// Weights[i] is the weight of sample Xs[i].  If Weights is
	// nil, all Xs have weight 1.  Weights must have the same
	// length of Xs and all values must be non-negative.
	Weights []float64

	// Sorted indicates that Xs is sorted in ascending order.
	Sorted bool
}
