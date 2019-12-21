package constservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

type CitiesReply struct {
	Cities map[string]types.CityObject `json:"cities"`
}

func (svc *Service) GetCities(c controller.MContext) {
	c.JSON(http.StatusOK, svc.citiesReply)
}
