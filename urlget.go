package myhttp

import (
	"net/url"
	"strings"
)

type urlGet struct {
	url  string
	data url.Values
}

func (g *urlGet) parseUrl() string {
	if g.data == nil {
		return g.url
	} else {
		tmp := g.data.Encode()
		switch {
		case !strings.Contains(g.url, "?"):
			tmp = "?" + tmp
		case g.url[len(g.url)-1] == '?':
		case g.url[len(g.url)-1] != '&':
			tmp = "&" + tmp
		}
		return g.url + tmp
	}
}
