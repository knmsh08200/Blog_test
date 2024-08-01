package metrics

import (
	"log"
	"net/http"

	"github.com/knmsh08200/Blog_test/internal/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(addr string) {
	prometheus.MustRegister(
		router.ReqCounter,
		router.RequestDuration,
	)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Println("Listen port:8082")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Printf("Error:%v", err)
	}
}
