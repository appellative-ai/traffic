package redirect

import (
	"github.com/behavioral-ai/collective/eventing/eventtest"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/traffic/timeseries"
	"time"
)

func _ExampleMaster() {
	ch := make(chan struct{})
	agent := newAgent(eventtest.New())

	go func() {
		go masterAttend(agent, timeseries.Functions)

		agent.Message(messaging.NewMessage(messaging.ChannelMaster, messaging.PauseEvent))
		agent.Message(messaging.NewMessage(messaging.ChannelMaster, messaging.ResumeEvent))

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
	//msg := messaging.NewMessage(messaging.ChannelMaster, messaging.ObservationEvent)
	//msg.SetContent(contentTypeObservation, observation{origin: origin, latency: 2350, gradient: 15})
	//test.LoadResiliencyContent()
	//resolver, status := test.NewResiliencyResolver()
	//if !status.OK() {
	//	metrics.Notify(status)
	//}
	agent := newAgent(eventtest.New())

	go func() {
		go masterAttend(agent, timeseries.Functions)
		//agent.Message(msg)
		//time.Sleep(testDuration * 2)

		agent.Message(messaging.ShutdownMessage)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
