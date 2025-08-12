package representation1

import (
	"fmt"
)

var (
	m2 = map[string]string{
		RateLimitKey:    "1234",
		RateBurstKey:    "12",
		IntervalKey:     "750ms",
		OriginalPathKey: "/google/search:v1",
		NewPathKey:      "/yahoo/search:v2",
	}
)

func ExampleParseRedirect() {
	var redirect Redirect2
	parseRedirect(&redirect, m2)

	fmt.Printf("test: parseRedirect() -> %v\n", redirect)

	//Output:
	//test: parseRedirect() -> {false 1234 12 750ms /google/search:v1 /yahoo/search:v2 0 0 <nil> <nil>}

}
