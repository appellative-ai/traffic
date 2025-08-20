package cache

import (
	"github.com/appellative-ai/agency/logger"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/std"
	"io"
	"net/http"
	"time"
)

func do(a *agentT, method string, url string, h http.Header, r io.ReadCloser) (resp *http.Response, status *std.Status) {
	if a == nil {
		return serverErrorResponse, std.StatusNotFound
	}
	ctx, cancel := httpx.NewContext(nil, a.state.Load().TimeoutDuration)
	defer cancel()
	start := time.Now().UTC()
	req, err := http.NewRequestWithContext(ctx, method, url, r)
	if err != nil {
		return serverErrorResponse, std.NewStatus(std.StatusInvalidArgument, err)
	}
	req.Header = h
	resp, err = a.exchange(req)
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	logger.Agent.LogEgress(start, time.Since(start), cacheRouteName, req, resp, a.state.Load().TimeoutDuration)
	if err != nil {
		status = std.NewStatus(resp.StatusCode, err)
		return
	}
	if a.state.Load().TimeoutDuration > 0 {
		err = httpx.TransformBody(resp)
	}
	if err != nil {
		status = std.NewStatus(resp.StatusCode, err)
	} else {
		status = std.StatusOK
	}
	return
}
