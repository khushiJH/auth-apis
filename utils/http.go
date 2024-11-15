package utils

import (
	"net/http"
	"strings"
)

func GetTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return ""
	}

	return parts[1]
}
