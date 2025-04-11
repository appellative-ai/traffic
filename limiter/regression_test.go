package limiter

import (
	"fmt"
	"time"
)

func ExampleRegressionSample() {
	s := new(regressionSample)

	s.update(&event{duration: time.Second * 2})
	fmt.Printf("test: update()  -> x:%v y:%v\n", s.x, s.y)

	s.update(&event{duration: time.Millisecond * 1500})
	fmt.Printf("test: update()  -> x:%v y:%v\n", s.x, s.y)

	s.update(&event{duration: 0})
	fmt.Printf("test: update()  -> x:%v y:%v\n", s.x, s.y)

	//Output:
	//test: Update()  -> x:[] y:[2e+09]
	//test: Update()  -> x:[] y:[2e+09 1.5e+09]
	//test: Update()  -> x:[] y:[2e+09 1.5e+09 0]

}
