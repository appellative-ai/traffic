package access

import (
	"github.com/appellative-ai/core/iox"
	"net/http"
	"time"
)

func ExampleLog() {
	start := time.Date(2025, 7, 17, 12, 27, 10, 0, time.UTC)

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(contentEncoding, "gzip")
	//resp.Header.Add(CachedName, "true")
	//t := Threshold{Timeout: time.Millisecond * 456, RateLimit: 100, Redirect: 8}
	Log(nil, EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp)

	//Output:
	//{"traffic":"egress","start-time":"2025-07-17T12:27:10.000Z","duration-ms":1500,"route":"test-route","method":"PUT","url":"https://www.google.com/search?q=golang","status-code":418,"cached":null,"encoding":"gzip","bytes-received":12345,"timeout-ms":-1,"rate-limit":-1,"redirect":-1}

}

func ExampleLogWithOperators() {
	start := time.Date(2025, 7, 17, 12, 27, 10, 0, time.UTC)

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(contentEncoding, "gzip")
	SetCached(resp.Header, true)
	SetTimeout(resp.Header, time.Millisecond*456)
	SetRateLimit(resp.Header, 50)
	SetRedirect(resp.Header, 15)

	//resp.Header.Add(CachedName, "true")
	//resp.Header.Add(TimeoutName,"456ms")
	//resp.Header.Add(RateLimitName)
	//t := Threshold{Timeout: time.Millisecond * 456, RateLimit: 100, Redirect: 8}
	Log(nil, EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp)

	//Output:
	//fail

}

const (
	opsPath = "file://[cwd]/accesstest/logging-operators.json"
)

func readFunc() ([]byte, error) {
	return iox.ReadFile(opsPath)
}

/*
func ExampleLoadOperators() {
	err := agent.ConfigureOperators(readFunc)

	fmt.Printf("test: LoadOperators() -> [err:%v] [ops:%v]\n", err, agent.defaultOperators)

	//Output:
	//test: LoadOperators() -> [err:<nil>] [ops:[{start-time %START_TIME%} {duration-ms %DURATION%} {traffic %TRAFFIC%} {route %ROUTE%} {method %METHOD%} {host %HOST%} {path %PATH%} {status-code %STATUS_CODE%} {timeout-ms %TIMEOUT_DURATION%} {rate-limit %RATE_LIMIT%} {redirect %REDIRECT%}]]

}


*/
