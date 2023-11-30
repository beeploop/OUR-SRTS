package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogrusMiddleware(c *gin.Context) {
	// Starting time request
	startTime := time.Now()

	// Processing request
	c.Next()

	// End Time request
	endTime := time.Now()

	// execution time
	latencyTime := endTime.Sub(startTime)

	// Request method
	reqMethod := c.Request.Method

	// Request route
	reqUri := c.Request.RequestURI

	// status code
	statusCode := c.Writer.Status()

	// Request IP
	clientIP := c.ClientIP()

	logrus.WithFields(logrus.Fields{
		"METHOD":    reqMethod,
		"URI":       reqUri,
		"STATUS":    statusCode,
		"LATENCY":   latencyTime,
		"CLIENT_IP": clientIP,
	}).Info("HTTP REQUEST")

	c.Next()
}
