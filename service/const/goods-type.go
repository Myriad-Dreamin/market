package constservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

type GoodsTypesReply struct {
	Types map[string]types.GoodsType `json:"types"`
}

var goodsTypesReply = GoodsTypesReply{types.GoodsTypesMap}

func (svc *Service) GetGoodsTypes(c controller.MContext) {
	c.JSON(http.StatusOK, goodsTypesReply)
}
