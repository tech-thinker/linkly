package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/repository"
)

type URL interface {
	Redirect(ctx *gin.Context)
	Track(ctx *gin.Context)
	GenQR(ctx *gin.Context)
}

type url struct {
	url repository.URL
}

// Redirect redirects to the target url
// @Summary Redirect to the target url
// @Description Redirect to the target url
// @ID redirect
// @Tags URL Shortener
// @Accept  json
// @Produce  json
// @Success 302 {string} string
// @Failure 500 {object} models.Error
// @Router /{link} [get]
func (u *url) Redirect(ctx *gin.Context) {
	url := ctx.Param("link")
	link, err := u.url.Redirect(ctx, models.Link{Link: &url})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	if time.Since(*link.ExpireAt).Seconds() > 0 {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: "link has expired",
				Status: http.StatusInternalServerError,
			},
		})
		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	// Use http.StatusFound (302) to tell the http client to redirect
	target := *link.Target
	ctx.Redirect(http.StatusFound, target)
}

// Track tracks a url
// @Summary Track a url
// @Description Track a url
// @ID track
// @Tags URL Shortener
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Link
// @Failure 500 {object} models.Error
// @Router /{link}/track [get]
func (u *url) Track(ctx *gin.Context) {
	url := ctx.Param("link")

	ctx.JSON(http.StatusOK, models.Link{
		Link: &url,
	})
}

// GenQR generates a QR code for the short url
// @Summary Generate a QR code for the short url
// @Description Generate a QR code for the short url
// @ID generate-qr
// @Tags URL Shortener
// @Accept  json
// @Produce  json
// @Success 200 {object} models.QRCode.Image
// @Failure 500 {object} models.Error
// @Router /{link}/qrcode [get]
func (u *url) GenQR(ctx *gin.Context) {
	var qr models.QRCode
	qr.Content = ctx.Param("link")
	hostname := ctx.Request.Host
	qr.Content = "http://" + hostname + "/" + qr.Content
	qr, err := u.url.GenQR(ctx, qr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	ctx.Data(http.StatusOK, "image/png", qr.Image)
}

// NewURL returns a new url controller
func NewURL(urlRepo repository.URL) URL {
	return &url{
		url: urlRepo,
	}
}
