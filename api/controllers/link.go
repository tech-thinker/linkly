package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/repository"
	"github.com/tech-thinker/linkly/utils"
)

type Link interface {
	GetLinks(c *gin.Context)
	AddLink(c *gin.Context)
	GetLink(c *gin.Context)
	UpdateLink(c *gin.Context)
	DeleteLink(c *gin.Context)
	GetLinkStats(c *gin.Context)
	GenQRCode(c *gin.Context)
}

type link struct {
	link repository.Link
}

// GetLinks returns all links
// @Summary Get all links
// @Description Get list of links
// @ID get-all-links
// @Tags links
// @Produce json
// @Success 200 {array} models.Link
// @Failure 500 {object} models.Error
// @Router /api/v1/links [get]
func (l *link) GetLinks(c *gin.Context) {
	links, err := l.link.ReadAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, links)
}

// AddLink adds a new link
// @Summary Add a new link
// @Description Add a new link
// @ID add-new-link
// @Tags links
// @Accept json
// @Produce json
// @Param link body models.LinkBody true "Link"
// @Success 200 {object} models.Link
// @Failure 500 {object} models.Error
// @Router /api/v1/links [post]
func (l *link) AddLink(c *gin.Context) {
	var link models.LinkBody
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: err.Error(),
				Status: http.StatusBadRequest,
			},
		})
		return
	}
	if link.Target == "" {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: "Target is empty",
				Status: http.StatusBadRequest,
			},
		})
		return
	}

	newLink := models.Link{
		Target: &link.Target,
	}
	// check if link is already exists
	newLink, err := l.link.Read(c, newLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if newLink.ID != "" && time.Since(*newLink.ExpireAt).Seconds() < 0 {
		c.JSON(http.StatusConflict, models.Error{
			Error: models.ServiceError{
				Type:   "conflict",
				Title:  "Conflict",
				Detail: "Link already exists",
				Status: http.StatusConflict,
			},
		})
		return
	}

	// Convert link to model->Link
	if newLink.ID == "" {
		newLink.ID = utils.GenerateUUID()
	}
	newLink.Target = &link.Target
	if link.CustomURL != "" {
		url := c.Request.Host
		newLink.Link = &link.CustomURL
		url = url + "/" + link.CustomURL
		newLink.Address = &url
	} else {
		shortURL := utils.GenerateShortURL()
		newLink.Link = &shortURL
		url := c.Request.Host
		url = url + "/" + shortURL
		newLink.Address = &url
	}
	if link.Reusable != nil {
		newLink.Reusable = link.Reusable
	}
	if link.Description != "" {
		newLink.Description = &link.Description
	}
	if link.Password != "" {
		newLink.Password = &link.Password
	}
	// if link.Domain != "" {
	// 	url := link.Domain + "/" + *newLink.Link
	// 	newLink.Address = &url
	// }
	if link.ExpireIn != "" {
		expireAt := utils.GetExpireAt(link.ExpireIn)
		if !expireAt.IsZero() {
			newLink.ExpireAt = &expireAt
		}
	}
	ip := c.ClientIP()
	newLink.IP = &ip

	// create new link
	if err = l.link.Create(c, &newLink); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "Internal Error",
				Detail: "An internal error has occurred",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, newLink)
}

// GetLink returns a link
// @Summary Get a link
// @Description Get a link
// @ID get-a-link
// @Tags links
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Link
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/links/{id} [get]
func (l *link) GetLink(c *gin.Context) {
	id := c.Param("id")
	link, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	// check if link is exists
	if link.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	c.JSON(http.StatusOK, link)
}

// UpdateLink updates a link
// @Summary Update a link
// @Description Update a link
// @ID update-link
// @Tags links
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param link body models.LinkBody true "Link"
// @Success 200 {object} models.Link
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/links/{id} [put]
func (l *link) UpdateLink(c *gin.Context) {
	id := c.Param("id")
	var link models.LinkBody
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: err.Error(),
				Status: http.StatusBadRequest,
			},
		})
		return
	}

	// retrieve link
	existingLink, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if existingLink.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	// Update link fields as per data recieved
	if link.Target != "" {
		existingLink.Target = &link.Target
	}
	if link.CustomURL != "" {
		existingLink.Link = &link.CustomURL
		url := c.Request.Host + "/" + link.CustomURL
		existingLink.Address = &url
	}
	if link.Description != "" {
		existingLink.Description = &link.Description
	}
	if link.Password != "" {
		existingLink.Password = &link.Password
	}
	// if link.Domain != "" {
	// 	url := link.Domain + "/" + *existingLink.Link
	// 	existingLink.Address = &url
	// }
	if link.ExpireIn != "" {
		expireAt := utils.GetExpireAt(link.ExpireIn)
		if !expireAt.IsZero() {
			existingLink.ExpireAt = &expireAt
		}
	}
	if link.Reusable != nil {
		existingLink.Reusable = link.Reusable
	}
	if link.ExpireIn != "" {
		expireAt := utils.GetExpireAt(link.ExpireIn)
		if !expireAt.IsZero() {
			existingLink.ExpireAt = &expireAt
		}
	}

	// update link
	if err := l.link.Update(c, &existingLink); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, existingLink)
}

// DeleteLink deletes a link
// @Summary Delete a link
// @Description Delete a link
// @ID delete-link
// @Tags links
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Link
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/links/{id} [delete]
func (l *link) DeleteLink(c *gin.Context) {
	id := c.Param("id")
	link, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if link.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	if err := l.link.Delete(c, &link); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, link)
}

// GetLinkStats returns stats of a link
// @Summary Get stats of a link
// @Description Get stats of a link
// @ID get-link-stats
// @Tags links
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Stat
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/links/{id}/stats [get]
func (l *link) GetLinkStats(c *gin.Context) {
	id := c.Param("id")
	stats, err := l.link.GetStats(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// GenQRCode generates a qr code for a link
// @Summary Generate a qr code for a link
// @Description Generate a qr code for a link
// @ID generate-qr-code
// @Tags links
// @Produce json
// @Produce data:image/png
// @Param id path string true "ID"
// @Success 200 {object} models.QRCode.Image
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/v1/links/{id}/qrcode [get]
func (l *link) GenQRCode(c *gin.Context) {
	id := c.Param("id")
	link, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if link.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	qrCode, err := l.link.GenQRCode(c, link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.Data(http.StatusOK, "image/png", qrCode.Image)
}

// NewLink returns a new link controller
func NewLink(linkRepo repository.Link) Link {
	return &link{
		link: linkRepo,
	}
}
