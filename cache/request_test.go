package cache

import (
	"fmt"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/iox"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"net/http"
	"time"
)

type agentTest struct {
	timeout  time.Duration
	exchange rest.Exchange
}

func (a *agentTest) String() string               { return a.Uri() }
func (a *agentTest) Uri() string                  { return "agent:request" }
func (a *agentTest) Message(m *messaging.Message) {}
func (a *agentTest) Timeout() time.Duration       { return a.timeout }
func (a *agentTest) Do() rest.Exchange            { return a.exchange }

func ExampleDo_Get() {
	url := "https://www.google.com/search?q=golang"
	a := newAgent(nil)
	//a.exchange = httpx.Do
	//a.state.Load().Timeout = 0

	h := make(http.Header)
	h.Add(iox.AcceptEncoding, iox.GzipEncoding)
	h.Add(httpx.XRequestId, "1234-request-id")
	resp, status := do(a, http.MethodGet, url, h, nil)
	fmt.Printf("test: do() -> [resp:%v] [status:%v]\n", resp.StatusCode, status)

	if resp.StatusCode == http.StatusOK {
		buf, err1 := iox.ReadAll(resp.Body, resp.Header)
		fmt.Printf("test: iox.ReadAll() -> [buf:%v] [err:%v]\n", len(buf), err1)
	}

	//Output:
	//test: do() -> [resp:200] [status:OK]
	//test: iox.ReadAll() -> [buf:82676] [err:<nil>]

}

func ExampleDo_Get_Timeout() {
	url := "https://www.google.com/search?q=golang"
	a := newAgent(nil)
	a.exchange = httpx.Do
	a.state.Load().TimeoutDuration = time.Millisecond * 10

	h := make(http.Header)
	h.Add(iox.AcceptEncoding, "gzip")
	h.Add(httpx.XRequestId, "1234-request-id")
	resp, status := do(a, http.MethodGet, url, h, nil)
	fmt.Printf("test: do() -> [resp:%v] [status:%v]\n", resp.StatusCode, status)

	if resp.StatusCode == http.StatusOK {
		buf, err1 := iox.ReadAll(resp.Body, resp.Header)
		fmt.Printf("test: iox.ReadAll() -> [buf:%v] [err:%v]\n", len(buf), err1)
	}

	//Output:
	//test: do() -> [resp:504] [status:Timeout - Get "https://www.google.com/search?q=golang": context deadline exceeded]

}
