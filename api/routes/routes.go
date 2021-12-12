package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/api/services"
)

// New returns a new router
func New(routes *gin.Engine) {
	// init services
	svc := services.NewServices()
	// routers
	routes.LoadHTMLGlob("templates/*")
	routes.GET("/", func(c *gin.Context) {
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
	routes.GET("/api/health", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"health": "ok",
			},
		)
	})
	routes.GET("/api/links", func(c *gin.Context) {
		svc.URLService().FindAll(c)
	})
	routes.POST("/api/links", func(c *gin.Context) {
		svc.URLService().Add(c)
	})
	routes.GET("/api/links/:id", func(c *gin.Context) {
		svc.URLService().Find(c)
	})
	routes.DELETE("/api/links/:id", func(c *gin.Context) {
		svc.URLService().Delete(c)
	})
	routes.PATCH("/api/links/:id", func(c *gin.Context) {
		svc.URLService().Update(c)
	})
	routes.GET("/api/links/:id/stats", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"stats": "not implemented",
			},
		)
	})
}
