package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chichigami/EMR/internal/config"
	"github.com/chichigami/EMR/internal/handlers"
	"github.com/chichigami/EMR/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	//need to add cron that runs a db loader at 6 am?
	cfg, cleanup := config.LoadConfig()
	defer cleanup()

	h := handlers.NewHandlerConfig(cfg)

	router := gin.Default()
	routes.AddRoutes(router, h)

	c := cron.New()

	// This cron job refreshes the dashboard cache every start of day to ensure the latest appointments are included.
	//
	// Why always refresh?
	// - New appointments may be added after the last cache update but before the start of the day.
	// - Avoids checks and calls between Redis cache, the appointments table, and the dashboards table.
	// - Performance optimizations from comparison are uncertain, so a full refresh is simpler and reliable.
	//
	// If the cron job fails to register, the program exits with a fatal error.
	_, err := c.AddFunc("* * * * *", func() {
		h.CacheDashboard(time.Now())
	})
	if err != nil {
		log.Fatalln(err)
	}

	server := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	c.Start()
	defer c.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server running on :8000")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("ListenAndServe error: %v", err)
		}
	}()

	<-stop
	log.Println("Server shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server stopped cleanly")
}
