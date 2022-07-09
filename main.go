package main

import (
	"embed"
	"log"
	"os"
	"time"

	"github.com/tech-thinker/linkly/api/routes"

	"github.com/gin-gonic/gin"
)

//go:embed views/*
var viewsFs embed.FS

var (
	// startTime is the time when the server starts
	startTime time.Time = time.Now()
)

func main() {
	// Get port from env
	port := ":3000"
	_, present := os.LookupEnv("PORT")
	if present {
		port = ":" + os.Getenv("PORT")

	}

	// Set the router as the default one shipped with Gin
	server := gin.Default()
	// Initialize the routes
	routes.StartTime = startTime
	routes.ViewsFs = viewsFs
	routes.InitRoutes(server)
	routes.BootTime = time.Since(startTime)
	// Start and run the server
	log.Fatal(server.Run(port))
}
