package limiter

import (
	"fmt"
	"github.com/behavioral-ai/core/eventing/eventtest"
)

func ExampleNewAgent() {
	a := newAgent(eventtest.New())
	fmt.Printf("test: newAgent() -> [limiter:%v] [burst:%v] [%v}\n", a.limiter.Limit(), a.limiter.Burst(), a.Uri())

	//agent := agentT{}
	//t := reflect.TypeOf(agent)
	//fmt.Printf("test: agenT -> [%v] [name:%v] [path:%v] [kind:%v]\n", t, t.Name(), t.PkgPath(), t.Kind())

	//t = reflect.TypeOf(New)
	//fmt.Printf("test: New() -> [%v] [name:%v] [path:%v] [kind:%v]\n", t, t.Name(), t.PkgPath(), t.Kind())

	//Output:
	//test: newAgent() -> [limiter:50] [burst:10] [resiliency:agent/request/rate-limiting}

}

/*
func ExampleAgent_LoadContent() {
	ch := make(chan struct{})
	dispatcher := event.NewTraceDispatcher()
	origin := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	s := messagingtest.NewTestSpanner(time.Second*2, testDuration)
	//test.LoadResiliencyContent()
	agent := newAgent(origin, eventtest.New(dispatcher))

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
