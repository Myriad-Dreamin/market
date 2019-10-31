package mock

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ContextRecorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("here")
		c.Header("Gin-Context-Matched-Path-Method", c.FullPath() + "@" + c.Request.Method)
	}
}


