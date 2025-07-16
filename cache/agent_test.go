package cache

import (
	"fmt"
	"github.com/appellative-ai/collective/operations/operationstest"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/iox"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/traffic/cache/representation1"
	"net/http"
)

func ExampleNew() {
	//url := "https://www.google.com/search"
	a := newAgent(representation1.Initialize(nil), nil, operationstest.NewService())

	fmt.Printf("test: newAgent() -> %v\n", a.Name())
	m := make(map[string]string)
	m[representation1.HostKey] = "google.com"
	a.Message(messaging.NewMapMessage(m))
	fmt.Printf("test: Message() -> %v\n", a.state.Host)

	//Output:
	//test: newAgent() -> test:resiliency:agent/cache/request/http
	//test: Message() -> google.com

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
