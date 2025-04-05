package analytics

import "fmt"

func ExampleNew() {
	a := New(nil)

	fmt.Printf("test: ExampleNew() -> [%v]\n", a)

	//Output:
	//fail
}
