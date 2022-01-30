package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"

	"github.com/tech-thinker/linkly/repository"
)

type URL interface {
	Add(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAndRedirect(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type url struct {
	urlRepo repository.URLRepo
}

// Add a new url
func (u *url) Add(ctx *gin.Context) {
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var url models.Input
	err = json.Unmarshal(bytes, &url)
	if err != nil {
		log.Fatal(err)
	}
	u.urlRepo.Add(ctx, &url)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"urls":    url,
	})
}

// Get a url by short url
func (u *url) Get(ctx *gin.Context) {
	var url models.URL
	url.ShortURL = ctx.Param("short_url")

	url, err := u.urlRepo.Get(ctx, url)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "URL not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"urls":    url,
	})
}

// GetAndRedirect redirects to the original url
func (u *url) GetAndRedirect(ctx *gin.Context) {
	var url models.URL
	url.ShortURL = ctx.Param("short_url")

	url, err := u.urlRepo.GetAndRedirect(ctx, url)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "URL not found",
		})
		return
	}
	redirectURL := "http://" + url.URL
	ctx.Redirect(http.StatusMovedPermanently, redirectURL)
}

// GetAll : get all urls
func (u *url) GetAll(ctx *gin.Context) {
	urls, err := u.urlRepo.GetAll(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "URL not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"urls":    urls,
	})
}

// Update a url
func (u *url) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatal(err)
	}

	var url models.URL
	err = json.Unmarshal(bytes, &url)
	if err != nil {
		log.Fatal(err)
	}
	url.ID = uint64(id)

	err = u.urlRepo.Update(ctx, &url)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "URL not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"urls":    url,
	})
}

// Delete a url
func (u *url) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	url := models.URL{
		ID: uint64(id),
	}
	err = u.urlRepo.Delete(ctx, &url)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "URL not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
		"urls":    url,
	})
}

// NewURL returns a new url controller
func NewURL(urlRepo repository.URLRepo) URL {
	return &url{
		urlRepo: urlRepo,
	}
}
