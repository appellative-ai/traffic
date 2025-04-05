package metrics

import (
	"fmt"
	"github.com/behavioral-ai/collective/timeseries"
	"time"
)

func ExamplePercentileSample() {
	s := new(PercentileSample)

	s.Update(&timeseries.Event{Duration: time.Second * 2})
	fmt.Printf("test: Update()  -> x:%v\n", s.X)

	s.Update(&timeseries.Event{Duration: time.Millisecond * 1500})
	fmt.Printf("test: Update()  -> x:%v\n", s.X)

	s.Update(&timeseries.Event{Duration: 0})
	fmt.Printf("test: Update()  -> x:%v\n", s.X)

	//Output:
	//test: Update()  -> x:[2e+09]
	//test: Update()  -> x:[2e+09 1.5e+09]
	//test: Update()  -> x:[2e+09 1.5e+09 0]

}
