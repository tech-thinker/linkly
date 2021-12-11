package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/api/handler"
)

// New returns a new router
func New() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.Home())
	router.GET("/api/healthcheck", handler.HealthCheck())
	router.POST("/api/links", handler.CreateLink())
	router.GET("/api/links", handler.GetLinks())
	router.GET("/api/links/:id", handler.GetLink())
	router.DELETE("/api/links/:id", handler.DeleteLink())
	router.PATCH("/api/links/:id", handler.UpdateLink())
	router.GET("/api/links/:id/stats", handler.GetLinkStats())
	return router
}
