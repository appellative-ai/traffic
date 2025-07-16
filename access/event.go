package access

import (
	"fmt"
	"github.com/appellative-ai/core/fmtx"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// event - struct for all logging
type event struct {
	Traffic    string
	Start      time.Time
	Duration   time.Duration
	Route      string
	Req        any
	Resp       any
	Thresholds Threshold
	NewReq     *http.Request
	NewResp    *http.Response
	Url        string
	Parsed     *Parsed
}

func newEvent(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) *event {
	e := new(event)
	e.Traffic = traffic
	e.Start = start
	e.Duration = duration
	e.Route = route
	e.Req = req
	e.Resp = resp
	e.Thresholds = thresholds
	e.NewReq = BuildRequest(req)
	e.NewResp = BuildResponse(resp)
	e.Url, e.Parsed = ParseURL(e.NewReq.Host, e.NewReq.URL)
	return e
}

func (e *event) AddRequest(r *http.Request) {
	e.NewReq = BuildRequest(r)
}

func (e *event) AddResponse(r *http.Response) {
	e.NewResp = BuildResponse(r)
}

func (e *event) Value(value string) string {
	switch value {
	case TrafficOperator:
		return e.Traffic
	case StartTimeOperator:
		return fmtx.FmtRFC3339Millis(e.Start)
	case DurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.Duration))
	case DurationStringOperator:
		return e.Duration.String()
	case RouteOperator:
		return e.Route

		// Origin
	case OriginRegionOperator:
		return agent.origin.Region
	case OriginZoneOperator:
		return agent.origin.Zone
	case OriginSubZoneOperator:
		return agent.origin.SubZone
	case OriginHostOperator:
		return agent.origin.Host
	case OriginInstanceIdOperator:
		return agent.origin.InstanceId

		// Request
	case RequestMethodOperator:
		return e.NewReq.Method
	case RequestProtocolOperator:
		return e.NewReq.Proto
	case RequestPathOperator:
		return e.NewReq.URL.Path
	case RequestUrlOperator:
		return e.NewReq.URL.String()
	case RequestHostOperator:
		return e.NewReq.Host
	case RequestIdOperator:
		return e.NewReq.Header.Get(RequestIdHeaderName)
	case RequestFromRouteOperator:
		return e.NewReq.Header.Get(FromRouteHeaderName)
	case RequestUserAgentOperator:
		return e.NewReq.Header.Get(UserAgentHeaderName)
	case RequestAuthorityOperator:
		return ""
	case RequestForwardedForOperator:
		return e.NewReq.Header.Get(ForwardedForHeaderName)

		// Response
	case ResponseBytesReceivedOperator:
		return fmt.Sprintf("%v", e.NewResp.ContentLength) //strconv.Itoa(e.NewResp.ContentLength) //l.BytesReceived))
	case ResponseBytesSentOperator:
		return fmt.Sprintf("%v", 0) //l.BytesSent)
	case ResponseStatusCodeOperator:
		if e.NewResp == nil {
			return "0"
		} else {
			return strconv.Itoa(e.NewResp.StatusCode)
		}
	case ResponseContentEncodingOperator:
		return Encoding(e.NewResp)
	case ResponseCachedOperator:
		s := e.NewResp.Header.Get(CachedName)
		if s == "" {
			s = "false"
		}
		return s

	// Thresholds
	case TimeoutDurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.Thresholds.TimeoutT())) //strconv.Itoa(l.Timeout)
	case RateLimitOperator:
		return fmt.Sprintf("%v", e.Thresholds.RateLimitT())
	case RedirectOperator:
		return strconv.Itoa(e.Thresholds.RedirectT())
	}
	if strings.HasPrefix(value, RequestReferencePrefix) {
		name := requestOperatorHeaderName(value)
		return e.NewReq.Header.Get(name)
	}
	if !strings.HasPrefix(value, OperatorPrefix) {
		return value
	}
	return ""
}
