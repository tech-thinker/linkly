package routes

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	jwtMiddleware "github.com/mrinjamul/mrinjamul-auth/middleware"
	"github.com/tech-thinker/linkly/api/services"
	"github.com/tech-thinker/linkly/docs"
	"github.com/tech-thinker/linkly/middleware"

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

	// URL Redirect
	router.GET("/:link", func(c *gin.Context) {
		svc.URLService().Redirect(c)
	})
	// Generate QR Code for the short url
	router.GET("/:link/qrcode", func(c *gin.Context) {
		svc.URLService().GenQR(c)
	})

	// Backend API
	docs.SwaggerInfo.BasePath = "/"
	api := router.Group("/api")
	api.Use(middleware.CORSMiddleware())
	v1 := api.Group("/v1")
	{
		// health check
		api.GET("/health", func(c *gin.Context) {
			svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
		})

		// links routes
		links := v1.Group("/links")
		{
			// link routes
			links.GET("", jwtMiddleware.JWTAuthAdmin(), func(c *gin.Context) {
				svc.LinkService().GetLinks(c)
			})
			links.POST("", jwtMiddleware.JWTAuth(), func(c *gin.Context) {
				svc.LinkService().AddLink(c)
			})
			links.GET("/:id", func(c *gin.Context) {
				svc.LinkService().GetLink(c)
			})
			links.GET("/:id/qrcode", func(c *gin.Context) {
				svc.LinkService().GenQRCode(c)
			})
			links.PATCH("/:id", jwtMiddleware.JWTAuth(), func(c *gin.Context) {
				svc.LinkService().UpdateLink(c)
			})
			links.DELETE("/:id", jwtMiddleware.JWTAuth(), func(c *gin.Context) {
				svc.LinkService().DeleteLink(c)
			})
			links.GET("/:id/stats", func(c *gin.Context) {
				svc.LinkService().GetLinkStats(c)
			})
		}

		// domains routes
		domains := v1.Group("/domains")
		{
			// domain routes
			domains.GET("", func(c *gin.Context) {
				svc.DomainService().GetDomains(c)
			})
		}

		// tracker routes
		tracker := v1.Group("/trackers")
		{
			// tracker routes
			tracker.GET("", func(c *gin.Context) {
				svc.TrackerService().GetTrackers(c)
			})
			tracker.GET("/gen", func(c *gin.Context) {
				svc.TrackerService().GenerateTracker(c)
			})
			tracker.GET("/:id", func(c *gin.Context) {
				svc.TrackerService().GetTracker(c)
			})
			tracker.GET("/:id/qr.png", func(c *gin.Context) {
				svc.TrackerService().QRCode(c)
			})
			tracker.GET("/:id/status", func(c *gin.Context) {
				svc.TrackerService().Status(c)
			})
			tracker.DELETE("/:id", func(c *gin.Context) {
				svc.TrackerService().DeleteTracker(c)
			})
		}
	}

	// swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
