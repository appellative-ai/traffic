package analytics

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
	"strings"
)

const (
	RegionKey                 = "region"
	ZoneKey                   = "zone"
	SubZoneKey                = "sub-zone"
	HostKey                   = "host"
	InstanceIdKey             = "id"
	RegionZoneHostFmt         = "%v:%v.%v.%v"
	RegionZoneSubZoneHostFmt  = "%v:%v.%v.%v.%v"
	uriFmt                    = "%v:%v"
	RegionZoneSubZoneHostFmt2 = "%v.%v.%v.%v"
)

// Origin - location
type Origin struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func (o Origin) Tag2() string {
	tag := o.Region
	if o.Zone != "" {
		tag += ":" + o.Zone
	}
	if o.SubZone != "" {
		tag += ":" + o.SubZone
	}
	if o.Host != "" {
		tag += ":" + o.Host
	}
	return tag
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

func NewValues(o Origin) url.Values {
	values := make(url.Values)
	if o.Region != "" {
		values.Add(RegionKey, o.Region)
	}
	if o.Zone != "" {
		values.Add(ZoneKey, o.Zone)
	}
	if o.SubZone != "" {
		values.Add(SubZoneKey, o.SubZone)
	}
	if o.Host != "" {
		values.Add(HostKey, o.Host)
	}
	return values
}

func NewOrigin(values url.Values) Origin {
	o := Origin{}
	if values != nil {
		o.Region = values.Get(RegionKey)
		o.Zone = values.Get(ZoneKey)
		o.SubZone = values.Get(SubZoneKey)
		o.Host = values.Get(HostKey)
	}
	return o
}

func OriginMatch(target Origin, filter Origin) bool {
	isFilter := false
	if filter.Region != "" {
		if filter.Region == "*" {
			return true
		}
		isFilter = true
		if !StringMatch(target.Region, filter.Region) {
			return false
		}
	}
	if filter.Zone != "" {
		isFilter = true
		if !StringMatch(target.Zone, filter.Zone) {
			return false
		}
	}
	if filter.SubZone != "" {
		isFilter = true
		if !StringMatch(target.SubZone, filter.SubZone) {
			return false
		}
	}
	if filter.Host != "" {
		isFilter = true
		if !StringMatch(target.Host, filter.Host) {
			return false
		}
	}
	return isFilter
}

func StringMatch(target, filter string) bool {
	//if filter == "" {
	//	return true
	//}
	return strings.ToLower(target) == strings.ToLower(filter)
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
