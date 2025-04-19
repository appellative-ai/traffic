package timeseries

import "fmt"

func ExampleNewOrigin() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "instance-id",
	}
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

func ExampleOrigin_Uri() {
	target := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "instance-id",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target.Uri("class"))

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
		InstanceId: "instance-id",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target)

	target.SubZone = ""
	fmt.Printf("test: Origin_Uri_No_SubZone()    -> [%v]\n", target)

	//Output:
	//test: Origin_Uri_SubZone()       -> [region.zone.sub-zone.host]
	//test: Origin_Uri_No_SubZone()    -> [region.zone.host]

}
