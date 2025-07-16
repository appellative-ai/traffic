package cachetest

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/host"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/iox"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/traffic/cache"
	"github.com/appellative-ai/traffic/cache/representation1"
	"net/http"
	"net/http/httptest"
)

func nextExchange(next rest.Exchange) rest.Exchange {
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

func ExampleExchange() {
	// configure exchange and host name
	//repository.Message(httpx.NewConfigExchangeMessage(Exchange))
	cfg := make(map[string]string)
	cfg[representation1.HostKey] = "localhost:8082"
	exchange.Message(messaging.NewMapMessage(cfg))

	// create request
	url := "https://localhost:8081/search?q=golang"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	httpx.AddRequestId(req)

	// create endpoint and server Http
	e := host.NewEndpoint("pattern", []any{exchange.Agent(cache.NamespaceName), nextExchange})
	r := httptest.NewRecorder()
	e.ServeHTTP(r, req)
	r.Flush()
	buf, err := iox.ReadAll(r.Result().Body, r.Result().Header)
	if err != nil {
		fmt.Printf("test: iox.RedAll() -> [err:%v]\n", err)
	}
	fmt.Printf("test: CacheAgent [status:%v ] [encoding:%v] [buff:%v]\n", r.Result().StatusCode, r.Result().Header.Get(iox.ContentEncoding), len(buf))

	r = httptest.NewRecorder()
	e.ServeHTTP(r, req)
	r.Flush()
	buf, err = iox.ReadAll(r.Result().Body, nil)
	if err != nil {
		fmt.Printf("test: iox.RedAll() -> [err:%v]\n", err)
	}
	fmt.Printf("test: CacheAgent [status:%v ] [encoding:%v] [buff:%v]\n", r.Result().StatusCode, r.Result().Header.Get(iox.ContentEncoding), len(buf))

	//Output:
	//test: CacheAgent [status:200 ] [encoding:] [buff:82654]
	//test: CacheAgent [status:200 ] [encoding:gzip] [buff:41182]

}
