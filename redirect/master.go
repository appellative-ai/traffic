package redirect

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
)

// master attention
func masterAttend(agent *agentT, resolver *content.Resolution) {
	agent.dispatch(agent.master, messaging.StartupEvent)
	paused := false
	if paused {
	}

	for {
		select {
		case msg := <-agent.master.C:
			agent.dispatch(agent.master, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.masterShutdown()
				return
			default:
			}
		default:
		}
	}
}
