package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns a healthcheck response
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

// Home returns a home response
func Home() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	}
}

func GetLinks() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sorry, not implemented yet",
		})
	}
}

func CreateLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sorry, not implemented yet",
		})
	}
}

func GetLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sorry, not implemented yet",
		})
	}
}

func UpdateLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sorry, not implemented yet",
		})
	}
}

func DeleteLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sorry, not implemented yet",
		})
	}
}

func GetLinkStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sorry, not implemented yet",
		})
	}
}
