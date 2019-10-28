package needsservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	goods_service "github.com/Myriad-Dreamin/market/service/goods-service"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.NeedsDB
	logger types.Logger
	filterFunc func(c *gin.Context) interface{}
}


func NewService(logger types.Logger, provider *model.Provider, _ *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.db = provider.NeedsDB()
	a.logger = logger
	a.CRUDService = base_service.NewCRUDService(a, "nid")
	a.ListService = base_service.NewListService(a, "nid")
	a.filterFunc = goods_service.ListFilter(a.db)
	return
}

/*
type Needs struct {
}
*/
