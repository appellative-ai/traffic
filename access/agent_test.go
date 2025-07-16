package access

import (
	"fmt"
	"github.com/appellative-ai/core/iox"
	"net/http"
	"time"
)

func ExampleLog() {
	start := time.Now().UTC()

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(ContentEncoding, "gzip")
	resp.Header.Add(CachedName, "true")
	t := Threshold{Timeout: time.Millisecond * 456, RateLimit: 100, Redirect: 8}
	agent.Log(EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp, t)

	//Output:
	//fail

}

func ExampleLogWithOperators() {
	start := time.Now().UTC()

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(ContentEncoding, "gzip")
	resp.Header.Add(CachedName, "true")
	t := Threshold{Timeout: time.Millisecond * 456, RateLimit: 100, Redirect: 8}
	LogWithOperators(nil, EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp, t)

	//Output:
	//fail

}

const (
	opsPath = "file://[cwd]/accesstest/logging-operators.json"
)

func readFunc() ([]byte, error) {
	return iox.ReadFile(opsPath)
}

func ExampleLoadOperators() {
	err := agent.ConfigureOperators(readFunc)

	fmt.Printf("test: LoadOperators() -> [err:%v] [ops:%v]\n", err, agent.defaultOperators)

	//Output:
	//test: LoadOperators() -> [err:<nil>] [ops:[{start-time %START_TIME%} {duration-ms %DURATION%} {traffic %TRAFFIC%} {route %ROUTE%} {method %METHOD%} {host %HOST%} {path %PATH%} {status-code %STATUS_CODE%} {timeout-ms %TIMEOUT_DURATION%} {rate-limit %RATE_LIMIT%} {redirect %REDIRECT%}]]

}
