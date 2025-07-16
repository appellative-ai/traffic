package access

/*
const (
	RegionKey                = "region"
	ZoneKey                  = "zone"
	SubZoneKey               = "sub-zone"
	HostKey                  = "host"
	InstanceIdKey            = "id"
	RegionZoneHostFmt        = "%v:%v.%v.%v"
	RegionZoneSubZoneHostFmt = "%v:%v.%v.%v.%v"
)


*/

// origin - location
type originT struct {
	//AppId      string
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func (o originT) Tag() string {
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

/*
func (o originT) Uri(class string) string {
	var uri string
	if o.SubZone == "" {
		uri = fmt.Sprintf(RegionZoneHostFmt, class, o.Region, o.Zone, o.Host)
	} else {
		uri = fmt.Sprintf(RegionZoneSubZoneHostFmt, class, o.Region, o.Zone, o.SubZone, o.Host)
	}
	return uri
}

func NewValues(o originT) url.Values {
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

func NewOrigin(values url.Values) originT {
	o := originT{}
	if values != nil {
		o.Region = values.Get(RegionKey)
		o.Zone = values.Get(ZoneKey)
		o.SubZone = values.Get(SubZoneKey)
		o.Host = values.Get(HostKey)
		//o.Route = values.Get(RouteKey)
	}
	return o
}


*/
