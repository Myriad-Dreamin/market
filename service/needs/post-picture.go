package needsservice

import (
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (srv *Service) PostPicture(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, srv.key)
	if !ok {
		return
	}

	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}

	if err = c.SaveUploadedFile(file, srv.cfg.BaseParametersConfig.NeedsPicturePath + strconv.Itoa(int(id))); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
}

