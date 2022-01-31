package main

import (
	"embed"
	"log"
	"os"

	"github.com/tech-thinker/linkly/api/routes"

	"github.com/gin-gonic/gin"
)

//go:embed views/*
var viewsFs embed.FS

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
	routes.ViewsFs = viewsFs
	routes.InitRoutes(server)
	// Start and run the server
	log.Fatal(server.Run(port))
}
