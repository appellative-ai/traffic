package cache

import (
	"fmt"
	"github.com/appellative-ai/collective/notification/notificationtest"
)

func ExampleNewAgent() {
	a := newAgent(notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> %v\n", a.Name())

	//Output:
	//test: newAgent() -> common:resiliency:agent/cache/request/http
}

/*
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


*/
