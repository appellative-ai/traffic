package routing

import (
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/std"
	"io"
	"net/http"
)

var (
// serverErrorResponse = httpx.NewResponse(http.StatusInternalServerError, nil, nil)
)

func do(a *agentT, method string, url string, h http.Header, r io.ReadCloser) (resp *http.Response, status *std.Status) {
	if a == nil {
		return serverErrorResponse, std.StatusNotFound
	}
	ctx, cancel := httpx.NewContext(nil, a.state.Timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, r)
	if err != nil {
		return serverErrorResponse, std.NewStatus(std.StatusInvalidArgument, "", err)
	}
	req.Header = h
	resp, err = a.exchange(req)
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	if err != nil {
		status = std.NewStatus(resp.StatusCode, "", err)
		return
	}
	if a.state.Timeout > 0 {
		err = httpx.TransformBody(resp)
	}
	if err != nil {
		status = std.NewStatus(resp.StatusCode, "", err)
	} else {
		status = std.StatusOK
	}

	return
}
