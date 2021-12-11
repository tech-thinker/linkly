package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/repository"
)

type URL interface {
	Add(ctx *gin.Context)
	Find(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type url struct {
	userRepo repository.URLRepository
}

// Add a new URL
func (svc *url) Add(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Add"})
}

// Find a URL
func (svc *url) Find(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Find"})
}

// Update a URL
func (svc *url) Update(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update"})
}

// Delete a URL
func (svc *url) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete"})
}

// FindAll a URL
func (svc *url) FindAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "FindAll"})
}

// NewURL returns a new URL interface
func NewURL(userRepo repository.URLRepository) URL {
	return &url{userRepo}
}
