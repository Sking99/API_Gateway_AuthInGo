package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseURL string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseURL)
	if err != nil {
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		originalPath := r.URL.Path

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)

		fmt.Println("Original Path:", r.URL.Path, r.URL.Host, r.Host)
		fmt.Println("Stripped Path:", target.Host, target.Path)
		r.URL.Host = target.Host
		r.URL.Path = target.Path + strippedPath
		r.Host = target.Host

		fmt.Println("Updated Path:", r.URL.Path, r.URL.Host, r.Host)

		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-User-Id", userId)
		}
	}

	return proxy.ServeHTTP
}
