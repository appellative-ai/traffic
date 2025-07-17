package access

import (
	"strings"
)

const (
	RequestIdHeaderName    = "X-REQUEST-ID"
	FromRouteHeaderName    = "FROM-ROUTE"
	UserAgentHeaderName    = "USER-AGENT"
	ForwardedForHeaderName = "X-FORWARDED-FOR"
)

const (
	OperatorPrefix         = "%"
	RequestReferencePrefix = "%REQ("

	TrafficOperator        = "%TRAFFIC%"      // ingress, egress, ping
	StartTimeOperator      = "%START_TIME%"   // start time
	DurationOperator       = "%DURATION%"     // Total duration in milliseconds of the request from the start time to the last byte out.
	DurationStringOperator = "%DURATION_STR%" // Time package formatted
	RouteOperator          = "%ROUTE%"        // Route name

	OriginAppIdOperator      = "%ORIGIN_APP_ID%"      // origin application id
	OriginRegionOperator     = "%ORIGIN_REGION%"      // origin region
	OriginZoneOperator       = "%ORIGIN_ZONE%"        // origin zone
	OriginSubZoneOperator    = "%ORIGIN_SUB_ZONE%"    // origin sub zone
	OriginHostOperator       = "%ORIGIN_HOST%"        // origin host
	OriginInstanceIdOperator = "%ORIGIN_INSTANCE_ID%" // origin instance id

	RequestProtocolOperator = "%PROTOCOL%" // HTTP Protocol
	RequestMethodOperator   = "%METHOD%"   // HTTP method
	RequestUrlOperator      = "%URL%"
	RequestPathOperator     = "%PATH%"
	RequestHostOperator     = "%HOST%"

	RequestIdOperator           = "%X-REQUEST-ID%"    // X-REQUEST-ID request header value
	RequestFromRouteOperator    = "%FROM-ROUTE%"      // request from route name
	RequestUserAgentOperator    = "%USER-AGENT%"      // user agent request header value
	RequestAuthorityOperator    = "%AUTHORITY%"       // authority request header value
	RequestForwardedForOperator = "%X-FORWARDED-FOR%" // client IP address (X-FORWARDED-FOR request header value)

	ResponseStatusCodeOperator      = "%STATUS_CODE%"      // HTTP status code
	ResponseBytesReceivedOperator   = "%BYTES_RECEIVED%"   // bytes received
	ResponseBytesSentOperator       = "%BYTES_SENT%"       // bytes sent
	ResponseContentEncodingOperator = "%CONTENT-ENCODING%" // content encoding
	ResponseCachedOperator          = "%CACHED%"           // cached flag

	TimeoutDurationOperator = "%TIMEOUT_DURATION%" // threshold timeout
	RateLimitOperator       = "%RATE_LIMIT%"       // threshold rate limit
	RedirectOperator        = "%REDIRECT%"         // threshold redirect percentage

)

var (
	defaultOperators, _ = createOperators([]string{TrafficOperator, StartTimeOperator,
		DurationOperator, RouteOperator,
		RequestMethodOperator, RequestUrlOperator,
		ResponseStatusCodeOperator, ResponseCachedOperator,
		ResponseContentEncodingOperator, ResponseBytesReceivedOperator,
		TimeoutDurationOperator, RateLimitOperator,
		RedirectOperator,
	})
)

func isDirectOperator(op Operator) bool {
	return !strings.HasPrefix(op.Value, OperatorPrefix)
}

func isRequestOperator(op Operator) bool {
	if !strings.HasPrefix(op.Value, RequestReferencePrefix) {
		return false
	}
	if len(op.Value) < (len(RequestReferencePrefix) + 2) {
		return false
	}
	return op.Value[len(op.Value)-2:] == ")%"
}

func requestOperatorHeaderName(v any) (value string) {
	if v == nil {
		return
	}
	if op, ok := v.(Operator); ok {
		if op.Name != "" {
			return op.Name
		}
		value = op.Value
	} else {
		if v2, ok1 := v.(string); ok1 {
			value = v2
		}
	}
	if value != "" {
		if len(value) < (len(RequestReferencePrefix) + 2) {
			return ""
		}
		return value[len(RequestReferencePrefix) : len(value)-2]
	}
	return
}

func requestOperatorHeaderName23(value string) string {
	if len(value) < (len(RequestReferencePrefix) + 2) {
		return ""
	}
	return value[len(RequestReferencePrefix) : len(value)-2]
}

func isStringValue(op Operator) bool {
	switch op.Value {
	case DurationOperator, TimeoutDurationOperator,
		RateLimitOperator, RedirectOperator,
		ResponseStatusCodeOperator, ResponseCachedOperator,
		ResponseBytesSentOperator, ResponseBytesReceivedOperator:
		return false
	}
	return true
}
