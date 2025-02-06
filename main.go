package main

import (
	"github.com/chichigami/EMR/internal/config"
	"github.com/chichigami/EMR/internal/handlers"
	"github.com/chichigami/EMR/internal/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, dbConnection := config.LoadConfig()
	defer dbConnection.Close()

	h := handlers.NewHandlerConfig(cfg)

	router := gin.Default()

	routes.LoadRoutes(router, h)

	router.Run(":8000")
}
