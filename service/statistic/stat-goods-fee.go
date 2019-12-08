package statisticservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type StatFeeReply struct {
	Code int `json:"code"`
}

func ToStatFeeReply() StatFeeReply {
	return StatFeeReply{Code:types.CodeOK}
}


func (svc *Service) StatGoodsFee(c *gin.Context) {
	svc.statGoodsFeeService.List(c)
	return
}

type statGoodsFeeService struct {
}

func (svc statGoodsFeeService) CreateFilter() interface{} {
	return new(model.StatFeeRequest)
}

func (svc statGoodsFeeService) ProcessListResults(c *gin.Context, r interface{}) interface{} {
	return ToStatFeeReply()
}
