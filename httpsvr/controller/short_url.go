package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gsurl/httpsvr/vo"
	"gsurl/service"
)

func GenShortUrl(c *gin.Context) {
	req := &vo.ShortUrlGenReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	shortUrl := service.GenShortUrl(req.Url)
	c.JSON(http.StatusOK, gin.H{"Data": shortUrl})
}
