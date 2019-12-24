package needsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var picField = []string{"pic_name"}

func (svc *Service) PutPicture(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, svc.key)
	if !ok {
		return
	}

	needs, err := svc.needsDB.ID(id)
	if ginhelper.MaybeSelectError(c, needs, err) {
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

	if len(needs.PicName) != 0 {
		if err = os.Remove(needs.PicName); err != nil {
			c.JSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeFSExecError,
				Error: err.Error(),
			})
			return
		}
	}

	needs.PicName = strconv.Itoa(int(id)) + filepath.Ext(file.Filename)
	if !ginhelper.UpdateFields(c, needs, picField) {
		return
	}

	if err = c.SaveUploadedFile(file, svc.cfg.BaseParametersConfig.GoodsPicturePath+needs.PicName); err != nil {
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
