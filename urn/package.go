package urn

import (
	"github.com/behavioral-ai/traffic/analytics"
	"github.com/behavioral-ai/traffic/limiter"
	"github.com/behavioral-ai/traffic/redirect"
)

const (
	AnalyticsAgent = analytics.NamespaceName
	LimiterAgent   = limiter.NamespaceName
	RedirectAgent  = redirect.NamespaceName
)
