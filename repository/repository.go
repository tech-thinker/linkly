package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/storage"
)

type URLRepository interface {
	Add(ctx *gin.Context, url *models.URL) error
	Find(ctx *gin.Context, url models.URL) (models.URL, error)
	Update(ctx *gin.Context, url *models.URL) error
	Delete(ctx *gin.Context, url *models.URL) error
	FindAll(ctx *gin.Context) ([]models.URL, error)
}

type urlRepository struct {
	storage storage.Service
}

// Find returns a URL by its ID
func (repo *urlRepository) Find(ctx *gin.Context, url models.URL) (models.URL, error) {
	return models.URL{}, nil
}

// Add adds a new URL to the database
func (repo *urlRepository) Add(ctx *gin.Context, url *models.URL) error {
	return nil
}

// Update updates an existing URL
func (repo *urlRepository) Update(ctx *gin.Context, url *models.URL) error {
	return nil
}

// Delete deletes an existing URL
func (repo *urlRepository) Delete(ctx *gin.Context, url *models.URL) error {
	return nil
}

// FindAll returns all URLs
func (repo *urlRepository) FindAll(ctx *gin.Context) ([]models.URL, error) {
	return []models.URL{}, nil
}

// NewURLRepository returns a new URLRepository
func NewURLRepository(storage storage.Service) URLRepository {
	return &urlRepository{storage}
}
