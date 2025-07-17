package access

import (
	"net/http"
	"strings"
)

// Request - request interface for non HTTP traffic
type Request interface {
	Url() string
	Header() http.Header
	Method() string
	Protocol() string
}

// Response - response interface for non HTTP traffic
type Response interface {
	StatusCode() int
	Header() http.Header
}

// RequestImpl - non HTTP request attributes
type RequestImpl struct {
	Url      string
	Header   http.Header
	Method   string
	Protocol string
}

func BuildRequest(r any) *http.Request {
	if r == nil {
		newReq, _ := http.NewRequest("", failsafeUri, nil)
		return newReq
	}
	if req, ok := r.(*http.Request); ok {
		return req
	}
	if req, ok := r.(Request); ok {
		newReq, _ := http.NewRequest(req.Method(), req.Url(), nil)
		newReq.Header = req.Header()
		newReq.Proto = req.Protocol()
		return newReq
	}
	newReq, _ := http.NewRequest("", "https://somehost.com/search?q=test", nil)
	return newReq
}

func BuildResponse(r any) *http.Response {
	if r == nil {
		newResp := &http.Response{StatusCode: http.StatusOK}
		newResp.Header = make(http.Header)
		return newResp
	}
	if newResp, ok := r.(*http.Response); ok {
		if newResp.Header == nil {
			newResp.Header = make(http.Header)
		}
		return newResp
	}
	if resp, ok := r.(Response); ok {
		return &http.Response{StatusCode: resp.StatusCode(), Header: resp.Header()}
	}
	if sc, ok := r.(int); ok {
		return &http.Response{StatusCode: sc, Header: make(http.Header)}
	}
	if status, ok := r.(int); ok {
		return &http.Response{StatusCode: status, Header: make(http.Header)}
	}
	newResp := &http.Response{StatusCode: http.StatusOK, Header: make(http.Header)}
	return newResp
}

func Encoding(resp *http.Response) string {
	encoding := ""
	if resp != nil && resp.Header != nil {
		encoding = resp.Header.Get(contentEncoding)
	}
	// normalize encoding
	if strings.Contains(strings.ToLower(encoding), "none") {
		encoding = ""
	}
	return encoding
}
