package representation1

import (
	"encoding/json"
	"fmt"
)

var (
	m = map[string]string{
		AppHostKey:   "localhost:8080",
		CacheHostKey: "localhost:8081",
		TimeoutKey:   "750ms",
		IntervalKey:  "5m",
	}

	r1 = Route{
		Name: "test-route-one",
		Path: "/resource/test2",
		Redirect: Redirect{
			Path:                "/resource/test/main",
			StatusCodes:         []string{"5xx"},
			StatusCodeThreshold: 10,
			Percentile:          "99/1500ms",
			PercentileThreshold: 10,
		},
	}

	r2 = Route{
		Name: "test-route-two",
		Path: "/resource/old",
		Redirect: Redirect{
			Path:                "/resource/test/new",
			StatusCodes:         []string{"5xx", "4xx"},
			StatusCodeThreshold: 15,
			Percentile:          "95/2000ms",
			PercentileThreshold: 15,
		},
	}
)

func ExampleParseRouting() {
	var routing Routing
	parseRouting(&routing, m)

	fmt.Printf("test: parseRouting() -> %v\n", routing)

	//Output:
	//test: parseRouting() -> {false false true localhost:8080 localhost:8081 app2 5m0s 750ms}

}

func ExampleRoutingTable() {
	rt := RoutingTable{Routes: []*Route{&r1, &r2}}

	fmt.Printf("test: Route() -> %v\n", rt)

	buf, err := json.Marshal(rt)
	fmt.Printf("test: json.Marshal() -> %v [err:%v]\n", string(buf), err)

	//Output:
	//test: Route() -> {test-route /resource/test2 {/resource/test/main [5xx] 10 99/1500ms 10 <nil> <nil>}}

}
