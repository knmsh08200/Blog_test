package middleware

import (
	"net/http"
	"strings"
)

const BearerPrefix = "Bearer "

// BearerTokenMiddleware проверяет наличие и валидность Bearer Token
func BearerTokenMiddleware(next http.Handler, validToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, BearerPrefix)
		if token != validToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Валидация прошла успешно, продолжаем обработку запроса
		next.ServeHTTP(w, r)
	})
}
