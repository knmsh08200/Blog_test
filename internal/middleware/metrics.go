package middleware

import (
	"net/http"
	"time"

	"github.com/knmsh08200/Blog_test/internal/metrics"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	// size       int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Пример использования функции RequestDuration в обработчике HTTP-запросов
func MetricsMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time
		start := time.Now()

		// Create a response writer that captures the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Record metrics
		duration := time.Since(start).Seconds()
		// path := r.URL.Path
		// method := r.Method
		status := rw.statusCode

		metrics.IncRequestCounter(status)
		metrics.ObserveRequestDuration(duration)
	})
}
