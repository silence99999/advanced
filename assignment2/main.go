package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"assignment2/internal/api"
	"assignment2/internal/model"
	"assignment2/internal/queue"
	"assignment2/internal/store"
	"assignment2/internal/worker"
)

func main() {
	repo := store.NewRepository[string, *model.Task]()
	taskQueue := queue.NewTaskQueue()

	go worker.Start(1, taskQueue, repo)
	go worker.Start(2, taskQueue, repo)

	stopMonitor := make(chan struct{})
	go startMonitoring(repo, stopMonitor)

	apiHandler := api.New(repo, taskQueue)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", apiHandler.CreateTaskHandler)
	mux.HandleFunc("GET /tasks", apiHandler.GetTasksHandler)
	mux.HandleFunc("GET /task/{id}", apiHandler.GetTaskHandler)
	mux.HandleFunc("GET /stats", apiHandler.GetStatsHandler)
	mux.HandleFunc("GET /slow", apiHandler.SlowHandler)
	server := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Println("server started on http://localhost:9000")
		server.ListenAndServe()
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")

	close(stopMonitor)
	taskQueue.Close()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)

	log.Println("server stopped gracefully")
}

func startMonitoring(repo *store.Repository[string, *model.Task], stop <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			stats := api.CountStats(repo)

			log.Printf("MONITOR submitted=%d in_progress=%d completed=%d", stats.Submitted, stats.InProgress, stats.Completed)

		case <-stop:
			log.Println("MONITOR stopped")
			return
		}
	}
}
