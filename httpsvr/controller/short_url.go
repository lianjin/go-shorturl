package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gsurl/httpsvr/middleware"
	"gsurl/httpsvr/vo"
	"gsurl/log"
	"gsurl/service"
)

func GenShortUrl(c *gin.Context) {
	req := &vo.ShortUrlGenReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	shortUrl, err := service.GenShortUrl(c, req.Url)
	if err != nil {
		log.Logger.Errorf("Error generating short URL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to generate short URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"short-url": shortUrl})
}

func GetShortUrl(c *gin.Context) {
	shortCode := c.Param("short_code")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "short_code parameter is required"})
		return
	}
	originalUrl, err := service.GetShortUrl(c, shortCode)
	if err != nil {
		log.Logger.Errorf("Error retrieving original URL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve original URL"})
		return
	}
	if originalUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Short URL not found"})
		return
	}
	middleware.IncShortUrlReqCounter(shortCode)
	c.Redirect(http.StatusSeeOther, originalUrl)
}
