package router

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/handler"
	"gorm.io/gorm"
)

//go:embed front/build
var staticFiles embed.FS

type Router struct {
	db *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{db: db}
}

func (r *Router) InitRouter() *gin.Engine {
	router := gin.Default()

	// API routes group
	api := router.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		// Note routes
		api.GET("/notes/:slug", handler.ViewNote(r.db))
		api.PUT("/notes/:slug", handler.EditNote(r.db))
		api.DELETE("/notes/:slug", handler.DeleteNote(r.db))
	}

	// Serve frontend in production
	if gin.Mode() == gin.ReleaseMode {
			// Get the embedded frontend files
		frontend, err := fs.Sub(staticFiles, "front/build")
		if err != nil {
			panic(err)
		}
		
		// Serve the entire frontend for root and handle HTML5 history mode
		router.NoRoute(func(c *gin.Context) {
			c.FileFromFS(c.Request.URL.Path, http.FS(frontend))
		})
	}

	return router
}
