package statisticservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type StatFeeXYReply struct {
	Code    types.CodeType          `json:"code"`
	Results []model.StatFeeXYResult `json:"results"`
}

func (svc *Service) StatGoodsFeeXY(c controller.MContext) {
	svc.statGoodsFeeXYService.List(c)
}

func ToStatFeeXYReply(results []model.StatFeeXYResult) StatFeeXYReply {
	return StatFeeXYReply{Code: types.CodeOK, Results: results}
}

type statGoodsFeeXYService struct {
}

func (svc statGoodsFeeXYService) CreateFilter() interface{} {
	return new(model.StatFeeRequest)
}

func (svc statGoodsFeeXYService) ProcessListResults(_ controller.MContext, r interface{}) interface{} {
	return ToStatFeeXYReply(r.([]model.StatFeeXYResult))
}
