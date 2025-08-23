package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// ① 定义监控指标
var (
	// 请求总数
	reqTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	// 请求耗时直方图
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latencies",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path"},
	)
	// 短链接请求总数
	shortUrlReqCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "short_url_req_total",
			Help: "业务自定义计数",
		},
		[]string{"code"},
	)
)

// ② 注册指标
func init() {
	prometheus.MustRegister(reqTotal)
	prometheus.MustRegister(reqDuration)
	prometheus.MustRegister(shortUrlReqCounter)
}

// ③ Gin 中间件：记录指标
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// 收集标签
		method := c.Request.Method
		path := c.FullPath() // 路由模板，如 /user/:id
		status := strconv.Itoa(c.Writer.Status())

		// 上报
		reqTotal.WithLabelValues(method, path, status).Inc()
		reqDuration.WithLabelValues(method, path).Observe(float64(time.Since(start).Milliseconds()))
	}
}

func IncShortUrlReqCounter(code string) {
	if code == "" {
		return
	}
	shortUrlReqCounter.WithLabelValues(code).Inc()
}
