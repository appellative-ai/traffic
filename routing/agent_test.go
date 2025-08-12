package routing

import (
	"fmt"
	"github.com/appellative-ai/collective/notification/notificationtest"
)

func ExampleNewAgent() {
	a := newAgent(notificationtest.NewNotifier())

	fmt.Printf("test: newAgent() -> %v\n", a.Name())

	//Output:
	//test: newAgent() -> common:resiliency:agent/routing/request/http

}
