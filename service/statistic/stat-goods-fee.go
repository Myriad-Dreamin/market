package statisticservice

import (
	"github.com/gin-gonic/gin"
	"time"
)

type StatFeeRequest struct {
	LEThan time.Time `json:"le" form:"le" binding:"required"`
	GEThan time.Time `json:"ge" form:"ge" binding:"required"`
}


//func (srv *Service) FilterOn(c *gin.Context) (interface{}, error) {
//
//	result := srv.filterFunc(c)
//	if c.IsAborted() {
//		return nil, nil
//	}
//
//	return srv.GoodssToListReply(c, result.([]model.Goods)), nil
//}
//
func (svc *Service) StatGoodsFee(c *gin.Context) {

}
