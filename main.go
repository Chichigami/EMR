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
	defer c.Stop()

	_, err := c.AddFunc("6 * * * *", h.CacheToday)
	if err != nil {
		log.Fatalln(err.Error())
	}

	server := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

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
