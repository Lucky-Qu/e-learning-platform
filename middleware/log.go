package middleware

import (
	"e-learning-platform/log/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		duration := end.Sub(start)
		ip := c.ClientIP()
		reqMethod := c.Request.Method
		logger.DefaultLogger.Logger.Info("请求成功",
			zap.String("请求ip：", ip),
			zap.String("请求方法：", reqMethod),
			zap.String("uri：", c.Request.RequestURI),
			zap.Int("状态：", c.Writer.Status()),
			zap.Duration("运行时间：", duration),
		)
	}
}
