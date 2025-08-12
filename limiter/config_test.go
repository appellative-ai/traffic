package limiter

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/traffic/limiter/representation1"
)

func _ExampleConfig_Dispatcher() {
	a := newAgent(nil)
	fmt.Printf("test: newAgent() -> [dispatcher:%v]\n", a.dispatcher)

	a.Message(messaging.NewConfigMessage(messaging.NewTraceDispatcher()))
	fmt.Printf("test: Message() -> [dispatcher:%v]\n", a.dispatcher != nil)

	//Output:
	//test: newAgent() -> [dispatcher:<nil>]
	//test: Message() -> [dispatcher:true]

}

func ExampleConfig_Limiter() {
	a := newAgent(nil)
	fmt.Printf("test: newAgent() -> [limit:%v] [limiter-burst:%v]\n", a.state.Load().Limit, a.limiter.Limit())

	m := map[string]string{
		representation1.RateLimitKey: "125",
	}
	a.Message(messaging.NewConfigMessage(m))
	fmt.Printf("test: Message() -> [limit:%v] [limiter-limit:%v]\n", a.state.Load().Limit, a.limiter.Limit())

	//Output:
	//test: newAgent() -> [limit:50] [limiter-burst:50]
	//test: Message() -> [limit:125] [limiter-limit:125]

}

func ExampleConfig_Review() {
	a := newAgent(nil)
	fmt.Printf("test: newAgent() -> [dur:%v] [review-dur:%v]\n", a.state.Load().ReviewDuration, a.review.Load().Duration())

	m := map[string]string{
		representation1.ReviewDurationKey: "25",
	}
	a.Message(messaging.NewConfigMessage(m))
	fmt.Printf("test: Message() -> [dur:%v] [review-dur:%v]\n", a.state.Load().ReviewDuration, a.review.Load().Duration())

	//Output:
	//test: newAgent() -> [dur:0s] [review-dur:0s]
	//test: Message() -> [dur:25s] [review-dur:25s]

}
