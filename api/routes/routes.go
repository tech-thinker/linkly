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

func InitRoutes(router *gin.Engine) {
	// Initialize the services
	svc := services.NewServices()

	// Serve the frontend
	fsRoot, err := fs.Sub(ViewsFs, "views")
	if err != nil {
		log.Println(err)
	}
	router.NoRoute(gin.WrapH(http.FileServer(http.FS(fsRoot))))

	// Backend API
	docs.SwaggerInfo.BasePath = "/"

	// URL Redirect
	router.GET("/:short_url", func(c *gin.Context) {
		svc.URLService().GetAndRedirect(c)
	})
	// Generate QR Code for the short url
	router.GET("/:short_url/qr", func(c *gin.Context) {
		svc.URLService().GenQR(c)
	})

	// API Routes
	api := router.Group("/api")
	// v1 := api.Group("/v1")
	// api.Use(middleware.CORSMiddleware())
	{
		// health check
		api.GET("/health", func(c *gin.Context) {
			svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
		})
		// links routes
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
		// domains routes
		domains := api.Group("/domains")
		{
			// domain routes
			domains.GET("", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{
					"message": "success",
					"domains": "",
				})

			})
		}
		// tracker routes
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
