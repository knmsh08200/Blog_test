package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	ProcItemStatusFailed = "failed"
	ProcItemStatusDone   = "done"
)

var (
	reqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "myapp",
		Subsystem: "api",
		Name:      "http_total_request",
		Help:      "Общее количесвто HTTP-запросов ",
	}, []string{"status"})
	requestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: []float64{0.01, 0.05}, // 90-й и 95-й перцентили
		})
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

		reqCounter.WithLabelValues(strconv.Itoa(status)).Inc()
		requestDuration.Observe(duration)
	})
}
