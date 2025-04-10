package limiter

import (
	"fmt"
	"time"
)

func ExamplePercentileSample() {
	s := new(PercentileSample)

	s.Update(&event{Duration: time.Second * 2})
	fmt.Printf("test: Update()  -> x:%v\n", s.X)

	s.Update(&event{Duration: time.Millisecond * 1500})
	fmt.Printf("test: Update()  -> x:%v\n", s.X)

	s.Update(&event{Duration: 0})
	fmt.Printf("test: Update()  -> x:%v\n", s.X)

	//Output:
	//test: Update()  -> x:[2e+09]
	//test: Update()  -> x:[2e+09 1.5e+09]
	//test: Update()  -> x:[2e+09 1.5e+09 0]

}
