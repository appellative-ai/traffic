package redirect

import (
	"github.com/behavioral-ai/core/eventing/eventtest"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/redirect/representation1"
	"time"
)

const (
	testDuration = time.Second * 5
)

func ExampleEmissary() {
	ch := make(chan struct{})
	agent := newAgent(eventtest.New(), representation1.Initialize())

	go func() {
		go emissaryAttend(agent)

		agent.Message(messaging.NewMessage(messaging.ChannelEmissary, messaging.PauseEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.ChannelEmissary, messaging.ResumeEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}

func ExampleEmissary_Observation() {
	ch := make(chan struct{})
	agent := newAgent(eventtest.New(), representation1.Initialize())

	go func() {
		go emissaryAttend(agent)
		time.Sleep(testDuration * 2)

		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration * 3)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}
