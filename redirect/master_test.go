package redirect

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/eventing/eventtest"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

func _ExampleMaster() {
	ch := make(chan struct{})
	agent := newAgent(eventtest.New())

	go func() {
		go masterAttend(agent, content.Resolver)
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))

		agent.Message(messaging.NewMessage(messaging.Master, messaging.PauseEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ResumeEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))

		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)
	//Output:
	//fail
}

func _ExampleMaster_Observation() {
	ch := make(chan struct{})
	msg := messaging.NewMessage(messaging.Master, messaging.ObservationEvent)
	//msg.SetContent(contentTypeObservation, observation{origin: origin, latency: 2350, gradient: 15})
	//test.LoadResiliencyContent()
	//resolver, status := test.NewResiliencyResolver()
	//if !status.OK() {
	//	messaging.Notify(status)
	//}
	agent := newAgent(eventtest.New())

	go func() {
		go masterAttend(agent, content.Resolver)
		agent.Message(msg)
		time.Sleep(testDuration * 2)

		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
