package httpsvr

import (
	"gsurl/httpsvr/controller"
	"gsurl/log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

func (z *zapLogger) Write(p []byte) (n int, err error) {
	z.logger.Info(strings.TrimSpace(string(p)))
	return len(p), nil
}

func Init() {
	r := gin.Default()
	gin.DefaultWriter = &zapLogger{logger: zap.L().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false))}
	gin.DefaultErrorWriter = &zapLogger{logger: zap.L().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false))}
	r.POST("/short-url", controller.GenShortUrl)
	r.GET("/:short_code", controller.GetShortUrl)
	log.Logger.Infof("Starting server on port 8080")
	r.Run(":8080")
}
