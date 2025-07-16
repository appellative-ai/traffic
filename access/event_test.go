package access

import (
	"fmt"
	"net/http"
	"time"
)

func _ExampleValue_Duration() {
	start := time.Now()

	time.Sleep(time.Second * 2)
	data := &event{}
	data.Duration = time.Since(start)
	fmt.Printf("test: Value(\"Duration\") -> [%v]\n", data.Value(DurationOperator))
	fmt.Printf("test: Value(\"DurationString\") -> [%v]\n", data.Value(DurationStringOperator))

	//Output:
	//test: Value("Duration") -> [2011]
	//test: Value("DurationString") -> [2.0117949s]

}

func ExampleValue_Origin() {
	agent.SetOrigin("region", "zone", "subZone", "host-name", "instanceId")

	data := event{}
	fmt.Printf("test: Value(\"%v\") -> [%v]\n", "region", data.Value(OriginRegionOperator))
	fmt.Printf("test: Value(\"%v\") -> [%v]\n", "zone", data.Value(OriginZoneOperator))
	fmt.Printf("test: Value(\"%v\") -> [%v]\n", "sub-zone", data.Value(OriginSubZoneOperator))
	fmt.Printf("test: Value(\"%v\") -> [%v]\n", "service", data.Value(OriginHostOperator))
	fmt.Printf("test: Value(\"%v\") -> [%v]\n", "instance-id", data.Value(OriginInstanceIdOperator))

	//Output:
	//test: Value("region") -> [region]
	//test: Value("zone") -> [zone]
	//test: Value("sub-zone") -> [subZone]
	//test: Value("service") -> [host-name]
	//test: Value("instance-id") -> [instanceId]

}

func ExampleValue_Thresholds() {
	name := "test-route"
	start := time.Now().UTC()

	//data := Event{}
	//fmt.Printf("test: Value(\"%v\") -> [%v]\n", name, data.Value(op))
	//data = Event{ControllerName: name}
	//fmt.Printf("test: Value(\"%v\") -> [controller:%v]\n", name, data.Value(op))

	data1 := newEvent(EgressTraffic, start, time.Since(start), "test-route", nil, nil, Threshold{
		Timeout:   time.Millisecond * 1500,
		RateLimit: 125,
		Redirect:  10})
	fmt.Printf("test: Value(\"%v\") -> [traffic:%v]\n", name, data1.Value(TrafficOperator))
	fmt.Printf("test: Value(\"%v\") -> [route:%v]\n", name, data1.Value(RouteOperator))
	fmt.Printf("test: Value(\"%v\") -> [timeout:%v]\n", name, data1.Value(TimeoutDurationOperator))
	fmt.Printf("test: Value(\"%v\") -> [limit:%v]\n", name, data1.Value(RateLimitOperator))
	fmt.Printf("test: Value(\"%v\") -> [redirect:%v]\n", name, data1.Value(RedirectOperator))

	//data = Event{Timeout: 500}
	//fmt.Printf("test: Value(\"%v\") -> [timeout:%v]\n", name, data.Value(TimeoutDurationOperator))

	//Output:
	//test: Value("test-route") -> [traffic:egress]
	//test: Value("test-route") -> [route:test-route]
	//test: Value("test-route") -> [timeout:1500]
	//test: Value("test-route") -> [limit:125]
	//test: Value("test-route") -> [redirect:10]

}

func ExampleValue_Request() {
	op := RequestMethodOperator

	data := &event{}
	//fmt.Printf("test: Value(\"method\") -> [%v]\n", data.Value(op))

	req, _ := http.NewRequest("POST", "www.google.com", nil)
	req.Header.Add(RequestIdHeaderName, "123-456-789")
	req.Header.Add(FromRouteHeaderName, "calling-route")
	data = &event{}
	data.AddRequest(req)
	fmt.Printf("test: Value(\"method\") -> [%v]\n", data.Value(op))

	fmt.Printf("test: Value(\"headers\") -> [request-id:%v] [from-route:%v]\n", data.Value(RequestIdOperator), data.Value(RequestFromRouteOperator))

	//Output:
	//test: Value("method") -> [POST]
	//test: Value("headers") -> [request-id:123-456-789] [from-route:calling-route]
}

func ExampleValue_Response() {
	op := ResponseStatusCodeOperator

	data := &event{}
	fmt.Printf("test: Value(\"code\") -> [%v]\n", data.Value(op))

	resp := &http.Response{StatusCode: 200}
	data = &event{}
	data.AddResponse(resp)
	fmt.Printf("test: Value(\"code\") -> [%v]\n", data.Value(op))

	//Output:
	//test: Value("code") -> [0]
	//test: Value("code") -> [200]
}

func ExampleValue_Request_Header() {
	req, _ := http.NewRequest("", "www.google.com", nil)
	req.Header.Add("customer", "Ted's Bait & Tackle")
	data := event{}
	data.AddRequest(req)
	fmt.Printf("test: Value(\"customer\") -> [%v]\n", data.Value("%REQ(customer)%"))

	//Output:
	//test: Value("customer") -> [Ted's Bait & Tackle]
}
