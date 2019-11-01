package mock

import (
	"github.com/gin-gonic/gin"
)

func ContextRecorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Header("Gin-Context-Matched-Path-Method", c.FullPath() + "@" + c.Request.Method)
	}
}


