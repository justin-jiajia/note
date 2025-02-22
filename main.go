package main

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/config"
	"github.com/justin-jiajia/note/database"
	"github.com/justin-jiajia/note/router"
)

//go:embed front/dist/*
var frontFS embed.FS

func main() {

	// Initialize configuration
	cfg := config.NewConfig()

	// Set gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	db := database.InitDB(cfg)

	// Initialize router
	r := router.NewRouter(db, &frontFS).InitRouter(cfg)

	// Start server
	r.Run(":" + cfg.ServerPort)
}
