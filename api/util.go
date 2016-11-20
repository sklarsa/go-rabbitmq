package api

import (
	"net/url"
)

func escapeVhost(vhost string) string {
	if vhost == "/" {
		return "%2f"
	}

	return url.QueryEscape(vhost)
}
