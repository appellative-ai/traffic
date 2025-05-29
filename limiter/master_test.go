package limiter

import (
	"github.com/behavioral-ai/core/eventing/eventtest"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	testDuration = time.Second * 2
)

func ExampleMaster() {
	ch := make(chan struct{})
	agent := newAgent(eventtest.New(), nil, nil)

	go func() {
		go masterAttend(agent, nil)

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

func ExampleMaster_Observation() {
	ch := make(chan struct{})
	//msg := metrics.NewMessage(metrics.Master, metrics.ObservationEvent)
	//msg.SetContent(contentTypeObservation, observation{origin: origin, latency: 2350, gradient: 15})
	//test.LoadResiliencyContent()
	//resolver, status := test.NewResiliencyResolver()
	//if !status.OK() {
	//	metrics.Notify(status)
	//}
	agent := newAgent(eventtest.New(), nil, nil)

	go func() {
		go masterAttend(agent, nil)
		//agent.Message(msg)
		time.Sleep(testDuration * 2)

		messaging.Shutdown(agent)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
