package access

import (
	"net/http"
	"time"
)

func ExampleLogConfigure() {
	start := time.Date(2025, 7, 17, 12, 27, 10, 0, time.UTC)

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(contentEncoding, "gzip")
	Log(nil, EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp)

	a := newAgent()
	ops, _ := createOperators([]string{TrafficOperator, StartTimeOperator,
		DurationOperator, RouteOperator,
		RequestMethodOperator, RequestUrlOperator,
	})
	a.Message(NewOperatorsMessage(ops))
	a.log(start, time.Millisecond*1500, req, resp)

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
