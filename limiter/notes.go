package limiter

// Need to capture all traffic for a given time period.
//  1. Need to include rate limited requests
//  2. Need to have access to timeout value to determine when to start rate limiting
//  3. Determine saturation, percentile latency, and gradiant for a window

// Saturation = percentile latency/threshold latency - threshold latency is in the collective
//https://pkg.go.dev/golang.org/x/perf/internal/stats#Sample.Percentile
//

// Gradiant  =
// https://thelinuxcode.com/introduction-to-gonum-enabling-scientific-computing-in-go/
// spend := []float64{8, 12, 15, 20, 22, 35, 60, 90, 150, 250}

//model := stat.LinearRegression(spend, revenue, nil, false)

//fmt.Println(model.Coeff(1)) // Slope = 9.8
//fmt.Print("%0.2f", model.Predicted([]float64{100})[0]) // Predicts $968.32 revenue
// for $100 ad spend
