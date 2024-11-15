package middleware

import (
	"auth-api/utils"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromHeader(r)
		if tokenString == "" {
			utils.ErrorResponse(w, "Authorization header missing or invalid", http.StatusUnauthorized)
			return
		}

		if utils.IsTokenBlacklisted(tokenString) {
			utils.ErrorResponse(w, "Token has been revoked", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.ErrorResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			utils.ErrorResponse(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
