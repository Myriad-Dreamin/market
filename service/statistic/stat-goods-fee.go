package statisticservice

import (
	"github.com/gin-gonic/gin"
	"time"
)

type StatFeeRequest struct {
	LEThan time.Time `json:"le" form:"le" binding:"required"`
	GEThan time.Time `json:"ge" form:"ge" binding:"required"`
}




func (svc *Service) StatGoodsFee(c *gin.Context) {

}
