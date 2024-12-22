package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/config"
	"github.com/justin-jiajia/note/database"
	"github.com/justin-jiajia/note/router"
	"os"
)

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
	r := router.NewRouter(db).InitRouter()

	// Start server
	r.Run(":" + cfg.ServerPort)
}
