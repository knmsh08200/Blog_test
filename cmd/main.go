package main

import (
	"log"
	"net/http"
	"os"

	"github.com/knmsh08200/Blog_test/internal/db"
	"github.com/knmsh08200/Blog_test/internal/metrics"
	"github.com/knmsh08200/Blog_test/internal/router"
)

func main() {
	// Initialize observability (metrics)
	go metrics.Init(":8082")

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	// Initialize DB PROVIDER
	if err := db.InitDB(dbUrl); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize route handler
	handler := router.NewHandler()

	// Start the server
	log.Println("Server is listening on port 3001...")
	if err := http.ListenAndServe(":3001", handler); err != nil {
		log.Fatalf("Server failed to start,%v", err)
	}
}
