package logger

import (
	"github.com/appellative-ai/core/logx"
	"github.com/appellative-ai/core/messaging"
	"net/http"
	"time"
)

const (
	contentEncoding = "Content-Encoding"
)

func ExampleLogConfigure() {
	a := new(agentT)
	start := time.Date(2025, 7, 17, 12, 27, 10, 0, time.UTC)

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(contentEncoding, "gzip")
	a.LogEgress(nil, start, time.Millisecond*1500, "test-route", req, resp, 0)

	//a := newAgent()
	ops, _ := logx.CreateOperators([]string{logx.TrafficOperator, logx.StartTimeOperator,
		logx.DurationOperator, logx.RouteOperator,
		logx.RequestMethodOperator, logx.RequestUrlOperator,
	})
	a.Message(messaging.NewConfigMessage(ops))
	a.LogEgress(nil, start, time.Millisecond*1500, "test-route", req, resp, 0)

	//Output:
	//fail

}

/*
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


*/
