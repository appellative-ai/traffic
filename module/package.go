package module

import (
	"github.com/appellative-ai/traffic/access"
	"github.com/appellative-ai/traffic/authorization"
	"github.com/appellative-ai/traffic/cache"
	"github.com/appellative-ai/traffic/limiter"
	"github.com/appellative-ai/traffic/routing"
)

var (
	LimiterNamespaceName       = limiter.NamespaceName
	CacheNamespaceName         = cache.NamespaceName
	RoutingNamespaceName       = routing.NamespaceName
	AccessNamespaceName        = access.NamespaceName
	AuthorizationNamespaceName = authorization.NamespaceName
)
