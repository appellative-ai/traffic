package access

import (
	"fmt"
	"net/http"
	"time"
)

func ExampleSet() {
	h := make(http.Header)
	SetTimeout(h, time.Millisecond*456)
	SetRateLimit(h, float64(123))
	SetRedirect(h, 35)
	SetCached(h, false)

	fmt.Printf("test: Set() -> %v\n", h)
	RemoveThresholds(h)
	fmt.Printf("test: Set() -> %v\n", h)

	//Output:
	//test: Set() -> map[X-Thresholds:[timeout=456ms rate-limit=123 redirect=35 cached=false]]
	//test: Set() -> map[]

}
