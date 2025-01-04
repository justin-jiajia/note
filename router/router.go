package router

import (
	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/docs"
	_ "github.com/justin-jiajia/note/docs"
	"github.com/justin-jiajia/note/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

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
		api.POST("/notes", handler.CreateNote(r.db))
	}

	docs.SwaggerInfo.BasePath="/api/v1"
	docs.SwaggerInfo.Title="Note API"
	docs.SwaggerInfo.Version="1.0"
	docs.SwaggerInfo.Description="A simple note-taking API"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
