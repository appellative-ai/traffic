package cache

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"io"
	"net/http"
	"time"
)

var (
// serverErrorResponse = httpx.NewResponse(http.StatusInternalServerError, nil, nil)
)

/*
type Requester interface {
	Timeout() time.Duration
	Do() rest.Exchange
}

*/

func Do(timeout time.Duration, ex rest.Exchange, method string, url string, h http.Header, r io.ReadCloser) (resp *http.Response, status *messaging.Status) {
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return serverErrorResponse, messaging.NewStatus(messaging.StatusInvalidArgument, err)
	}
	req.Header = h
	resp, err = httpx.ExchangeWithTimeout(timeout, ex)(req)
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	if err != nil {
		status = messaging.NewStatus(resp.StatusCode, err)
		return
	}
	status = messaging.StatusOK()
	return
}
