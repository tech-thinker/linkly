package repository

import (
	"time"

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
	urlModel, err := repo.storage.LoadInfo(url.URL)
	if err != nil {
		return models.URL{}, err
	}
	return *urlModel, nil
}

// Add adds a new URL to the database
func (repo *urlRepository) Add(ctx *gin.Context, url *models.URL) error {
	layoutISO := "2006-01-02 15:04:05"
	expires, err := time.Parse(layoutISO, url.Expires)
	if err != nil {
		return err
	}
	_, err = repo.storage.Save(url.URL, expires)
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing URL
func (repo *urlRepository) Update(ctx *gin.Context, url *models.URL) error {
	expires := time.Time{}
	_, err := repo.storage.Update(url.URL, expires)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an existing URL
func (repo *urlRepository) Delete(ctx *gin.Context, url *models.URL) error {
	_, err := repo.storage.Delete(url.OriginalURL)
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns all URLs
func (repo *urlRepository) FindAll(ctx *gin.Context) ([]models.URL, error) {
	urls, err := repo.storage.LoadInfoAll()
	return []models.URL{}, nil
}

// NewURLRepository returns a new URLRepository
func NewURLRepository(storage storage.Service) URLRepository {
	return &urlRepository{storage}
}
