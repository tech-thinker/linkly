package services

import (
	"github.com/tech-thinker/linkly/api/controllers"
	"github.com/tech-thinker/linkly/repository"
	"github.com/tech-thinker/linkly/storage"
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
	db := storage.GetDB()
	return &services{url: controllers.NewURL(
		repository.NewURLRepository(db),
	)}
}
