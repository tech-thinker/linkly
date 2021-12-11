package main

import (
	"log"

	"github.com/tech-thinker/linkly/api/router"
	"github.com/tech-thinker/linkly/config"
	"github.com/tech-thinker/linkly/storage/redis"
)

func main() {
	config, err := config.FromFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	service, err := redis.New(config.Redis.Host, config.Redis.Port, config.Redis.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()
	// create new router
	router := router.New()
	// serve rest api
	log.Fatal(router.Run(":" + config.Server.Port))
}
