package statisticservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type StatFeeCountXYReply struct {
	Code    types.CodeType               `json:"code"`
	Results []model.StatFeeCountXYResult `json:"results"`
}

func (svc *Service) StatGoodsCountXY(c controller.MContext) {
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

func (svc statGoodsFeeCountXYService) ProcessListResults(c controller.MContext, r interface{}) interface{} {
	return ToStatFeeCountXYReply(r.([]model.StatFeeCountXYResult))
}
