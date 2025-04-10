package config

import (
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
	"strconv"
	"time"
)

const (
	AppHostKey = "app-host" // routing host name
	TimeoutKey = "timeout"  // routing timeout
	ScoreKey   = "score"    // limiter threshold score
	LatencyKey = "latency"  // limiter threshold latency
)

func Timeout(agent messaging.Agent, m *messaging.Message) (time.Duration, bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(agent), agent.Uri())
		return 0, false
	}
	timeout := cfg[TimeoutKey]
	if timeout == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Uri())
		return 0, false
	}
	dur, err := fmtx.ParseDuration(timeout)
	if err != nil {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, TimeoutKey), agent.Uri())
		return 0, false
	}
	return dur, true
}

func AppHostName(agent messaging.Agent, m *messaging.Message) (string, bool) {
	return hostName(agent, m, AppHostKey)
}

func Percentile(agent messaging.Agent, m *messaging.Message) (p timeseries.Percentile, ok bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(agent), agent.Uri())
		return timeseries.Percentile{}, false
	}
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

	latency := cfg[LatencyKey]
	if latency == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, LatencyKey), agent.Uri())
		return timeseries.Percentile{}, false
	}
	dur, err1 := fmtx.ParseDuration(latency)
	if err1 != nil {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, LatencyKey), agent.Uri())
		return timeseries.Percentile{}, false
	}
	p.Latency = fmtx.Milliseconds(dur)
	return p, true
}

func hostName(agent messaging.Agent, m *messaging.Message, key string) (string, bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(agent), agent.Uri())
		return "", false
	}
	host := cfg[key]
	if host == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(agent, key), agent.Uri())
		return "", false
	}
	return host, true
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
