package cache

import (
	"fmt"
	"github.com/appellative-ai/collective/notification/notificationtest"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/traffic/cache/cachetest"
	"github.com/appellative-ai/traffic/cache/representation1"
)

func ExampleConfig_Host() {
	a := newAgent(notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> [host:%v]\n", a.state.Load().Host)

	m := make(map[string]string)
	m[representation1.HostKey] = "google.com"
	a.Message(messaging.NewConfigMessage(m))
	fmt.Printf("test: Message() -> [host:%v]\n", a.state.Load().Host)

	//Output:
	//test: newAgent() -> [host:]
	//test: Message() -> [host:google.com]

}

func ExampleConfig_Exchange() {
	a := newAgent(notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> [ex:%v]\n", a.exchange)

	m := messaging.NewConfigMessage(rest.Exchange(cachetest.Exchange))
	a.Message(m)
	fn := rest.Exchange(cachetest.Exchange)
	fmt.Printf("test: Message() -> [ex:%v] %v\n", a.exchange, fn)

	//Output:
	//test: newAgent() -> [ex:0x11882a0]
	//test: Message() -> [ex:0x118bce0] 0x118bce0

}
