package representation1

import (
	"fmt"
)

var (
	m = map[string]string{
		AppHostKey:         "localhost:8080",
		CacheHostKey:       "localhost:8081",
		TimeoutDurationKey: "750ms",
		ReviewDurationKey:  "45m",
	}
)

func ExampleParseRouting() {
	var routing Routing
	parseRouting(&routing, m)

	fmt.Printf("test: parseRouting() -> %v\n", routing)

	//Output:
	//test: parseRouting() -> {localhost:8080 localhost:8081 750ms 45m0s}
	
}
