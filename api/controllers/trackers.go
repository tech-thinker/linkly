package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
)

type Trackers interface {
	GenerateTracker(ctx *gin.Context)
	GetTrackers(ctx *gin.Context)
	GetTracker(ctx *gin.Context)
	Status(ctx *gin.Context)
}

type tracker struct {
}

// GenerateTracker generates a new tracker
// @Summary Generate a new tracker
// @Description Generate a new tracker
// @ID generate-tracker
// @Tags trackers
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Tracker
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers [post]
func (t *tracker) GenerateTracker(ctx *gin.Context) {
	var tracker models.Tracker
	ctx.JSON(http.StatusOK, tracker)
}

// GetTrackers returns all trackers
// @Summary Get all trackers
// @Description Get all trackers
// @ID get-trackers
// @Tags trackers
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tracker
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers [get]
func (t *tracker) GetTrackers(ctx *gin.Context) {
	var trackers []models.Tracker
	ctx.JSON(http.StatusNotImplemented, trackers)
}

// GetTracker returns a tracker
// @Summary Get a tracker
// @Description Get a tracker
// @ID get-tracker
// @Tags trackers
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Tracker
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers/{id} [get]
func (t *tracker) GetTracker(ctx *gin.Context) {
	var tracker models.Tracker
	ctx.JSON(http.StatusNotImplemented, tracker)
}

// Status returns the status of the tracker
// @Summary Get the status of the tracker
// @Description Get the status of the tracker
// @ID get-tracker-status
// @Tags trackers
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TrackerStatus
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers/{id}/status [get]
func (t *tracker) Status(ctx *gin.Context) {
	var trackerStatus models.TrackerStatus
	ctx.JSON(http.StatusNotImplemented, trackerStatus)
}

// NewTrackers returns a new trackers controller
func NewTrackers() Trackers {
	return &tracker{}
}
