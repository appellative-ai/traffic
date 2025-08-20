package module

import (
	"github.com/appellative-ai/traffic/authorization"
	"github.com/appellative-ai/traffic/cache"
	"github.com/appellative-ai/traffic/limiter"
	"github.com/appellative-ai/traffic/routing"
)

var (
	LimiterAgentName         = limiter.AgentName
	CacheAgentName           = cache.AgentName
	RoutingAgentName         = routing.AgentName
	AuthorizationHandlerName = authorization.HandlerName
)
