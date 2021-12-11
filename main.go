package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/api/router"
	"github.com/tech-thinker/linkly/config"
)

func main() {
	config, err := config.FromFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	// Create gin server
	server := gin.Default()

	router.New(server)
	// serve rest api
	log.Fatal(server.Run(":" + config.Server.Port))

	// service, err := redis.New(config.Redis.Host, config.Redis.Port, config.Redis.Password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer service.Close()

}
