package metrics

import (
	"fmt"
	"github.com/behavioral-ai/collective/timeseries"
)

func ExampleStatusCodeSample() {
	s := new(StatusCodeSample)

	s.Update(&timeseries.Event{StatusCode: 100})
	s.Update(&timeseries.Event{StatusCode: 199})

	s.Update(&timeseries.Event{StatusCode: 200})
	s.Update(&timeseries.Event{StatusCode: 201})
	s.Update(&timeseries.Event{StatusCode: 222})
	s.Update(&timeseries.Event{StatusCode: 229})
	s.Update(&timeseries.Event{StatusCode: 299})

	s.Update(&timeseries.Event{StatusCode: 300})
	s.Update(&timeseries.Event{StatusCode: 399})

	s.Update(&timeseries.Event{StatusCode: 400})
	s.Update(&timeseries.Event{StatusCode: 401})
	s.Update(&timeseries.Event{StatusCode: 402})
	s.Update(&timeseries.Event{StatusCode: 429})
	s.Update(&timeseries.Event{StatusCode: 499})

	s.Update(&timeseries.Event{StatusCode: 500})
	s.Update(&timeseries.Event{StatusCode: 501})
	s.Update(&timeseries.Event{StatusCode: 502})
	s.Update(&timeseries.Event{StatusCode: 504})
	s.Update(&timeseries.Event{StatusCode: 599})

	fmt.Printf("test: Update()  -> [2xx:%v]\n", s.Status2xx)
	fmt.Printf("test: Update()  -> [4xx:%v]\n", s.Status4xx)
	fmt.Printf("test: Update()  -> [429:%v]\n", s.Status429)
	fmt.Printf("test: Update()  -> [504:%v]\n", s.Status504)
	fmt.Printf("test: Update()  -> [5xx:%v]\n", s.Status5xx)

	//Output:
	//test: Update()  -> [2xx:5]
	//test: Update()  -> [4xx:4]
	//test: Update()  -> [429:1]
	//test: Update()  -> [504:1]
	//test: Update()  -> [5xx:4]

}
