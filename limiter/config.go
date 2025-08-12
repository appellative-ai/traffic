package limiter

import "github.com/appellative-ai/core/messaging"

func (a *agentT) config(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	//if a.running {
	//	return
	//}
	if t, ok := messaging.ConfigContent[map[string]string](m); ok {
		state := a.state.Load()
		limit := state.Limit
		reviewDuration := state.ReviewDuration
		// Changed?
		if !state.Update(t) {
			return
		}
		// TODO : reset burst based on new limit??
		if limit != state.Limit {
			a.limiter.SetLimit(state.Limit)
		}
		dur := a.state.Load().ReviewDuration
		if reviewDuration != dur && dur > 0 {
			review := a.review.Load()
			if !review.Started() {
				review.Start(dur)
			}
		}
		return
	}
	if d, ok1 := messaging.ConfigContent[messaging.Dispatcher](m); ok1 {
		a.dispatcher = d
		return
	}
}
