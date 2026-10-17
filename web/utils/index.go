package utils

import (
	"net"
	"net/http"
	"strings"
)

/*
 Nota
 Para restaurar o endereço IP do visitante original
 no seu servidor web de origem, a Cloudflare recomenda que
 seus logs ou aplicativos sejam analisados
 CF-Connecting-IP ou True-Client-IP em vez de
 X-Forwarded-For. CF-Connecting-IP e True-Client-IP
 ambos possuem um formato consistente contendo apenas
 um endereço IP.
*/

func IsHTTPS(r *http.Request) bool {
	if r.TLS != nil {
		return true
	}

	// Se estiver atrás de um proxy reverso (Nginx, Cloudflare, etc.)
	if r.Header.Get("X-Forwarded-Proto") == "https" {
		return true
	}

	return false
}

func GetClientIP(r *http.Request) string {
	// 1. Sua preferência: Tenta pegar direto pelo CF-Connecting-IP
	cfIP := r.Header.Get("CF-Connecting-IP")
	if cfIP != "" {
		if net.ParseIP(strings.TrimSpace(cfIP)) != nil {
			return strings.TrimSpace(cfIP)
		}
	}

	// 2. Segundo lugar: X-Forwarded-For (caso venha de outro proxy)
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		// O cabeçalho pode vir como uma lista: "IP_Cliente, IP_Proxy1, IP_Proxy2"
		ips := strings.Split(forwardedFor, ",")
		clientIP := strings.TrimSpace(ips[0])
		if net.ParseIP(clientIP) != nil {
			return clientIP
		}
	}

	// 3. Último caso: IP direto da conexão TCP (RemoteAddr)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return ip
	}

	return r.RemoteAddr
}
