package metrics

import (
	"fmt"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/fmtx"
	"time"
)

func ExampleLatencySample() {
	s := new(LatencySample)
	s.Update(&timeseries.Event{Start: time.Now().UTC(), Duration: time.Second * 2})

	fmt.Printf("test: Update()  -> [first:%v] [last:%v]\n", fmtx.FmtRFC3339Millis(s.First), fmtx.FmtRFC3339Millis(s.Last))
	fmt.Printf("test: Elapsed() -> [elapsed:%v]\n", s.Elapsed())

	s.Update(&timeseries.Event{Start: time.Date(2025, 4, 5, 22, 8, 45, 0, time.UTC), Duration: time.Millisecond * 1500})
	fmt.Printf("test: Update()  -> [first:%v] [last:%v]\n", fmtx.FmtRFC3339Millis(s.First), fmtx.FmtRFC3339Millis(s.Last))
	fmt.Printf("test: Elapsed() -> [elapsed:%v]\n", s.Elapsed())

	//Output:
	//test: Update()  -> [first:2025-04-05T22:19:08.286Z] [last:2025-04-05T22:19:10.286Z]
	//test: Elapsed() -> [elapsed:2s]
	//test: Update()  -> [first:2025-04-05T22:08:45.000Z] [last:2025-04-05T22:19:10.286Z]
	//test: Elapsed() -> [elapsed:10m25.2867158s]

}
