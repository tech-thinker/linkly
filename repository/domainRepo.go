package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
)

type DomainRepo interface {
	Get(ctx *gin.Context, url models.URL) (models.URL, error)
}
