package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

//go:generate package-attach-to-path -generate_register_map

type ConstService interface {
	ConstSignatureXXX() interface{}

	// @Title ServiceCode
	// @Description ServiceCode Const
	GetServiceCode(c controller.MContext)

	// @Title GetGoodsTypes
	// @Description GetGoodsTypes
	GetGoodsTypes(c controller.MContext)

	// @Title GetCities
	// @Description GetCities Needs
	GetCities(c controller.MContext)
}
