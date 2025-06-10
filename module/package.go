package module

import (
	"github.com/behavioral-ai/traffic/cache"
	"github.com/behavioral-ai/traffic/limiter"
	"github.com/behavioral-ai/traffic/redirect"
	"github.com/behavioral-ai/traffic/routing"
)

var (
	LimiterNamespaceName  = limiter.NamespaceName
	RedirectNamespaceName = redirect.NamespaceName
	CacheNamespaceName    = cache.NamespaceName
	RoutingNamespaceName  = routing.NamespaceName
)

func Resolve(name string) (bool, any) {
	switch name {
	case limiter.NamespaceName:
		return true, nil
	case redirect.NamespaceName:
		return true, nil
	default:
		return false, nil
	}
}
