package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/api/services"
)

// New returns a new router
func New(router *gin.Engine) {
	// init services
	svc := services.NewServices()
	// routers
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"health": "ok",
			},
		)
	})
	router.GET("/api/links", func(c *gin.Context) {
		svc.URLService().FindAll(c)
	})
	router.POST("/api/links", func(c *gin.Context) {
		svc.URLService().Add(c)
	})
	router.GET("/api/links/:id", func(c *gin.Context) {
		svc.URLService().Find(c)
	})
	router.DELETE("/api/links/:id", func(c *gin.Context) {
		svc.URLService().Delete(c)
	})
	router.PATCH("/api/links/:id", func(c *gin.Context) {
		svc.URLService().Update(c)
	})
	router.GET("/api/links/:id/stats", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"stats": "not implemented",
			},
		)
	})
}
