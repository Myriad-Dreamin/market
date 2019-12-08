package statisticservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)


type StatFeeXYReply struct {
	Code int `json:"code"`
	Results []model.StatFeeXYResult `json:"results"`
}

func (svc *Service) StatGoodsFeeXY(c *gin.Context) {
	svc.statGoodsFeeXYService.List(c)
}

func ToStatFeeXYReply(results []model.StatFeeXYResult) StatFeeXYReply {
	return StatFeeXYReply{Code:types.CodeOK, Results:results}
}

type statGoodsFeeXYService struct {
}

func (svc statGoodsFeeXYService) CreateFilter() interface{} {
	return new(model.StatFeeRequest)
}

func (svc statGoodsFeeXYService) ProcessListResults(_ *gin.Context, r interface{}) interface{} {
	return ToStatFeeXYReply(r.([]model.StatFeeXYResult))
}


