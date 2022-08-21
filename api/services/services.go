package services

import (
	"github.com/tech-thinker/linkly/api/controllers"
	"github.com/tech-thinker/linkly/database"
	"github.com/tech-thinker/linkly/repository"
)

type Services interface {
	HealthCheckService() controllers.HealthCheck
	LinkService() controllers.Link
	URLService() controllers.URL
	DomainService() controllers.Domain
	TrackerService() controllers.Trackers
}

type services struct {
	healthCheck controllers.HealthCheck
	link        controllers.Link
	url         controllers.URL
	domain      controllers.Domain
	trackers    controllers.Trackers
}

// HealthCheckService returns a health check service
func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

// LinkService returns a link service
func (svc *services) LinkService() controllers.Link {
	return svc.link
}

// URLService returns a url service
func (svc *services) URLService() controllers.URL {
	return svc.url
}

// DomainService returns a domain service
func (svc *services) DomainService() controllers.Domain {
	return svc.domain
}

// TrackerService returns a tracker service
func (svc *services) TrackerService() controllers.Trackers {
	return svc.trackers
}

// NewServices initializes services
func NewServices() Services {
	db := database.GetDB()
	return &services{
		healthCheck: controllers.NewHealthCheck(),
		link: controllers.NewLink(
			repository.NewLink(db),
		),
		url: controllers.NewURL(
			repository.NewURLRepo(db),
		),
		domain: controllers.NewDomain(),
		trackers: controllers.NewTrackers(
			repository.NewTracker(db),
		),
	}
}
