package access

import (
	"fmt"
	"github.com/appellative-ai/core/fmtx"
	"github.com/appellative-ai/core/messaging"
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
	Thresholds threshold
	NewReq     *http.Request
	NewResp    *http.Response
	Url        string
	Parsed     *parsed
}

func newEvent(traffic string, start time.Time, duration time.Duration, route string, req any, resp any) *event {
	e := new(event)
	e.Traffic = traffic
	e.Start = start
	e.Duration = duration
	e.Route = route
	e.Req = req
	e.Resp = resp
	e.NewReq = buildRequest(req)
	e.NewResp = buildResponse(resp)
	e.Thresholds = newThreshold(e.NewResp)
	e.Url, e.Parsed = parseURL(e.NewReq.Host, e.NewReq.URL)
	return e
}

func (e *event) addRequest(r *http.Request) {
	e.NewReq = buildRequest(r)
}

func (e *event) addResponse(r *http.Response) {
	e.NewResp = buildResponse(r)
}

func (e *event) value(value string) string {
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
		return messaging.Origin.Region
	case OriginZoneOperator:
		return messaging.Origin.Zone
	case OriginSubZoneOperator:
		return messaging.Origin.SubZone
	case OriginHostOperator:
		return messaging.Origin.Host
	case OriginInstanceIdOperator:
		return messaging.Origin.InstanceId

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
		return encoding(e.NewResp)
	case ResponseCachedOperator:
		return e.Thresholds.cached()

	// Thresholds
	case TimeoutDurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.Thresholds.timeout())) //strconv.Itoa(l.Timeout)
	case RateLimitOperator:
		return fmt.Sprintf("%v", e.Thresholds.rateLimit())
	case RedirectOperator:
		return strconv.Itoa(e.Thresholds.redirect())
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
