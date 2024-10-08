package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/knmsh08200/Blog_test/internal/blog"
	"github.com/knmsh08200/Blog_test/internal/db"
	"github.com/knmsh08200/Blog_test/internal/handlers"
	"github.com/knmsh08200/Blog_test/internal/metrics"
	"github.com/knmsh08200/Blog_test/internal/router"
)

func main() {

	ctx := context.Background()

	ctxWithCancel, cancelFunc := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunc()
	}()
	// обработка сигналов остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go metrics.InitProvider(":8082")

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	// Initialize DB PROVIDER
	dbProvider, err := db.InitDB(ctxWithCancel, dbUrl)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	defer dbProvider.Close()

	blogDBListProvider := blog.NewRep(dbProvider)
	blogDBIDProvider := blog.NewRep(dbProvider)
	if err != nil {
		log.Fatal(err)
	}

	blogListProvider, blogIDProvider := handlers.NewBlogHandler(blogDBListProvider, blogDBIDProvider)

	// Initialize route handler
	handler := router.NewHandler(blogListProvider, blogIDProvider)
	server := &http.Server{
		Addr:    ":3001",
		Handler: handler,
	}
	go func() {
		// Start the server
		log.Println("Server is listening on port 3001...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start,%v", err)
		}
	}()

	<-stop

	// Создаем контекст с таймаутом для корректного завершения работы сервера
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	// Завершаем работу сервера
	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
