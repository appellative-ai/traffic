package access

import (
	"errors"
	"net/url"
	"strings"
)

type parsed struct {
	valid    bool
	host     string
	domain   string
	version  string
	resource string
	path     string
	query    string
	err      error
}

func (p *parsed) pathQuery() *url.URL {
	rawURL := p.path
	if p.query != "" {
		rawURL = p.path + "?" + p.query
	}
	u, _ := url.Parse(rawURL)
	return u
}

// parseURL - create the URL, host, path, and query, un-escaping path and query
func parseURL(hostOverride string, u *url.URL) (uri string, p *parsed) {
	if u == nil {
		uri = "url-is-nil"
		return uri, &parsed{valid: false, err: errors.New("invalid argument: URL is nil")}
	}
	// Set scheme
	scheme := u.Scheme
	if scheme == "" {
		scheme = "http"
	}
	// Set host
	host := hostOverride
	if len(host) == 0 {
		host = u.Host
	}
	// Set path, checking for a domain
	urlPath, _ := url.PathUnescape(u.Path)
	path := urlPath
	i := strings.Index(path, ":")
	if i >= 0 {
		path = path[i+1:]
	}

	// Set query
	query := ""
	if u.RawQuery != "" {
		query, _ = url.QueryUnescape(u.RawQuery)
	}
	if query != "" {
		uri = scheme + "://" + host + urlPath + "?" + query
	} else {
		uri = scheme + "://" + host + urlPath + query
	}
	return uri, &parsed{
		valid:    true,
		host:     host,
		domain:   "",
		version:  "",
		resource: "",
		path:     path,
		query:    query,
		err:      nil,
	}
}
