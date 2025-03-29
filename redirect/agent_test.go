package redirect

import (
	"fmt"
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/collective/eventing/eventtest"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/messaging/messagingtest"
	"time"
)

func ExampleNewAgent() {
	a := newAgent(nil)
	a.hostName = "localhost:8080"
	fmt.Printf("test: newAgent() -> [host:%v] [uri:%v}\n", a.hostName, a.Uri())

	//Output:
	//test: newAgent() -> [host:localhost:8080] [uri:resiliency:agent/behavioral-ai/resiliency/redirect}

}

func _ExampleAgent_LoadContent() {
	ch := make(chan struct{})
	dispatcher := eventing.NewTraceDispatcher()
	s := messagingtest.NewTestSpanner(time.Second*2, testDuration)
	//test.LoadResiliencyContent()
	agent := newAgent(eventtest.New(dispatcher))

	go func() {
		go masterAttend(agent, content.Resolver)
		go emissaryAttend(agent, content.Resolver, s)
		time.Sleep(testDuration * 5)

		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}

func _ExampleAgent_NotFound() {
	ch := make(chan struct{})
	dispatcher := eventing.NewTraceDispatcher()
	agent := newAgent(eventtest.New(dispatcher))

	go func() {
		agent.Message(messaging.StartupMessage)
		time.Sleep(testDuration * 5)
		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}

func _ExampleAgent_Resolver() {
	ch := make(chan struct{})
	dispatcher := eventing.NewTraceDispatcher()
	agent := newAgent(eventtest.New(dispatcher))
	//test2.Startup()

	go func() {
		agent.Message(messaging.StartupMessage)
		time.Sleep(testDuration * 5)
		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}
