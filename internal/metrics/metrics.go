package metrics

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	ProcItemStatusFailed = "failed"
	ProcItemStatusDone   = "done"
)

var (
	ReqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "myapp",
		Subsystem: "api",
		Name:      "http_total_request",
		Help:      "Общее количесвто HTTP-запросов ",
	}, []string{"status"})
	RequestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: []float64{0.01, 0.05}, // 90-й и 95-й перцентили
		})

	CacheHitResponseDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "myapp",
			Subsystem: "cache",
			Name:      "cache_response_duration_seconds",
			Help:      "Среднее время ответа кэша в секундах.",
			Buckets:   prometheus.DefBuckets, // Используем стандартные бакеты
		})
	CacheMissResponseDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "myapp",
			Subsystem: "cache",
			Name:      "cache__miss_response_duration_seconds",
			Help:      "Среднее время ответа при промахе кэша в секундах.",
			Buckets:   prometheus.DefBuckets, // Используем стандартные бакеты
		})
)

func InitProvider(addr string) {
	prometheus.MustRegister(
		ReqCounter,
		RequestDuration,
		CacheHitResponseDuration,
		CacheMissResponseDuration,
	)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Println("Listen port:8082")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Printf("Error:%v", err)
	}
}

func ObserveCacheHit(duration float64) {
	CacheHitResponseDuration.Observe(duration)
}

func ObserveCacheMiss(duration float64) {
	CacheMissResponseDuration.Observe(duration)
}

func IncRequestCounter(status int) {
	ReqCounter.WithLabelValues(strconv.Itoa(status)).Inc()
}

// ObserveRequestDuration наблюдает за продолжительностью запроса
func ObserveRequestDuration(duration float64) {
	RequestDuration.Observe(duration)
}
