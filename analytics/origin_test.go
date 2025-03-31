package analytics

import "fmt"

func ExampleNewValues() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	values := NewValues(o)
	fmt.Printf("test: NewValues() -> [%v]\n", values)

	//Output:
	//test: NewValues() -> [map[host:[host] region:[region] sub-zone:[sub-zone] zone:[zone]]]

}

func ExampleNewOrigin() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	//values := NewValues(o)
	//o = NewOrigin(values)
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = ""
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = ""
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = ""
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = "host"
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	//Output:
	//test: NewOrigin() -> [region.zone.sub-zone.host]
	//test: NewOrigin() -> [region.sub-zone.host]
	//test: NewOrigin() -> [region.zone.host]
	//test: NewOrigin() -> [region.zone.sub-zone]
	//test: NewOrigin() -> [region.zone.sub-zone.host]

}

/*
func ExampleOrigin_Tag() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	o.Zone = ""
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	o.Host = ""
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	o.SubZone = ""
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	//Output:
	//test: Tag() -> [region:zone:sub-zone:host]
	//test: Tag() -> [region:sub-zone:host]
	//test: Tag() -> [region:sub-zone]
	//test: Tag() -> [region]

}


*/

func ExampleOriginMatch() {
	target := Origin{
		Region:     "Region",
		Zone:       "zonE",
		SubZone:    "sub-zone",
		Host:       "hOst",
		InstanceId: "",
	}
	filter := Origin{
		Region:     "",
		Zone:       "",
		SubZone:    "",
		Host:       "",
		InstanceId: "",
	}

	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.Region = "region"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.Zone = "zone"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.SubZone = "sub-zone"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.Host = "host"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.SubZone = ""
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.SubZone = "invalid"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	//Output:
	//test: OriginMatch({Region zonE sub-zone hOst  },{     }) -> [match:false]
	//test: OriginMatch({Region zonE sub-zone hOst  },{region     }) -> [match:true]
	//test: OriginMatch({Region zonE sub-zone hOst  },{region zone    }) -> [match:true]
	//test: OriginMatch({Region zonE sub-zone hOst  },{region zone sub-zone   }) -> [match:true]
	//test: OriginMatch({Region zonE sub-zone hOst  },{region zone sub-zone host  }) -> [match:true]
	//test: OriginMatch({Region zonE sub-zone hOst  },{region zone  host  }) -> [match:true]
	//test: OriginMatch({Region zonE sub-zone hOst  },{region zone invalid host  }) -> [match:false]

}

func ExampleOrigin_Uri() {
	target := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target.Uri("class"))
	//fmt.Printf("test: Origin_Uri_SubZone_Route() -> [%v]\n", target.Uri("class"))

	target.SubZone = ""
	fmt.Printf("test: Origin_Uri_No_SubZone()    -> [%v]\n", target.Uri("class"))

	//Output:
	//test: Origin_Uri_SubZone()       -> [class:region.zone.sub-zone.host]
	//test: Origin_Uri_No_SubZone()    -> [class:region.zone.host]

}

func ExampleOrigin_String() {
	target := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target)
	//fmt.Printf("test: Origin_Uri_SubZone_Route() -> [%v]\n", target)

	target.SubZone = ""
	fmt.Printf("test: Origin_Uri_No_SubZone()    -> [%v]\n", target)

	//Output:
	//test: Origin_Uri_SubZone()       -> [region.zone.sub-zone.host]
	//test: Origin_Uri_No_SubZone()    -> [region.zone.host]

}
