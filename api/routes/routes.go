package routes

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/mrinjamul/gnote/middleware"
	"github.com/tech-thinker/linkly/api/services"
	"github.com/tech-thinker/linkly/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ViewsFs for static files
var ViewsFs embed.FS

var (
	StartTime time.Time
	BootTime  time.Duration
)

func InitRoutes(routes *gin.Engine) {
	svc := services.NewServices()
	// Serve the frontend
	// This will ensure that the files are served correctly
	fsRoot, err := fs.Sub(ViewsFs, "views")
	if err != nil {
		log.Println(err)
	}
	routes.NoRoute(gin.WrapH(http.FileServer(http.FS(fsRoot))))

	// Backend API
	docs.SwaggerInfo.BasePath = "/"

	// redirect route
	routes.GET("/:short_url", func(c *gin.Context) {
		svc.URLService().GetAndRedirect(c)
	})
	routes.GET("/:short_url/qr", func(c *gin.Context) {
		svc.URLService().GenQR(c)
	})
	// api routes group
	api := routes.Group("/api")
	// v1 := api.Group("/v1")
	// api.Use(middleware.CORSMiddleware())
	{
		// health check
		api.GET("/health", func(c *gin.Context) {
			svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
		})
		// links routes group
		links := api.Group("/links")
		{
			// link routes
			links.GET("", func(c *gin.Context) {
				svc.URLService().GetAll(c)
			},
			)
			links.GET("/:short_url", func(c *gin.Context) {
				svc.URLService().Get(c)
			},
			)
			links.GET("/:short_url/qr", func(c *gin.Context) {
				svc.URLService().GenQR(c)
			},
			)
			links.POST("", func(c *gin.Context) {
				svc.URLService().Add(c)
			},
			)
			links.PATCH("", middleware.JWTAuth(), func(c *gin.Context) {
				svc.URLService().Update(c)
			},
			)
			links.DELETE("", middleware.JWTAuth(), func(c *gin.Context) {
				svc.URLService().Delete(c)
			},
			)
		}
		// domains routes group
		domains := api.Group("/domains")
		{
			// domains routes
			domains.GET("", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{
					"message": "success",
					"domains": "",
				})

			})
		}
		// tracker
		tracker := api.Group("/trackers")
		{
			// tracker routes
			tracker.GET("", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{
					"message":  "success",
					"trackers": "",
				})
			})
			tracker.GET("/gen", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "success",
					"url":     "",
				})
			})
		}
	}
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
