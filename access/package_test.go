package access

import (
	"fmt"
	"net/http"
	"time"
)

func ExampleSet() {
	h := make(http.Header)
	SetTimeout(h, time.Millisecond*456)
	SetRateLimit(h, float64(123))
	SetRedirect(h, 35)
	SetCached(h, false)

	fmt.Printf("test: Set() -> %v\n", h)
	RemoveThresholds(h)
	fmt.Printf("test: Set() -> %v\n", h)

	//Output:
	//test: Set() -> map[X-Threshold:[timeout=456ms rate-limit=123 redirect=35 cached=false]]
	//test: Set() -> map[]

}

func ExampleLog() {
	start := time.Date(2025, 7, 17, 12, 27, 10, 0, time.UTC)

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(contentEncoding, "gzip")
	SetCached(resp.Header, true)
	SetTimeout(resp.Header, time.Millisecond*456)
	SetRateLimit(resp.Header, 50)
	SetRedirect(resp.Header, 15)

	Log(nil, EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp)

	//Output:
	//{"traffic":"egress","start-time":"2025-07-17T12:27:10.000Z","duration-ms":1500,"route":"test-route","method":"PUT","url":"https://www.google.com/search?q=golang","status-code":418,"cached":null,"encoding":"gzip","bytes-received":12345,"timeout-ms":-1,"rate-limit":-1,"redirect":-1}

}
