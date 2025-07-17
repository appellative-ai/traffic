package access

import (
	"fmt"
	"time"
)

func ExampleThresholdTimeout() {
	s := fmt.Sprintf("%v", time.Millisecond*100)
	t := Threshold{Timeout: s}

	v := t.timeout()
	fmt.Printf("test: Threshold-String() -> [s:%v] [value:%v]\n", s, v)

	t = Threshold{Timeout: time.Second * 3}
	v = t.timeout()
	fmt.Printf("test: Threshold-Duration() ->  [value:%v]\n", v)

	t = Threshold{Timeout: 100}
	v = t.timeout()
	fmt.Printf("test: Threshold-Int() ->  [value:%v]\n", v)

	//Output:
	//test: Threshold-String() -> [s:100ms] [value:100ms]
	//test: Threshold-Duration() ->  [value:3s]
	//test: Threshold-Int() ->  [value:-1ns]

}

func ExampleThresholdRedirect() {
	s := "66"
	t := Threshold{Redirect: s}

	v := t.redirect()
	fmt.Printf("test: Redirect-String() -> [s:%v] [value:%v]\n", s, v)

	t = Threshold{Redirect: 33}
	v = t.redirect()
	fmt.Printf("test: Redirect-Int() ->  [value:%v]\n", v)

	t = Threshold{Redirect: time.Second * 4}
	v = t.redirect()
	fmt.Printf("test: Redirect-Duration() ->  [value:%v]\n", v)

	//Output:
	//test: Redirect-String() -> [s:66] [value:66]
	//test: Redirect-Int() ->  [value:33]
	//test: Redirect-Duration() ->  [value:-1]

}

func ExampleThresholdRateLimit() {
	s := "77"
	t := Threshold{RateLimit: s}

	v := t.rateLimit()
	fmt.Printf("test: RateLimit-String() -> [s:%v] [value:%v]\n", s, v)

	t = Threshold{RateLimit: float64(123)}
	v = t.rateLimit()
	fmt.Printf("test: RateLimit-Float64() ->  [value:%v]\n", v)

	t = Threshold{RateLimit: time.Second * 4}
	v = t.rateLimit()
	fmt.Printf("test: RateLimit-Duration() ->  [value:%v]\n", v)

	//Output:
	//test: RateLimit-String() -> [s:77] [value:77]
	//test: RateLimit-Float64() ->  [value:123]
	//test: RateLimit-Duration() ->  [value:-1]

}

func ExampleThresholdCached() {
	s := "true"
	t := Threshold{Cached: s}

	v := t.cached()
	fmt.Printf("test: Cached-String() -> [s:%v] [value:%v]\n", s, v)

	s = ""
	t = Threshold{}
	v = t.cached()
	fmt.Printf("test: Cached-String() -> [s:%v] [value:%v]\n", s, v)

	s = "false"
	t = Threshold{Cached: s}
	v = t.cached()
	fmt.Printf("test: Cached-String() -> [s:%v] [value:%v]\n", s, v)

	//Output:
	//test: Cached-String() -> [s:true] [value:true]
	//test: Cached-String() -> [s:] [value:false]
	//test: Cached-String() -> [s:false] [value:false]

}
