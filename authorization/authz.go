package authorization

import (
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/rest"
	"net/http"
)

const (
	NamespaceName = "test:resiliency:handler/authorization/http"
	AuthzName     = "Authorization"
)

func init() {
	exchange.RegisterExchangeHandler(NamespaceName, Authorization)
}

func Authorization(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(AuthzName)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		return next(r)
	}
}
