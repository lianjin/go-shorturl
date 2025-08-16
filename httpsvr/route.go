package httpsvr

import (
	"gsurl/httpsvr/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.POST("/short-url", controller.GenShortUrl)
	r.Run(":8080")
}
