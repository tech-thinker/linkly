package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
	"gorm.io/gorm"
)

type Tracker interface {
	GenerateTracker(ctx *gin.Context, track *models.Tracker) error
	GetTrackers(ctx *gin.Context) ([]models.Tracker, error)
	GetTracker(ctx *gin.Context, track models.Tracker) (models.Tracker, error)
	Update(ctx *gin.Context, track *models.Tracker) error
	Delete(ctx *gin.Context, track *models.Tracker) error
}

type tracker struct {
	db gorm.DB
}

// GenerateTracker generates a new tracker
func (repo *tracker) GenerateTracker(ctx *gin.Context, track *models.Tracker) error {
	// check if tracker is valid
	if track.ID == "" {
		return errors.New("id is empty")
	}

	// check if custom short link already exists
	existingLink := models.Tracker{}
	result := repo.db.Find(&existingLink, "id = ?", track.ID)
	if result.Error != nil {
		return result.Error
	}
	if existingLink.ID != "" {
		return errors.New("tracker already exists")
	}

	// Make transaction to add tracker
	tx := repo.db.Begin()
	if tx.Create(&track).Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// GetTrackers returns all trackers
func (repo *tracker) GetTrackers(ctx *gin.Context) ([]models.Tracker, error) {
	trackers := []models.Tracker{}
	if err := repo.db.Find(&trackers).Error; err != nil {
		return nil, err
	}
	return trackers, nil
}

// GetTracker returns a tracker
func (repo *tracker) GetTracker(ctx *gin.Context, track models.Tracker) (models.Tracker, error) {
	var tracker models.Tracker
	if err := repo.db.Find(&tracker, "id = ?", track.ID).Error; err != nil {
		return tracker, err
	}
	return tracker, nil
}

// Update updates a tracker
func (repo *tracker) Update(ctx *gin.Context, track *models.Tracker) error {
	// check if tracker is valid
	if track.ID == "" {
		return errors.New("id is empty")
	}

	// Make transaction to update tracker
	tx := repo.db.Begin()
	if tx.Save(&track).Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// Delete removes a tracker
func (repo *tracker) Delete(ctx *gin.Context, track *models.Tracker) error {
	// check if tracker is valid
	if track.ID == "" {
		return errors.New("id is empty")
	}

	// Make transaction to delete tracker
	tx := repo.db.Begin()
	if tx.Delete(&track).Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// NewTracker returns a new tracker
func NewTracker(db *gorm.DB) Tracker {
	return &tracker{db: *db}
}
