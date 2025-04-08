package analytics

import (
	"github.com/behavioral-ai/core/messaging"
)

// master attention
func masterAttend(a *agentT) {
	paused := false
	if paused {
	}

	for {
		select {
		case m := <-a.master.C:
			switch m.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				a.masterShutdown()
				return
			default:
			}
		default:
		}
	}
}
