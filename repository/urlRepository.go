package repository

import (
	"github.com/tech-thinker/linkly/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type URLRepo interface {
	Add(ctx *gin.Context, url *models.Input) error
	Get(ctx *gin.Context, url models.URL) (models.URL, error)
	GetAndRedirect(ctx *gin.Context, url models.URL) (models.URL, error)
	GetAll(ctx *gin.Context) ([]models.URL, error)
	Update(ctx *gin.Context, url *models.URL) error
	Delete(ctx *gin.Context, url *models.URL) error
}

type urlRepo struct {
	db gorm.DB
}

// Add a new url
func (repo *urlRepo) Add(ctx *gin.Context, url *models.Input) error {
	var shortUrl models.URL
	shortUrl.URL = url.URL
	shortUrl.ShortURL = url.ShortURL
	shortUrl.Expires = url.Expires
	result := repo.db.Create(&shortUrl)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get a url by short url
func (repo *urlRepo) Get(ctx *gin.Context, url models.URL) (models.URL, error) {
	result := repo.db.Find(&url, "short_url = ?", url.ShortURL)
	if result.Error != nil {
		return url, result.Error
	}
	return url, nil
}

// GetAndRedirect a url by short url
func (repo *urlRepo) GetAndRedirect(ctx *gin.Context, url models.URL) (models.URL, error) {
	result := repo.db.Find(&url, "short_url = ?", url.ShortURL)
	if result.Error != nil {
		return url, result.Error
	}
	return url, nil
}

// GetAll all urls
func (repo *urlRepo) GetAll(ctx *gin.Context) ([]models.URL, error) {
	var urls []models.URL
	result := repo.db.Find(&urls)
	if result.Error != nil {
		return urls, result.Error
	}
	return urls, nil
}

// Update a url
func (repo *urlRepo) Update(ctx *gin.Context, url *models.URL) error {
	result := repo.db.Save(&url)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete a url
func (repo *urlRepo) Delete(ctx *gin.Context, url *models.URL) error {
	result := repo.db.Delete(&url)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewURLRepo(db *gorm.DB) URLRepo {
	return &urlRepo{db: *db}
}
