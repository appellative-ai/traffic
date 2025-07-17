package access

import (
	"fmt"
)

func Example_IsDirectOperator() {
	op := Operator{Name: "test", Value: "   %"}
	fmt.Printf("test: IsDirectOperator() -> %v [value:%v]\n", isDirectOperator(op), op.Value)

	op = Operator{Name: "test", Value: "%"}
	fmt.Printf("test: IsDirectOperator() -> %v [value:%v]\n", isDirectOperator(op), op.Value)

	//Output:
	//test: IsDirectOperator() -> true [value:   %]
	//test: IsDirectOperator() -> false [value:%]
}

func Example_IsRequestOperator() {
	op := Operator{}
	ok := isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(<empty>) -> %v\n", ok)

	op = Operator{Name: " ", Value: " "}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(<empty>) -> %v\n", ok)

	op = Operator{Name: "", Value: "REQ "}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(header"}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(header)"}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ()"}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(1)%"}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(header-name)%"}
	ok = isRequestOperator(op)
	fmt.Printf("test: isRequestOperator(%v) -> %v\n", op, ok)

	//Output:
	//test: isRequestOperator(<empty>) -> false
	//test: isRequestOperator(<empty>) -> false
	//test: isRequestOperator({ REQ }) -> false
	//test: isRequestOperator({ %REQ(header}) -> false
	//test: isRequestOperator({ %REQ(header)}) -> false
	//test: isRequestOperator({ %REQ()}) -> false
	//test: isRequestOperator({ %REQ(1)%}) -> true
	//test: isRequestOperator({ %REQ(header-name)%}) -> true

}

func Example_RequestOperatorHeaderName() {
	op := Operator{}
	name := requestOperatorHeaderName(op)
	fmt.Printf("test: requestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ("}
	name = requestOperatorHeaderName(op)
	fmt.Printf("test: requestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ()"}
	name = requestOperatorHeaderName(op)
	fmt.Printf("test: requestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ()%"}
	name = requestOperatorHeaderName(op)
	fmt.Printf("test: requestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ(1)%"}
	name = requestOperatorHeaderName(op)
	fmt.Printf("test: requestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ(name)%"}
	name = requestOperatorHeaderName(op)
	fmt.Printf("test: requestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	//Output:
	//test: requestOperatorHeaderName() ->  [op:]
	//test: requestOperatorHeaderName() ->  [op:%REQ(]
	//test: requestOperatorHeaderName() ->  [op:%REQ()]
	//test: requestOperatorHeaderName() ->  [op:%REQ()%]
	//test: requestOperatorHeaderName() -> 1 [op:%REQ(1)%]
	//test: requestOperatorHeaderName() -> name [op:%REQ(name)%]

}

func Example_IsStringValue() {
	op := Operator{Name: "test", Value: "   %"}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: "%"}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: DurationOperator}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: TimeoutDurationOperator}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: RateLimitOperator}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: ResponseStatusCodeOperator}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: ResponseBytesSentOperator}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	op = Operator{Name: "test", Value: ResponseBytesReceivedOperator}
	fmt.Printf("test: isStringValue() -> %v [value:%v]\n", isStringValue(op), op.Value)

	//Output:
	//test: isStringValue() -> true [value:   %]
	//test: isStringValue() -> true [value:%]
	//test: isStringValue() -> false [value:%DURATION%]
	//test: isStringValue() -> false [value:%TIMEOUT_DURATION%]
	//test: isStringValue() -> false [value:%RATE_LIMIT%]
	//test: isStringValue() -> false [value:%STATUS_CODE%]
	//test: isStringValue() -> false [value:%BYTES_SENT%]
	//test: isStringValue() -> false [value:%BYTES_RECEIVED%]

}
