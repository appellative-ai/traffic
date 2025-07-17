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
	traffic    string
	start      time.Time
	duration   time.Duration
	route      string
	req        any
	resp       any
	thresholds threshold
	newReq     *http.Request
	newResp    *http.Response
	url        string
	parsed     *parsed
}

func newEvent(traffic string, start time.Time, duration time.Duration, route string, req any, resp any) *event {
	e := new(event)
	e.traffic = traffic
	e.start = start
	e.duration = duration
	e.route = route
	e.req = req
	e.resp = resp
	e.newReq = buildRequest(req)
	e.newResp = buildResponse(resp)
	e.thresholds = newThreshold(e.newResp)
	e.url, e.parsed = parseURL(e.newReq.Host, e.newReq.URL)
	return e
}

func (e *event) addRequest(r *http.Request) {
	e.newReq = buildRequest(r)
}

func (e *event) addResponse(r *http.Response) {
	e.newResp = buildResponse(r)
}

func (e *event) value(value string) string {
	switch value {
	case TrafficOperator:
		return e.traffic
	case StartTimeOperator:
		return fmtx.FmtRFC3339Millis(e.start)
	case DurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.duration))
	case DurationStringOperator:
		return e.duration.String()
	case RouteOperator:
		return e.route

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
		return e.newReq.Method
	case RequestProtocolOperator:
		return e.newReq.Proto
	case RequestPathOperator:
		return e.newReq.URL.Path
	case RequestUrlOperator:
		return e.newReq.URL.String()
	case RequestHostOperator:
		return e.newReq.Host
	case RequestIdOperator:
		return e.newReq.Header.Get(RequestIdHeaderName)
	case RequestFromRouteOperator:
		return e.newReq.Header.Get(FromRouteHeaderName)
	case RequestUserAgentOperator:
		return e.newReq.Header.Get(UserAgentHeaderName)
	case RequestAuthorityOperator:
		return ""
	case RequestForwardedForOperator:
		return e.newReq.Header.Get(ForwardedForHeaderName)

		// Response
	case ResponseBytesReceivedOperator:
		return fmt.Sprintf("%v", e.newResp.ContentLength) //strconv.Itoa(e.NewResp.ContentLength) //l.BytesReceived))
	case ResponseBytesSentOperator:
		return fmt.Sprintf("%v", 0) //l.BytesSent)
	case ResponseStatusCodeOperator:
		if e.newResp == nil {
			return "0"
		} else {
			return strconv.Itoa(e.newResp.StatusCode)
		}
	case ResponseContentEncodingOperator:
		return encoding(e.newResp)
	case ResponseCachedOperator:
		return e.thresholds.cached()

	// Thresholds
	case TimeoutDurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.thresholds.timeout())) //strconv.Itoa(l.Timeout)
	case RateLimitOperator:
		return fmt.Sprintf("%v", e.thresholds.rateLimit())
	case RedirectOperator:
		return strconv.Itoa(e.thresholds.redirect())
	}

	if strings.HasPrefix(value, RequestReferencePrefix) {
		name := requestOperatorHeaderName(value)
		return e.newReq.Header.Get(name)
	}
	if !strings.HasPrefix(value, OperatorPrefix) {
		return value
	}
	return ""
}
