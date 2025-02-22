package router

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/config"
	"github.com/justin-jiajia/note/docs"
	"github.com/justin-jiajia/note/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Router struct {
	db      *gorm.DB
	frontFS *embed.FS
}

func NewRouter(db *gorm.DB, frontFS *embed.FS) *Router {
	return &Router{db: db, frontFS: frontFS}
}

func (r *Router) InitRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{cfg.AllowOrigin}
	config.AllowHeaders = []string{"Content-Type", "X-Encryption-Tag"}
	router.Use(cors.New(config))

	// Serve embedded static files
	frontDist, _ := fs.Sub(r.frontFS, "front/dist/assets")
	router.StaticFS("/assets", http.FS(frontDist))

	// API routes group
	api := router.Group("/api/v1")
	{
		// Note routes
		api.GET("/notes/:slug", handler.ViewNote(r.db))
		api.PUT("/notes/:slug", handler.EditNote(r.db))
		api.DELETE("/notes/:slug", handler.DeleteNote(r.db))
		api.POST("/notes", handler.CreateNote(r.db))
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Note API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "A simple note-taking API"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Serve embedded index.html for all other routes
	router.NoRoute(func(c *gin.Context) {
		indexFile, _ := r.frontFS.ReadFile("front/dist/index.html")
		c.Data(http.StatusOK, "text/html", indexFile)
	})

	return router
}
