package config

import (
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
)

const (
	ThresholdLatencyKey = "threshold-latency" // limiter agent threshold latency
)

//ScoreKey   = "score"    // limiter threshold score

func Threshold(agent messaging.Agent, m *messaging.Message) (latency int, ok bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(agent), agent.Uri())
		return -1, false
	}
	/*
		score := cfg[ScoreKey]
		if score == "" {
			messaging.Reply(m, messaging.ConfigContentStatusError(agent, ScoreKey), agent.Uri())
			return timeseries.Percentile{}, false
		}
		var err error
		p.Score, err = strconv.Atoi(score)
		if err != nil {
			messaging.Reply(m, messaging.ConfigContentStatusError(agent, ScoreKey), agent.Uri())
			return timeseries.Percentile{}, false
		}


	*/

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

/*
func SetOrigin(agent metrics.Agent, m *metrics.Message) (o Origin, ok bool) {
	a := agent
	cfg := metrics.ConfigMapContent(m)
	if cfg == nil {
		metrics.Reply(m, metrics.ConfigEmptyStatusError(a), a.Uri())
		return
	}
	region := cfg[regionKey]
	if region == "" {
		return
	}
	o.Region = region
	o.Zone = cfg[zoneKey]
	if o.Zone == "" {
		metrics.Reply(m, metrics.ConfigContentStatusError(a, zoneKey), a.Uri())
		return
	}
	o.SubZone = cfg[subZoneKey]
	if o.SubZone == "" {
		metrics.Reply(m, metrics.ConfigContentStatusError(a, subZoneKey), a.Uri())
		return
	}
	o.Host = cfg[hostKey]
	if o.Host == "" {
		metrics.Reply(m, metrics.ConfigContentStatusError(a, hostKey), a.Uri())
		return
	}
	o.InstanceId = cfg[instanceIdKey]
	if o.Host == "" {
		metrics.Reply(m, metrics.ConfigContentStatusError(a, instanceIdKey), a.Uri())
		return
	}
	origin = o
	set = true
	metrics.Reply(m, metrics.StatusOK(), a.Uri())
	return o, true
}


*/
