package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"

	"github.com/tech-thinker/linkly/repository"
)

type URL interface {
	Add(ctx *gin.Context)
	Get(ctx *gin.Context)
	GenQR(ctx *gin.Context)
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
	var url models.URL
	err = json.Unmarshal(bytes, &url)
	if err != nil {
		log.Fatal(err)
	}
	err = u.urlRepo.Add(ctx, &url)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": err.Error(),
			"urls":    url,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
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
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"urls":    url,
	})
}

// GenQR generates qr code and returns the png
func (u *url) GenQR(ctx *gin.Context) {
	var qr models.QRCode
	qr.Content = ctx.Param("short_url")
	qr.Content = "http://" + "cut.mrinjamul.in" + "/" + qr.Content
	qr, err := u.urlRepo.GenQR(ctx, qr)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Data(200, "image/png", qr.Image)
}

// GetAndRedirect redirects to the original url
func (u *url) GetAndRedirect(ctx *gin.Context) {
	var url models.URL
	url.ShortURL = ctx.Param("short_url")

	url, err := u.urlRepo.GetAndRedirect(ctx, url)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// ctx.JSON(200, gin.H{
	// 	"message": "Success",
	// 	"urls":    url,
	// })
	redirectURL := "http://" + url.URL
	// Use http.StatusFound (302) to tell the http client to redirect
	ctx.Redirect(http.StatusFound, redirectURL)
}

// GetAll : get all urls
func (u *url) GetAll(ctx *gin.Context) {
	urls, err := u.urlRepo.GetAll(ctx)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
			"urls":    urls,
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
	// Read the body of the request
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Unmarshal the json into a url struct
	var url models.URL
	err = json.Unmarshal(bytes, &url)
	if err != nil {
		log.Fatal(err)
	}
	// Update the url in the database
	err = u.urlRepo.Update(ctx, &url)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
			"urls":    url,
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
	var url models.URL
	err := u.urlRepo.Delete(ctx, &url)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
			"urls":    url,
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
