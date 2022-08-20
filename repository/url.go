package repository

import (
	"errors"
	"time"

	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/utils"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type URL interface {
	Redirect(ctx *gin.Context, link models.Link) (models.Link, error)
	GenQR(ctx *gin.Context, qr models.QRCode) (models.QRCode, error)
}

type url struct {
	db gorm.DB
}

// Redirect redirects to the target url
func (repo *url) Redirect(ctx *gin.Context, link models.Link) (models.Link, error) {
	// check if link is valid
	if link.Link == nil {
		return link, errors.New("link is empty")
	}

	// retrieve link from database
	result := repo.db.Find(&link, "link = ?", link.Link)
	if result.Error != nil {
		return link, result.Error
	}
	if link.ID == "" {
		return link, errors.New("link not found")
	}
	// update link's click count
	link.VisitCount++
	link.UpdatedAt = time.Now()

	// make transaction to update link
	tx := repo.db.Begin()
	if tx.Error != nil {
		return link, tx.Error
	}
	if err := tx.Save(&link).Error; err != nil {
		tx.Rollback()
		return link, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return link, err
	}

	return link, nil
}

// GenQR generate qr code
func (u *url) GenQR(ctx *gin.Context, qr models.QRCode) (models.QRCode, error) {
	var err error
	// generate qr code
	qr.Image, err = utils.GenerateQRCode(qr.Content)
	if err != nil {
		return qr, err
	}
	return qr, nil
}

func NewURLRepo(db *gorm.DB) URL {
	return &url{db: *db}
}
