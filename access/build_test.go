package access

import "fmt"

func ExampleBuildRequest() {
	r := buildRequest(nil)

	fmt.Printf("test: buildRequest() -> %v\n", r)

	//Output:
	//fail
}
