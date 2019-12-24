package goodsservice

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

	goods, err := svc.goodsDB.ID(id)
	if ginhelper.MaybeSelectError(c, goods, err) {
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

	if len(goods.PicName) != 0 {
		if err = os.Remove(goods.PicName); err != nil {
			c.JSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeFSExecError,
				Error: err.Error(),
			})
			return
		}
	}

	goods.PicName = strconv.Itoa(int(id)) + filepath.Ext(file.Filename)
	if !ginhelper.UpdateFields(c, goods, picField) {
		return
	}

	if err = c.SaveUploadedFile(file, svc.cfg.BaseParametersConfig.GoodsPicturePath+goods.PicName); err != nil {
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
