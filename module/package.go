package module

import (
	"github.com/appellative-ai/traffic/authorization"
	"github.com/appellative-ai/traffic/cache"
	"github.com/appellative-ai/traffic/limiter"
	"github.com/appellative-ai/traffic/logger"
	"github.com/appellative-ai/traffic/routing"
)

var (
	LimiterNamespaceName       = limiter.NamespaceName
	CacheNamespaceName         = cache.NamespaceName
	RoutingNamespaceName       = routing.NamespaceName
	LoggerNamespaceName        = logger.NamespaceName
	AuthorizationNamespaceName = authorization.NamespaceName
)
