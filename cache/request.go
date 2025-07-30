package cache

import (
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/std"
	"io"
	"net/http"
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

func do(agent *agentT, method string, url string, h http.Header, r io.ReadCloser) (resp *http.Response, status *std.Status) {
	if agent == nil {
		return serverErrorResponse, std.StatusNotFound
	}
	ctx, cancel := httpx.NewContext(nil, agent.state.Timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, r)
	if err != nil {
		return serverErrorResponse, std.NewStatus(std.StatusInvalidArgument, "", err)
	}
	req.Header = h
	resp, err = agent.exchange(req)
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	if err != nil {
		status = std.NewStatus(resp.StatusCode, "", err)
		return
	}
	status = std.StatusOK
	return
}
