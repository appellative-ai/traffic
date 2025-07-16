package test

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	_ "github.com/appellative-ai/traffic/module"
)

func Load() {
	agent := exchange.Agent("resiliency:agent/rate-limiting/request/http")
	fmt.Printf("test: LoadTest() -> [%v]\n", agent)
	agent = exchange.Agent("resiliency:agent/redirect/request/http")
	fmt.Printf("test: LoadTest() -> [%v]\n", agent)

}
