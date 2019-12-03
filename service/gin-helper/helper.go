package ginhelper

import (
	"fmt"
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/market/lib/errorc"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"reflect"
	"strconv"
)

var ResponseOK = types.Response{Code: types.CodeOK}

func CheckInsertError(c *gin.Context, err error) bool {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			c.AbortWithStatusJSON(http.StatusOK, &types.Response{Code: types.CodeDuplicatePrimaryKey})
			return true
		}
	}
	return false
}

func MissID(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
		Code:  types.CodeInvalidParameters,
		Error: "id missing in the path",
	})
}

func AuthFailed(c *gin.Context, errorString string) {
	c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
		Code:  types.CodeAuthenticatePolicyError,
		Error: errorString,
	})
}

func ParseUint(c *gin.Context, key string) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative id",
		})
		return 0, false
	}
	return uint(id), true
}

func BindRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func ParseUintAndBind(c *gin.Context, key string, req interface{}) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative id",
		})
		return 0, false
	}
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return 0, false
	}
	return uint(id), true
}

func RosolvePageVariable(c *gin.Context) (int, int, bool) {
	spage, ok := c.GetQuery("page")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "missing page number",
		})
		return 0, 0, false
	}
	page, err := strconv.Atoi(spage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeUnserializeDataError,
			Error: "can not convert page number to integer",
		})
		return 0, 0, false
	}
	spageSize, ok := c.GetQuery("page_size")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "missing page size",
		})
		return 0, 0, false
	}
	pageSize, err := strconv.Atoi(spageSize)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeUnserializeDataError,
			Error: "can not convert page size to integer",
		})
		return 0, 0, false
	}
	if page <= 0 || pageSize <= 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative params",
		})
		return 0, 0, false
	}
	return page, pageSize, true
}

func MaybeGetRawDataError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeGetRawDataError,
			Error: err.Error(),
		})
		return true
	}
	return false
}

func MaybeCountError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}

	return false
}

type applyContext gin.Context

func (ctx *applyContext) applyError(code errorc.Code, errs string) bool {
	if code != types.CodeOK {
		((*gin.Context)(ctx)).AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: errs,
		})
		return true
	}
	return false
}

func MaybeSelectError(c *gin.Context, anyObj interface{}, err error) bool {
	return (*applyContext)(c).applyError(errorc.MaybeSelectError(anyObj, err))
}

func MaybeSelectErrorWithTip(c *gin.Context, anyObj interface{}, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if reflect.ValueOf(anyObj).IsNil() {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeNotFound,
			Error: missError,
		})
		return true
	}

	return false
}

func MaybeMissingError(c *gin.Context, has bool, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &types.Response{
			Code: types.CodeNotFound,
		})
		return true
	}

	return false
}

func MaybeMissingErrorWithTip(c *gin.Context, has bool, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeNotFound,
			Error: missError,
		})
		return true
	}

	return false
}
func MaybeOnlySelectError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}

	return false
}

type Deletable interface {
	Delete() (int64, error)
}

func DeleteObj(c *gin.Context, deleteObj Deletable) bool {
	affected, err := deleteObj.Delete()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.Response{
			Code: types.CodeDeleteNoEffect,
		})
		return false
	}
	return true
}

type Creatable interface {
	Create() (int64, error)
}

func CreateObj(c *gin.Context, createObj Creatable) bool {
	affected, err := createObj.Create()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.Response{
			Code: types.CodeInsertError,
		})
		return false
	}
	return true
}

func CreateObjWithTip(c *gin.Context, createObj Creatable) bool {
	affected, err := createObj.Create()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: fmt.Sprintf("create %T failed: %v", createObj, err.Error()),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.Response{
			Code: types.CodeInsertError,
		})
		return false
	}
	return true
}

type Updatable interface {
	Update() (int64, error)
}

func UpdateObj(c *gin.Context, updateObj Updatable) bool {
	affected, err := updateObj.Update()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &types.Response{
			Code: types.CodeUpdateError,
		})
		return false
	}
	return true
}

func UpdateFields(c *gin.Context, obj errorc.UpdateFieldsable, fields []string) bool {
	return !(*applyContext)(c).applyError(errorc.UpdateFields(obj, fields))
}

func GetCustomFields(c *gin.Context) *types.CustomFields {
	claims, _ := c.Get("claims")
	return claims.(*jwt.CustomClaims).CustomField.(*types.CustomFields)
}
