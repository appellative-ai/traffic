package metrics

import (
	"fmt"
	"time"
)

func Example_RPS() {
	rps := requestsSecond(time.Second*2, 234)
	fmt.Printf("test: requestSecond() -> [%v]\n", rps)

	rps = requestsSecond(0, -1)
	fmt.Printf("test: requestSecond() -> [%v]\n", rps)

	rps = requestsSecond(time.Second, -1)
	fmt.Printf("test: requestSecond() -> [%v]\n", rps)

	rps = requestsSecond(time.Millisecond, 12345)
	fmt.Printf("test: requestSecond() -> [%v]\n", rps)

	//Output:
	//test: requestSecond() -> [117]
	//test: requestSecond() -> [-1]
	//test: requestSecond() -> [0]
	//test: requestSecond() -> [12345000]

}
