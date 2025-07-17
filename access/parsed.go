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
	rawURL := p.Path
	if p.Query != "" {
		rawURL = p.Path + "?" + p.Query
	}
	u, _ := url.Parse(rawURL)
	return u
}

// parseURL - create the URL, host, path, and query, un-escaping path and query
func parseURL(hostOverride string, u *url.URL) (uri string, p *parsed) {
	if u == nil {
		uri = "url-is-nil"
		return uri, &parsed{Valid: false, Err: errors.New("invalid argument: URL is nil")}
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
		Valid:    true,
		Host:     host,
		Domain:   "",
		Version:  "",
		Resource: "",
		Path:     path,
		Query:    query,
		Err:      nil,
	}
}
