package cache

import (
	"fmt"
	"github.com/appellative-ai/collective/notification/notificationtest"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/iox"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/traffic/cache/cachetest"
	"github.com/appellative-ai/traffic/cache/representation1"
	"net/http"
)

func ExampleNew() {
	//url := "https://www.google.com/search"
	a := newAgent(notificationtest.NewNotifier())

	fmt.Printf("test: newAgent() -> %v\n", a.Name())
	m := make(map[string]string)
	m[representation1.HostKey] = "google.com"
	a.Message(messaging.NewMapMessage(m))
	fmt.Printf("test: Message() -> %v\n", a.state.Host)

	//Output:
	//test: newAgent() -> common:resiliency:agent/cache/request/http
	//test: Message() -> google.com

}

func ExampleConfig() {
	a := newAgent(notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> %v\n", a.Name())

	m := messaging.NewConfigMessage(rest.Exchange(cachetest.Exchange))
	a.Message(m)
	fn := rest.Exchange(cachetest.Exchange)
	fmt.Printf("test: Message() -> %v %v\n", a.exchange, fn)

	//Output:
	//test: newAgent() -> common:resiliency:agent/cache/request/http
	//test: Message() -> 0x7d4a00 0x7d4a00

}

func routingExchange(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		h := make(http.Header)
		h.Add(iox.AcceptEncoding, iox.GzipEncoding)
		req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
		req.Header = h
		resp, err = httpx.Do(req)
		if err != nil {
			fmt.Printf("test: httx.Do() -> [err:%v]\n", err)
		}
		return
	}
}
