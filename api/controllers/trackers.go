package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/repository"
	"github.com/tech-thinker/linkly/utils"
)

type Trackers interface {
	GenerateTracker(ctx *gin.Context)
	GetTrackers(ctx *gin.Context)
	GetTracker(ctx *gin.Context)
	QRCode(ctx *gin.Context)
	Status(ctx *gin.Context)
	DeleteTracker(ctx *gin.Context)
}

type trackers struct {
	track repository.Tracker
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
// @Router /api/v1/trackers/gen [get]
func (t *trackers) GenerateTracker(ctx *gin.Context) {
	tracker := models.Tracker{}
	var err error

	tracker.ID = utils.GenerateShortURL()
	ip := ctx.ClientIP()
	tracker.IP = &ip
	hostname := ctx.Request.Host
	url := "http://" + hostname + "/api/v1/trackers/" + tracker.ID + "/qr.png"
	tracker.URL = url

	// Save tracker to database
	err = t.track.GenerateTracker(ctx, &tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to save tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

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
func (t *trackers) GetTrackers(ctx *gin.Context) {
	trackers, err := t.track.GetTrackers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get trackers",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, trackers)
}

// GetTracker returns a tracker
// @Summary Get a tracker
// @Description Get a tracker
// @ID get-tracker
// @Tags trackers
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Tracker
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers/{id} [get]
func (t *trackers) GetTracker(ctx *gin.Context) {
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, tracker)
}

// QRCode returns a qr code png for a tracker
// @Summary Get a qr code png for a tracker
// @Description Get a qr code png for a tracker
// @ID get-qr-code
// @Tags trackers
// @Accept  json
// @Produce  data:image/png
// @Success 200 {object} models.Tracker.Image
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers/{id}/qr.png [get]
func (t *trackers) QRCode(ctx *gin.Context) {
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	tracker.VisitCount++
	err = t.track.Update(ctx, &tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to update tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	image := utils.GenerateTrackerImage()
	ctx.Data(http.StatusOK, "image/png", image)
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
func (t *trackers) Status(ctx *gin.Context) {
	var trackerStatus models.TrackerStatus
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	trackerStatus.ID = tracker.ID
	trackerStatus.URL = tracker.URL
	trackerStatus.Message = "not seen"
	if tracker.VisitCount > 0 {
		trackerStatus.Seen = true
		trackerStatus.Message = "seen"
	}

	ctx.JSON(http.StatusOK, trackerStatus)
}

// DeleteTracker deletes a tracker
// @Summary Delete a tracker
// @Description Delete a tracker
// @ID delete-tracker
// @Tags trackers
// @Accept  json
// @Produce  json
// @Success 204 {object} models.Message
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/trackers/{id} [delete]
func (t *trackers) DeleteTracker(ctx *gin.Context) {
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	err = t.track.Delete(ctx, &tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to delete tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	msg := models.Message{
		Code:    http.StatusOK,
		Message: "tracker deleted",
	}
	ctx.JSON(http.StatusOK, msg)
}

// NewTrackers returns a new trackers controller
func NewTrackers(tracker repository.Tracker) Trackers {
	return &trackers{
		track: tracker,
	}
}
