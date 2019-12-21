package constservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

type ServiceCodeReply struct {
	Codes map[types.CodeType]string `json:"codes"`
}

var serviceCodeReply = ServiceCodeReply{types.CodeDesc}

func (svc *Service) GetServiceCode(c controller.MContext) {
	c.JSON(http.StatusOK, serviceCodeReply)
}
