package config

import (
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	TimeoutKey = "timeout"
)

//ScoreKey   = "score"    // limiter threshold score
//ThresholdLatencyKey = "threshold-latency"

/*
	func Threshold(agent messaging.Agent, m *messaging.Message) (latency int, ok bool) {
		cfg := messaging.ConfigMapContent(m)
		if cfg == nil {
			messaging.Reply(m, messaging.ConfigEmptyStatusError(agent), agent.Uri())
			return -1, false
		}
		s := cfg[ThresholdLatencyKey]
		if s == "" {
			messaging.Reply(m, messaging.ConfigContentStatusError(agent, ThresholdLatencyKey), agent.Uri())
			return -1, false
		}
		dur, err1 := fmtx.ParseDuration(s)
		if err1 != nil {
			messaging.Reply(m, messaging.ConfigContentStatusError(agent, ThresholdLatencyKey), agent.Uri())
			return -1, false
		}
		return fmtx.Milliseconds(dur), true
	}
*/

func Timeout(agent messaging.Agent, m *messaging.Message) (time.Duration, bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(agent), agent.Name())
		return 0, false
	}
	timeout := cfg[TimeoutKey]
	if timeout == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
		return 0, false
	}
	dur, err := fmtx.ParseDuration(timeout)
	if err != nil {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Name())
		return 0, false
	}
	return dur, true
}
