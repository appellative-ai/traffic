package module

import (
	"github.com/behavioral-ai/traffic/cache"
	"github.com/behavioral-ai/traffic/limiter"
	"github.com/behavioral-ai/traffic/routing"
)

var (
	LimiterNamespaceName = limiter.NamespaceName
	CacheNamespaceName   = cache.NamespaceName
	RoutingNamespaceName = routing.NamespaceName
)
