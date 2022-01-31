package services

import (
	"github.com/tech-thinker/linkly/api/controllers"
	"github.com/tech-thinker/linkly/db"
	"github.com/tech-thinker/linkly/repository"
)

type Services interface {
	URLService() controllers.URL
}

type services struct {
	url controllers.URL
}

func (svc *services) URLService() controllers.URL {
	return svc.url
}

// NewServices initializes services
func NewServices() Services {
	db := db.GetDB()
	return &services{
		url: controllers.NewURL(
			repository.NewURLRepo(db),
		),
	}
}
