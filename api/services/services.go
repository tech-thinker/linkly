package services

import (
	"github.com/tech-thinker/linkly/api/controllers"
	"github.com/tech-thinker/linkly/database"
	"github.com/tech-thinker/linkly/repository"
)

type Services interface {
	URLService() controllers.URL
	HealthCheckService() controllers.HealthCheck
}

type services struct {
	url         controllers.URL
	healthCheck controllers.HealthCheck
}

func (svc *services) URLService() controllers.URL {
	return svc.url
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

// NewServices initializes services
func NewServices() Services {
	db := database.GetDB()
	return &services{
		url: controllers.NewURL(
			repository.NewURLRepo(db),
		),
		healthCheck: controllers.NewHealthCheck(),
	}
}
