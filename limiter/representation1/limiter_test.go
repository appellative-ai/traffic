package representation1

import (
	"fmt"
	"github.com/behavioral-ai/collective/resource"
)

const (
	NamespaceName = "resiliency:agent/rate-limiting/request/http"
)

var (
	m = map[string]string{
		RateLimitKey:       "1234",
		RateBurstKey:       "12",
		PeakDurationKey:    "750ms",
		OffPeakDurationKey: "5m",
		LoadSizeKey:        "567",
	}
)

func ExampleParseLimiter() {
	var limiter Limiter
	parseLimiter(&limiter, m)

	fmt.Printf("test: parseLimiter() -> %v\n", limiter)

	//Output:
	//test: parseLimiter() -> {false false 1234 12 750ms 5m0s 567}

}

func _ExampleNewLimiter() {
	resource.NewAgent()

	status := resource.Resolver.AddRepresentation(NamespaceName, Fragment, "author", m)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := resource.Resolver.Representation(NamespaceName, Fragment)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	l := NewLimiter(NamespaceName)
	fmt.Printf("test: NewLimiter() -> %v\n", l)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v1 type: application/json value: true] [status:OK]
	//test: Representation() -> [value:125] [status:OK]
	//test: NewLimiter() -> &{false false 1234 12 750ms 5m0s 567 2000}

}
