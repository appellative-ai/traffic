package config

import (
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
	"time"
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
func SetOrigin(agent messaging.Agent, m *messaging.Message) (o Origin, ok bool) {
	a := agent
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
		return
	}
	region := cfg[regionKey]
	if region == "" {
		return
	}
	o.Region = region
	o.Zone = cfg[zoneKey]
	if o.Zone == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, zoneKey), a.Uri())
		return
	}
	o.SubZone = cfg[subZoneKey]
	if o.SubZone == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, subZoneKey), a.Uri())
		return
	}
	o.Host = cfg[hostKey]
	if o.Host == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, hostKey), a.Uri())
		return
	}
	o.InstanceId = cfg[instanceIdKey]
	if o.Host == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, instanceIdKey), a.Uri())
		return
	}
	origin = o
	set = true
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
	return o, true
}


*/
