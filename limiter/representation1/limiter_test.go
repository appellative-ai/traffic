package representation1

import (
	"fmt"
)

var (
	m = map[string]string{
		RateLimitKey:       "1234",
		RateBurstKey:       "12",
		WindowSizeKey:      "567",
		PeakDurationKey:    "750ms",
		OffPeakDurationKey: "5m",
		ReviewDurationKey:  "5m",
	}
)

func ExampleParseLimiter_Int() {
	var limiter Limiter

	changed := parseLimiter(&limiter, m)
	fmt.Printf("test: parseLimiter() -> %v [changed:%v]\n", limiter, changed)

	m1 := map[string]string{
		RateLimitKey:       "",
		RateBurstKey:       "",
		WindowSizeKey:      "",
		PeakDurationKey:    "",
		OffPeakDurationKey: "",
		ReviewDurationKey:  "",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter() -> %v [changed:%v]\n", limiter, changed)

	// Rate limit
	m1 = map[string]string{
		RateLimitKey: "123",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(RateLimit) -> %v [changed:%v]\n", limiter, changed)
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(RateLimit) -> %v [changed:%v]\n", limiter, changed)
	m1 = map[string]string{
		RateLimitKey: "0",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(RateLimit) -> %v [changed:%v]\n", limiter, changed)

	// Rate burst
	m1 = map[string]string{
		RateBurstKey: "123",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(RateBurst) -> %v [changed:%v]\n", limiter, changed)
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(RateBurst) -> %v [changed:%v]\n", limiter, changed)
	m1 = map[string]string{
		RateBurstKey: "0",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(RateBurst) -> %v [changed:%v]\n", limiter, changed)

	// Window size
	m1 = map[string]string{
		WindowSizeKey: "123",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(WindowSize) -> %v [changed:%v]\n", limiter, changed)
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(WindowSize) -> %v [changed:%v]\n", limiter, changed)

	m1 = map[string]string{
		WindowSizeKey: "0",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(WindowSize) -> %v [changed:%v]\n", limiter, changed)

	//Output:
	//test: parseLimiter() -> {1234 12 567 750ms 5m0s 5m0s} [changed:true]
	//test: parseLimiter() -> {1234 12 567 750ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(RateLimit) -> {123 12 567 750ms 5m0s 5m0s} [changed:true]
	//test: parseLimiter(RateLimit) -> {123 12 567 750ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(RateLimit) -> {123 12 567 750ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(RateBurst) -> {123 123 567 750ms 5m0s 5m0s} [changed:true]
	//test: parseLimiter(RateBurst) -> {123 123 567 750ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(RateBurst) -> {123 123 567 750ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(WindowSize) -> {123 123 123 750ms 5m0s 5m0s} [changed:true]
	//test: parseLimiter(WindowSize) -> {123 123 123 750ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(WindowSize) -> {123 123 123 750ms 5m0s 5m0s} [changed:false]

}

func ExampleParseLimiter_Duration() {
	var limiter Limiter

	changed := parseLimiter(&limiter, m)
	fmt.Printf("test: parseLimiter() -> %v [changed:%v]\n", limiter, changed)

	// Peak duration
	m1 := map[string]string{
		PeakDurationKey: "100ms",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(PeakDuration) -> %v [changed:%v]\n", limiter, changed)
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(PeakDuration) -> %v [changed:%v]\n", limiter, changed)
	m1 = map[string]string{
		PeakDurationKey: "0",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(PeakDuration) -> %v [changed:%v]\n", limiter, changed)

	// Off-peak duration
	m1 = map[string]string{
		OffPeakDurationKey: "2m",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(OffPeakDuration) -> %v [changed:%v]\n", limiter, changed)
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(OffPeakDuration) -> %v [changed:%v]\n", limiter, changed)
	m1 = map[string]string{
		OffPeakDurationKey: "0",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(OffPeakDuration) -> %v [changed:%v]\n", limiter, changed)

	// Review duration
	m1 = map[string]string{
		ReviewDurationKey: "100ms",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(ReviewDuration) -> %v [changed:%v]\n", limiter, changed)
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(ReviewDuration) -> %v [changed:%v]\n", limiter, changed)
	m1 = map[string]string{
		ReviewDurationKey: "0",
	}
	changed = parseLimiter(&limiter, m1)
	fmt.Printf("test: parseLimiter(ReviewDuration) -> %v [changed:%v]\n", limiter, changed)

	//Output:
	//test: parseLimiter() -> {1234 12 567 750ms 5m0s 5m0s} [changed:true]
	//test: parseLimiter(PeakDuration) -> {1234 12 567 100ms 5m0s 5m0s} [changed:true]
	//test: parseLimiter(PeakDuration) -> {1234 12 567 100ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(PeakDuration) -> {1234 12 567 100ms 5m0s 5m0s} [changed:false]
	//test: parseLimiter(OffPeakDuration) -> {1234 12 567 100ms 2m0s 5m0s} [changed:true]
	//test: parseLimiter(OffPeakDuration) -> {1234 12 567 100ms 2m0s 5m0s} [changed:false]
	//test: parseLimiter(OffPeakDuration) -> {1234 12 567 100ms 2m0s 5m0s} [changed:false]
	//test: parseLimiter(ReviewDuration) -> {1234 12 567 100ms 2m0s 100ms} [changed:true]
	//test: parseLimiter(ReviewDuration) -> {1234 12 567 100ms 2m0s 100ms} [changed:false]
	//test: parseLimiter(ReviewDuration) -> {1234 12 567 100ms 2m0s 100ms} [changed:false]
	
}

/*
func _ExampleNewLimiter() {
	resource.NewAgent()

	status := resource.Resolver.AddRepresentation(NamespaceName, "author", messaging.ContentTypeText, "test content")
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := resource.Resolver.Representation(NamespaceName)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	//if buf, ok := ct.Value.([]byte); ok {
	//	fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	//}

	//l := NewLimiter(NamespaceName)
	//	fmt.Printf("test: NewLimiter() -> %v\n", l)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v1 type: application/json value: true] [status:OK]
	//test: Representation() -> [value:125] [status:OK]
	//test: NewLimiter() -> &{false false 1234 12 750ms 5m0s 567 2000}

}


*/
