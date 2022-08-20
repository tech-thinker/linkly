package controllers

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Views interface {
	Index(ctx *gin.Context, viewsFs embed.FS)
	ShortURL(ctx *gin.Context)
	ShortURLQR(ctx *gin.Context)
}

type views struct {
}

func (v *views) Index(ctx *gin.Context, viewsFs embed.FS) {
	fsRoot, err := fs.Sub(viewsFs, "views")
	if err != nil {
		log.Println(err)
	}
	gin.WrapH(http.FileServer(http.FS(fsRoot)))
}

// ShortURL redirects to the original url
func (v *views) ShortURL(ctx *gin.Context) {

}

// ShortURLQR generates a QR code for the short url
func (v *views) ShortURLQR(ctx *gin.Context) {

}

// NewViews returns a new views controller
func NewViews() Views {
	return &views{}
}
