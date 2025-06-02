package limiter

import (
	"fmt"
	"github.com/behavioral-ai/traffic/limiter/representation1"
)

func ExampleNewAgent() {
	a := newAgent(representation1.Initialize(nil), nil)
	fmt.Printf("test: newAgent() -> [limiter:%v] [burst:%v] [%v}\n", a.limiter.Limit(), a.limiter.Burst(), a.Name())

	//agent := agentT{}
	//t := reflect.TypeOf(agent)
	//fmt.Printf("test: agenT -> [%v] [name:%v] [path:%v] [kind:%v]\n", t, t.Name(), t.PkgPath(), t.Kind())

	//t = reflect.TypeOf(New)
	//fmt.Printf("test: New() -> [%v] [name:%v] [path:%v] [kind:%v]\n", t, t.Name(), t.PkgPath(), t.Kind())

	//Output:
	//test: newAgent() -> [limiter:50] [burst:10] [resiliency:agent/rate-limiting/request/http}

}

/*
func ExampleAgent_LoadContent() {
	ch := make(chan struct{})
	dispatcher := event.NewTraceDispatcher()
	s := messagingtest.NewTestSpanner(time.Second*2, testDuration)
	//test.LoadResiliencyContent()
	agent := newAgent()

	go func() {
		go masterAttend(agent, content.Resolver)
		go emissaryAttend(agent, content.Resolver, s)
		time.Sleep(testDuration * 5)

		agent.Message(metrics.ShutdownMessage)
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}


*/
