package representation1

import (
	"fmt"
)

const (
	NamespaceName = "test:resiliency:agent/cache/request/http"
)

var (
	m = map[string]string{
		HostKey:            "www.google.com",
		CacheControlKey:    "no-store, no-cache, max-age=0",
		TimeoutDurationKey: "750ms",
		SundayKey:          "13-15",
		MondayKey:          "8-16",
		TuesdayKey:         "6-10",
		WednesdayKey:       "12-12",
		ThursdayKey:        "0-23",
		FridayKey:          "22-23",
		SaturdayKey:        "3-8",
	}

	m2 = map[string]string{
		"host":          "www.google.com",
		"cache-control": "no-store, no-cache, max-age=0",
		"sun":           "13-15",
		"mon":           "8-16",
		"tue":           "14-21",
		"wed":           "12-12",
		"thu":           "19-23",
		"fri":           "22-23",
		"sat":           "3-8",
	}
)

func ExampleParseCache() {
	var cache Cache
	parseCache(&cache, m)

	fmt.Printf("test: parseCache() -> %v\n", cache)

	//Output:
	//test: parseCache() -> {750ms www.google.com map[Cache-Control:[no-store, no-cache, max-age=0]] map[fri:{22 23} mon:{8 16} sat:{3 8} sun:{13 15} thu:{0 23} tue:{6 10} wed:{12 12}]}

}
