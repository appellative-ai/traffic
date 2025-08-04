package cache

import (
	"bytes"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/collective/notification"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/core/uri"
	"github.com/appellative-ai/traffic/cache/representation1"
	"io"
	"net/http"
)

const (
	NamespaceName = "test:resiliency:agent/cache/request/http"
	cachedName    = "cached" // Sync with core/access

)

var (
	noContentResponse   = httpx.NewResponse(http.StatusNoContent, nil, nil)
	serverErrorResponse = httpx.NewResponse(http.StatusInternalServerError, nil, nil)
)

type agentT struct {
	state    *representation1.Cache
	exchange rest.Exchange //func(r *http.Request) (*http.Response,error)
	notifier *notification.Interface

	review   *messaging.Review
	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent()
	})
}

func newAgent() *agentT {
	a := new(agentT)
	a.state = representation1.Initialize(nil)
	a.notifier = notification.Notifier
	a.exchange = httpx.Do

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.state.Interval)
	a.emissary = messaging.NewEmissaryChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		if a.state.Running {
			return
		}
		messaging.UpdateContent[rest.Exchange](m, &a.exchange)
		messaging.UpdateContent[*messaging.Review](m, &a.review)
		messaging.UpdateMap(a.Name(), func(cfg map[string]string) {
			a.state.Update(cfg)
		}, m)
		return
	case messaging.StartupEvent:
		if a.state.Running {
			return
		}
		a.run()
		a.state.Running = true
		return
	case messaging.ShutdownEvent:
		if !a.state.Running {
			return
		}
		a.state.Running = false
	}
	if m.Channel() != messaging.ChannelMaster {
		a.emissary.C <- m
	}
}

// Run - run the agent
func (a *agentT) run() {
	go emissaryAttend(a)
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		if !a.cacheable(r) {
			return next(r)
		}
		var (
			url    string
			status *std.Status
		)
		// cache lookup
		url = uri.BuildURL(a.state.Host, r.URL.Path, r.URL.Query())
		h := make(http.Header)
		h.Add(httpx.XRequestId, r.Header.Get(httpx.XRequestId))
		resp, status = do(a, http.MethodGet, url, h, nil)
		if resp.StatusCode == http.StatusOK {
			resp.Header.Add(cachedName, "true")
			return resp, nil
		}
		resp.Header.Add(cachedName, "false")
		if status.Err != nil {
			a.notifier.Message(messaging.NewStatusMessage(status, a.Name())) //.WithLocation(a.Name()), a.Name()))
		}
		// cache miss, call next exchange
		resp, err = next(r)
		if resp.StatusCode == http.StatusOK {
			// cache update
			err = a.cacheUpdate(url, r, resp)
			if err != nil {
				return serverErrorResponse, err
			}
		}
		return
	}
}

func (a *agentT) trace(task, observation, action string) {
	if a.review == nil {
		return
	}
	if !a.review.Started() {
		a.review.Start()
	}
	if a.review.Expired() {
		return
	}
	a.notifier.Trace(a.Name(), task, observation, action)
}

func (a *agentT) cacheable(r *http.Request) bool {
	if a.state.Host == "" || r.Method != http.MethodGet || httpx.CacheControlNoCache(r.Header) {
		return false
	}
	return a.state.Enabled.Load()
}

func (a *agentT) emissaryShutdown() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) cacheUpdate(url string, r *http.Request, resp *http.Response) error {
	var (
		buf    []byte
		err    error
		status *std.Status
	)
	// TODO: Need to reset the body in the response after reading it.
	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		status = std.NewStatus(std.StatusIOError, a.Name(), err)
		a.notifier.Message(messaging.NewStatusMessage(status, a.Name()))
		return err
	}
	resp.ContentLength = int64(len(buf))
	resp.Body = io.NopCloser(bytes.NewReader(buf))

	// cache update
	go func() {
		h2 := httpx.CloneHeader(resp.Header)
		h2.Add(httpx.XRequestId, r.Header.Get(httpx.XRequestId))
		_, status = do(a, http.MethodPut, url, h2, io.NopCloser(bytes.NewReader(buf)))
		if status.Err != nil {
			a.notifier.Message(messaging.NewStatusMessage(status, a.Name())) //.WithLocation(a.Name()), a.Name()))
		}
	}()
	return nil
}

/*
func (a *agentT) configure2(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		cfg, status := messaging.MapContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.state.Update(cfg)
	case rest.ContentTypeExchange:
		ex, status := rest.ExchangeContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.exchange = ex
	case messaging.ContentTypeReview:
		r, status := messaging.ReviewContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		a.review = r
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}


*/
