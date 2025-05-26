package test

import (
	"fmt"
)

func ExampleLoad() {
	Load()
	fmt.Printf("test: LoadTest() -> [%v]\n", "test")

	//Output:
	//test: LoadTest() -> [resiliency:agent/rate-limiting/request/http]
	//test: LoadTest() -> [resiliency:agent/redirect/request/http]
	//test: LoadTest() -> [test]

}
