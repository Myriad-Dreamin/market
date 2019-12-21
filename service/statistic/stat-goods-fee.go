package statisticservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type StatFeeReply struct {
	Code types.CodeType `json:"code"`
}

func ToStatFeeReply() StatFeeReply {
	return StatFeeReply{Code: types.CodeOK}
}

func (svc *Service) StatGoodsFee(c controller.MContext) {
	svc.statGoodsFeeService.List(c)
	return
}

type statGoodsFeeService struct {
}

func (svc statGoodsFeeService) CreateFilter() interface{} {
	return new(model.StatFeeRequest)
}

func (svc statGoodsFeeService) ProcessListResults(c controller.MContext, r interface{}) interface{} {
	return ToStatFeeReply()
}
