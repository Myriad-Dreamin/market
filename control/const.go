package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

//go:generate package-attach-to-path -generate_register_map


var ConstCates []interface{}

// @Category Const - ServiceCode
// @Description Get Service Codes
// @Path /v2/const/service-codes
type serviceCodeCate interface {
}

// @Category Const - Cities
// @Description Get All valid city code and its corresponding city object
// @Path /v2/const/user-cities
type cityCate interface {
}

// @Category Const - Goods Types
// @Description Get All valid goods types
// @Path /v2/const/goods-types
type goodsTypeCate interface {
}

func init() {
	var (
		a serviceCodeCate = 0
		b cityCate = 0
		c goodsTypeCate = 0
	)
	ConstCates = []interface{}{
		&a,
		&b,
		&c,
	}
}
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
