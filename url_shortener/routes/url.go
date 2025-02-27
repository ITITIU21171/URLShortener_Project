package routes

import (
	"net/http"

	"github.com/ITITIU21171/url-shortener/services"
	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var request struct {
		OriginalURL string `json:"original_url"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	shortURL := services.GenerateShortURL(request.OriginalURL)
	err := services.SaveURL(request.OriginalURL, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "fail to connect to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func RedirectURL(c *gin.Context) {
	shortURl := c.Param("shortURL")

	urlData, err := services.GetOriginalURL(shortURl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
	}

	c.Redirect(http.StatusMovedPermanently, urlData.OriginalURL)
}
