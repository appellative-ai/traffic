package metrics

import (
	"fmt"
)

func ExampleStatusCodeSample() {
	s := new(StatusCodeSample)

	s.Update(&Event{StatusCode: 100})
	s.Update(&Event{StatusCode: 199})

	s.Update(&Event{StatusCode: 200})
	s.Update(&Event{StatusCode: 201})
	s.Update(&Event{StatusCode: 222})
	s.Update(&Event{StatusCode: 229})
	s.Update(&Event{StatusCode: 299})

	s.Update(&Event{StatusCode: 300})
	s.Update(&Event{StatusCode: 399})

	s.Update(&Event{StatusCode: 400})
	s.Update(&Event{StatusCode: 401})
	s.Update(&Event{StatusCode: 402})
	s.Update(&Event{StatusCode: 429})
	s.Update(&Event{StatusCode: 499})

	s.Update(&Event{StatusCode: 500})
	s.Update(&Event{StatusCode: 501})
	s.Update(&Event{StatusCode: 502})
	s.Update(&Event{StatusCode: 504})
	s.Update(&Event{StatusCode: 599})

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
