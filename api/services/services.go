package services

import (
	"log"

	"github.com/tech-thinker/linkly/api/controllers"
	"github.com/tech-thinker/linkly/config"
	"github.com/tech-thinker/linkly/repository"
	"github.com/tech-thinker/linkly/storage/redis"
)

// Services interface of service
type Services interface {
	URLService() controllers.URL
}

type services struct {
	url controllers.URL
}

// URLService is a service for url
func (svc *services) URLService() controllers.URL {
	return svc.url
}

// NewServices returns new instance of service
func NewServices() Services {
	config, err := config.FromFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	service, err := redis.New(config.Redis.Host, config.Redis.Port, config.Redis.Password)
	if err != nil {
		log.Fatal(err)
	}

	return &services{url: controllers.NewURL(
		repository.NewURLRepository(service),
	)}
}
