package access

import "fmt"

func ExampleParsed() {
	u, p := parseURL("", nil)

	fmt.Printf("test: parseURL() -> %v %v\n", u, p)

	//Output:
	//
}
