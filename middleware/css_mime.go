package middleware

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func MimeType(c *gin.Context) {
	path := c.Request.URL.Path
	ext := filepath.Ext(path)

	if ext == ".css" {
		c.Header("Content-Type", "text/css")
	}

	c.Next()
}
