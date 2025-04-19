package timeseries

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

const (
	RegionKey     = "region"
	ZoneKey       = "zone"
	SubZoneKey    = "sub-zone"
	HostKey       = "host"
	InstanceIdKey = "instance-id"
	uriFmt        = "%v:%v"
)

// Origin - location
type Origin struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func (o Origin) Uri(class string) string {
	return fmt.Sprintf(uriFmt, class, o)
}

func (o Origin) String() string {
	var uri = o.Region

	if o.Zone != "" {
		uri += "." + o.Zone
	}
	if o.SubZone != "" {
		uri += "." + o.SubZone
	}
	if o.Host != "" {
		uri += "." + o.Host
	}
	return uri
}

func NewOriginFromMessage(agent messaging.Agent, m *messaging.Message) (o Origin, ok bool) {
	a := agent
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
		return
	}
	region := cfg[RegionKey]
	if region == "" {
		return
	}
	o.Region = region
	o.Zone = cfg[ZoneKey]
	if o.Zone == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, ZoneKey), a.Uri())
		return
	}
	o.SubZone = cfg[SubZoneKey]
	if o.SubZone == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, SubZoneKey), a.Uri())
		return
	}
	o.Host = cfg[HostKey]
	if o.Host == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, HostKey), a.Uri())
		return
	}
	o.InstanceId = cfg[InstanceIdKey]
	if o.Host == "" {
		messaging.Reply(m, messaging.ConfigContentStatusError(a, InstanceIdKey), a.Uri())
		return
	}
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
	return o, true
}
