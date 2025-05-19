package test

import (
	"fmt"
	"github.com/behavioral-ai/collective/repository"
	_ "github.com/behavioral-ai/traffic/module"
)

func Load() {
	agent := repository.Agent("resiliency:agent/rate-limiting/request/http")
	fmt.Printf("test: LoadTest() -> [%v]\n", agent)
	agent = repository.Agent("resiliency:agent/redirect/request/http")
	fmt.Printf("test: LoadTest() -> [%v]\n", agent)

}
