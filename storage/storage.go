package storage

import (
	"time"

	"github.com/tech-thinker/linkly/models"
)

// Service is the interface that provides the basic storage functionality.
type Service interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*models.URL, error)
	Close() error
}
