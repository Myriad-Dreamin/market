package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

func (svc *Service) PutPicture(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, svc.key)
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

	if err = c.SaveUploadedFile(file, svc.cfg.BaseParametersConfig.GoodsPicturePath+strconv.Itoa(int(id))+filepath.Ext(file.Filename)); err != nil {
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
