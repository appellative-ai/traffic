package limiter

import (
	"fmt"
	"time"
)

func ExampleRegressionSample() {
	s := new(RegressionSample)

	s.Update(&event{Duration: time.Second * 2})
	fmt.Printf("test: Update()  -> x:%v y:%v\n", s.X, s.Y)

	s.Update(&event{Duration: time.Millisecond * 1500})
	fmt.Printf("test: Update()  -> x:%v y:%v\n", s.X, s.Y)

	s.Update(&event{Duration: 0})
	fmt.Printf("test: Update()  -> x:%v y:%v\n", s.X, s.Y)

	//Output:
	//test: Update()  -> x:[] y:[2e+09]
	//test: Update()  -> x:[] y:[2e+09 1.5e+09]
	//test: Update()  -> x:[] y:[2e+09 1.5e+09 0]

}
