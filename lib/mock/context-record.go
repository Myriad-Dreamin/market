package mock

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ContextRecorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("here")
		c.Header("gin_context_matched_path_method", c.FullPath() + "@" + c.Request.Method)
	}
}


