package routing

import (
	"fmt"
	"github.com/appellative-ai/collective/notification/notificationtest"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/traffic/routing/representation1"
	"github.com/appellative-ai/traffic/routing/routingtest"
)

func ExampleConfig_Host() {
	a := newAgent(notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> [app:%v] [cache:%v]\n", a.state.Load().AppHost, a.state.Load().CacheHost)

	m := make(map[string]string)
	m[representation1.AppHostKey] = "google.com"
	m[representation1.CacheHostKey] = "localhost:8080"
	a.Message(messaging.NewConfigMessage(m))
	fmt.Printf("test: Message() -> [app:%v] [cache:%v]\n", a.state.Load().AppHost, a.state.Load().CacheHost)

	a.running.Store(true)
	m = make(map[string]string)
	m[representation1.AppHostKey] = "google-reset.com"
	m[representation1.CacheHostKey] = "localhost-reset:8080"
	a.Message(messaging.NewConfigMessage(m))
	fmt.Printf("test: Message() -> [app:%v] [cache:%v]\n", a.state.Load().AppHost, a.state.Load().CacheHost)

	//Output:
	//test: newAgent() -> [app:] [cache:]
	//test: Message() -> [app:google.com] [cache:localhost:8080]
	//test: Message() -> [app:google.com] [cache:localhost:8080]

}

func ExampleConfig_Exchange() {
	a := newAgent(notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> [ex:%v]\n", a.exchange)

	m := messaging.NewConfigMessage(rest.Exchange(routingtest.Exchange))
	a.Message(m)
	fn := rest.Exchange(routingtest.Exchange)
	fmt.Printf("test: Message() -> [ex:%v] %v\n", a.exchange, fn)

	//Output:
	//test: newAgent() -> [ex:0x149f560]
	//test: Message() -> [ex:0x14a2e60] 0x14a2e60

}

func ExampleConfig_Review() {
	a := newAgent(nil)
	fmt.Printf("test: newAgent() -> [dur:%v] [review-dur:%v]\n", a.state.Load().ReviewDuration, a.review.Load().Duration())

	m := map[string]string{
		representation1.ReviewDurationKey: "25",
	}
	a.Message(messaging.NewConfigMessage(m))
	fmt.Printf("test: Message() -> [dur:%v] [review-dur:%v]\n", a.state.Load().ReviewDuration, a.review.Load().Duration())

	//Output:
	//test: newAgent() -> [dur:0s] [review-dur:0s]
	//test: Message() -> [dur:25s] [review-dur:25s]

}
