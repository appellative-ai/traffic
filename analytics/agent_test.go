package analytics

import (
	"fmt"
	"github.com/behavioral-ai/collective/exchange"
)

func ExampleNew() {
	fmt.Printf("test: New() -> [%v]\n", exchange.Agent(NamespaceName))

	//Output:
	//test: New() -> [resiliency:agent/behavioral-ai/traffic/analytics]

}
