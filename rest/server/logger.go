package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/tools/log"
	uuid "github.com/satori/go.uuid"
)

func newGinLogger(c *gin.Context, deps ...interface{}) log.LogRusEntry {
	return log.Get(deps...).
		WithField(log.LOG_FIELD_CORRELATION_ID, getCorrelationId(c)).
		WithField(log.LOG_FIELD_CONTROLLER, "Rest").
		WithField(log.LOG_FIELD_HTTP_METHOD, c.Request.Method).
		WithField(log.LOG_FIELD_HTTP_PATH, c.Request.URL.Path)
}

func GinLoggerMiddleware(deps ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := newGinLogger(c, deps...)
		c.Set("logger", logger)

		c.Next()

		if c.Request.Method != "OPTIONS" {
			deps := GinDeps(c)
			log.Get(deps...).WithField(log.LOG_FIELD_HTTP_STATUS, c.Writer.Status()).Info("Completed")
		}
	}
}

func ginLogger(c *gin.Context) log.LogRusEntry {
	logger, exist := c.Get("logger")
	if !exist {
		return newGinLogger(c)
	}
	return logger.(log.LogRusEntry)
}

func getCorrelationId(c *gin.Context) string {
	value := c.GetHeader(log.LOG_FIELD_CORRELATION_ID)

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
