package utils

import "net/http"

func isHTTPS(r *http.Request) bool {
	if r.TLS != nil {
		return true
	}

	// 2. Se estiver atrás de um proxy reverso (Nginx, Cloudflare, etc.)
	if r.Header.Get("X-Forwarded-Proto") == "https" {
		return true
	}

	return false
}
