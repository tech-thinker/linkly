package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
)

type Domain interface {
	GetDomains(ctx *gin.Context)
}

type domain struct {
}

// GetDomains returns all domains
// @Summary Get all domains
// @Description Get all domains
// @ID get-domains
// @Tags domains
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Domain
// @Failure 500 {object} models.Error
// @Router /api/v1/domains [get]
func (d *domain) GetDomains(ctx *gin.Context) {
	var domains []models.Domain
	ctx.JSON(http.StatusNotImplemented, domains)
}

// NewDomain returns a new domain controller
func NewDomain() Domain {
	return &domain{}
}
