package statisticservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type StatFeeCountXYReply struct {
	Code    int                          `json:"code"`
	Results []model.StatFeeCountXYResult `json:"results"`
}

func (svc *Service) StatGoodsCountXY(c *gin.Context) {
	svc.statGoodsFeeCountXYService.List(c)
}

func ToStatFeeCountXYReply(results []model.StatFeeCountXYResult) StatFeeCountXYReply {
	return StatFeeCountXYReply{Code: types.CodeOK, Results: results}
}

type statGoodsFeeCountXYService struct {
}

func (svc statGoodsFeeCountXYService) CreateFilter() interface{} {
	return new(dblayer.StatFeeRequest)
}

func (svc statGoodsFeeCountXYService) ProcessListResults(c *gin.Context, r interface{}) interface{} {
	return ToStatFeeCountXYReply(r.([]model.StatFeeCountXYResult))
}
