package representation1

import (
	"fmt"
	"github.com/behavioral-ai/collective/resource"
)

const (
	NamespaceName = "test:resiliency:agent/redirect/request/http"
)

var (
	m = map[string]string{
		RateLimitKey:    "1234",
		RateBurstKey:    "12",
		IntervalKey:     "750ms",
		OriginalPathKey: "/google/search:v1",
		NewPathKey:      "/yahoo/search:v2",
	}
)

func ExampleParseRedirect() {
	var redirect Redirect
	parseRedirect(&redirect, m)

	fmt.Printf("test: parseRedirect() -> %v\n", redirect)

	//Output:
	//test: parseRedirect() -> {false 1234 12 750ms /google/search:v1 /yahoo/search:v2 0 0 <nil> <nil>}

}

func _ExampleNewRedirect() {
	resource.NewAgent()

	status := resource.Resolver.AddRepresentation(NamespaceName, Fragment, "author", m)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := resource.Resolver.Representation(NamespaceName)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	//r := NewRedirect(NamespaceName)
	//fmt.Printf("test: NewRedirect() -> %v\n", r)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v1 type: application/json value: true] [status:OK]
	//test: Representation() -> [value:128] [status:OK]

}
