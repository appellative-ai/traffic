package limiter

import (
	"github.com/appellative-ai/collective/operations/operationstest"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/traffic/limiter/representation1"
	"time"
)

const (
	testDuration = time.Second * 2
)

func ExampleMaster() {
	ch := make(chan struct{})
	agent := newAgent(representation1.Initialize(nil), operationstest.NewNotifier())

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
	//test.LoadResiliencyContent()
	//resolver, status := test.NewResiliencyResolver()
	//if !status.OK() {
	//	metrics.Notify(status)
	//}
	agent := newAgent(representation1.Initialize(nil), operationstest.NewNotifier())

	go func() {
		go masterAttend(agent, nil)
		//agent.Message(msg)
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
