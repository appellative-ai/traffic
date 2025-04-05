package profile

import (
	"fmt"
	"github.com/behavioral-ai/collective/eventing/eventtest"
)

func ExampleNewAgent() {
	a := newAgent(eventtest.New())

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [resiliency:agent/behavioral-ai/traffic/profile]

}
