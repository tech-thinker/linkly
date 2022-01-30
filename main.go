package main

import (
	"embed"
	"log"

	"github.com/tech-thinker/linkly/api/routes"

	"github.com/gin-gonic/gin"
)

//go:embed views/*
var viewsFs embed.FS

func main() {
	// Set the router as the default one shipped with Gin
	server := gin.Default()
	// Initialize the routes
	routes.ViewsFs = viewsFs
	routes.InitRoutes(server)
	// Start and run the server
	log.Fatal(server.Run(":8080"))
}
