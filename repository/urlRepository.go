package repository

import (
	"errors"
	"time"

	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/utils"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type URLRepo interface {
	Add(ctx *gin.Context, url *models.URL) error
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
func (repo *urlRepo) Add(ctx *gin.Context, url *models.URL) error {
	// check if url is valid
	if url.URL == "" {
		return errors.New("url is empty")
	}

	expires := url.Expires
	// check if url is already exists
	result := repo.db.Find(&url, "url = ?", url.URL)
	if result.Error != nil {
		return result.Error
	}
	if url.ID > 0 {
		// url expiration removed
		// expires, _ := time.Parse("2006-01-02", url.Expires)
		// if time.Now().After(expires) {
		// 	url.Expires = time.Now().AddDate(0, 0, 30).Format("2006-01-02")
		// }
		url.Expires = expires
		go func() { _ = repo.db.Omit("ID").Save(&url) }()
		return errors.New("url already exists")
	}
	// check if custom short url already exists
	if url.ShortURL != "" {
		existingURL := models.URL{}
		result = repo.db.Find(&existingURL, "short_url = ?", url.ShortURL)
		if result.Error != nil {
			return result.Error
		}
		if url.ShortURL == existingURL.ShortURL {
			return errors.New("short url already exists")
		}
	}
	// generate short url 62^7
	if url.ShortURL == "" {
		url.ShortURL = utils.RandomChars(7)
	}

	// url expiration removed
	// // add expires in 30 days
	// expires := time.Now().AddDate(0, 0, 30).Format("2006-01-02")
	// url.Expires = expires

	result = repo.db.Omit("Visits", "CreatedAt").Create(&url)
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
	// compare time if expired
	if url.Expires != "" {
		expires, _ := time.Parse("2006-01-02", url.Expires)
		if time.Now().After(expires) {
			return models.URL{}, errors.New("url has been expired")
		}
	}
	return url, nil
}

// GetAndRedirect a url by short url
func (repo *urlRepo) GetAndRedirect(ctx *gin.Context, url models.URL) (models.URL, error) {
	result := repo.db.Find(&url, "short_url = ?", url.ShortURL)
	if result.Error != nil {
		return url, result.Error
	}
	// compare time if expired
	if url.Expires != "" {
		expires, _ := time.Parse("2006-01-02", url.Expires)
		if time.Now().After(expires) {
			return models.URL{}, errors.New("url has been expired")
		}
	}
	// increase visits field by one and save it
	// url.Visits++
	// result = repo.db.Save(&url)
	go func() {
		// increase visits field by one and save it
		url.Visits++
		if url.URL != "" {
			_ = repo.db.Save(&url)
		}
	}()
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
	var existingURL models.URL
	result := repo.db.Find(&existingURL, "short_url = ?", url.ShortURL)
	if result.Error != nil {
		return result.Error
	}
	existingURL.Expires = url.Expires
	// exclude unnecessary fields and update
	result = repo.db.Omit("UpdatedAt").Save(&existingURL)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete a url
func (repo *urlRepo) Delete(ctx *gin.Context, url *models.URL) error {
	result := repo.db.Delete(&url, "short_url = ?", url.ShortURL)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewURLRepo(db *gorm.DB) URLRepo {
	return &urlRepo{db: *db}
}
