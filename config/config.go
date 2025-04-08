package config

import (
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	AppHostKey = "app-host"
	TimeoutKey = "timeout"
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
