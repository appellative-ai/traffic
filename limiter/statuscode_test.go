package limiter

import (
	"fmt"
)

func ExampleStatusCodeSample() {
	s := new(StatusCodeSample)

	s.Update(&event{StatusCode: 100})
	s.Update(&event{StatusCode: 199})

	s.Update(&event{StatusCode: 200})
	s.Update(&event{StatusCode: 201})
	s.Update(&event{StatusCode: 222})
	s.Update(&event{StatusCode: 229})
	s.Update(&event{StatusCode: 299})

	s.Update(&event{StatusCode: 300})
	s.Update(&event{StatusCode: 399})

	s.Update(&event{StatusCode: 400})
	s.Update(&event{StatusCode: 401})
	s.Update(&event{StatusCode: 402})
	s.Update(&event{StatusCode: 429})
	s.Update(&event{StatusCode: 499})

	s.Update(&event{StatusCode: 500})
	s.Update(&event{StatusCode: 501})
	s.Update(&event{StatusCode: 502})
	s.Update(&event{StatusCode: 504})
	s.Update(&event{StatusCode: 599})

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
