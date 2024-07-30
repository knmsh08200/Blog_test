package main

import (
	"log"
	"net/http"

	storage "github.com/knmsh08200/Blog_test/internal/db"
	"github.com/knmsh08200/Blog_test/internal/metrics"
	routes "github.com/knmsh08200/Blog_test/internal/router"
)

func main() {
	// Initialize observability (metrics)
	go metrics.Init(":8082")

	// Initialize DB provider
	err := storage.InitDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer storage.Close()

	// Initialize route handler
	mux := routes.NewRouter()
	handler := metrics.MetricsMiddleware(mux)

	// Start the server
	log.Println("Server is listening on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", handler))

}
