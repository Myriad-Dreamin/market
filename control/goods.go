package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

var GoodsCates []interface{}

// @Category Goods - Post
// @Path /v1/goods
type goodsPostCate interface {
}

// @Category Goods - List
// @Path /v1/goods-list
type goodsListCate interface {
}

// @Category Goods - Put/Delete/Get Api Group
// @Path /v1/goods/:goid
type goodsIdGroupCate interface {
}

// @Category Goods - Force Delete
// @Path /v1/goods/:goid/force
type goodsForceDeleteCate interface {
}

// @Category Goods - Inspect
// @Path /v1/goods/:goid/inspect
type goodsInspectCate interface {
}

// @Category Goods - Picture Api Group
// @Path /v1/goods/:goid/picture
type goodsIdPictureCate interface {
}

func init() {
	var (
		a goodsPostCate        = 0
		b goodsListCate        = 0
		c goodsIdGroupCate     = 0
		d goodsForceDeleteCate = 0
		e goodsInspectCate     = 0
		f goodsIdPictureCate   = 0
	)
	GoodsCates = []interface{}{
		&a,
		&b,
		&c,
		&d,
		&e,
		&f,
	}
}

// GoodsService defines the interface of goods service
type GoodsService interface {
	GoodsSignatureXXX() interface{}

	// @Title Goods Post
	// @Description Post a Goods you want to sell to others.
	// The following is a description of the parameters
	//     + `g_type uint16`: the type of goods, optional values refer to the link [market types](https://github.com/Myriad-Dreamin/market/blob/master/types/goods_type.go).
	//
	//     + `name string`: the name of your goods.
	//
	//     + `end_at time.Time`: the alive time of goods.
	//
	//     + `min_price uint64`: the lowest price required to sell the goods, if you want to sell the goods at least $10.55 then `min_price` will equal to 1055 to avoid the loss of floating precision (10.55 * 100)
	//
	//     + `is_fixed bool`: If `is_fixed` is false, then the item will be auctioned. And after the end time, you can confirm and sell at the highest price. Otherwise, then the item will be sold on a first-come-first-get basis. In this case, you should also confirm the transaction.
	//
	//     + `description string`: a short description of the goods.
	Post(c controller.MContext)

	// @Title Goods Put
	// @Description Put a Goods that not be finished. you can not reset its price to be fixed or not. you also can not modify it before a minute of ddl
	// The following is a description of the parameters
	//     + `g_type uint16`: the type of goods, optional values refer to the link [market types](https://github.com/Myriad-Dreamin/market/blob/master/types/goods_type.go).
	//
	//     + `name string`: the name of your goods.
	//
	//     + `end_at time.Time`: the alive time of goods.
	//
	//     + `min_price uint64`: the lowest price required to sell the goods, if you want to sell the goods at least $10.55 then `min_price` will equal to 1055 to avoid the loss of floating precision (10.55 * 100)
	//
	//     + `description string`: a short description of the goods.
	Put(c controller.MContext)

	// @Title Goods Delete
	// @Description This is the public interface for deleting goods. You cannot use this interface to delete a transaction that has already been completed (`goods.Status == types.GoodsStatusFinished`). the request url is `v1/goods/:goid.DELETE` where `:goid` is the corresponding goods number
	Delete(c controller.MContext)

	// @Title Goods Force Delete
	// @Description This is the private interface for deleting an goods. Only administrators can call this api. It can delete any goods, so when you call it, you should know what you are doing. the request url is `v1/goods/:goid/force.DELETE` where `:goid` is the corresponding goods number
	ForceDelete(c controller.MContext)

	// @Title Goods Get
	// @Description This is a public interface for getting goods information, and only returns limited information. If you want to get all information of specified goods, you can access the Inspect interface. the request url is `v1/goods/:goid.GET` where `:goid` is the corresponding goods number
	Get(c controller.MContext)

	// @Title Goods Inspect
	// @Description This is a public interface for getting goods information, and will returns all the information. the request url is `v1/goods/:goid/inspect.GET` where `:goid` is the corresponding goods number
	Inspect(c controller.MContext)

	// @Title Goods List
	// @Description List Goods
	List(c controller.MContext)

	// @Title Upload Goods' Picture
	// @Description This interface is used to upload pictures of goods. It is not merged into the Put interface just because the man who writes backend is too lazy. the request url is `v1/goods/:goid/picture.PUT` where `:goid` is the corresponding goods number
	PutPicture(c controller.MContext)

	// @Title Download Goods' Picture
	// @Description This interface is used to download pictures of goods. the request url is `v1/goods/:goid/picture.GET` where `:goid` is the corresponding goods number
	GetPicture(c controller.MContext)

	// @Title GetGoodsTypes
	// @Description GetGoodsTypes
	GetTypes(c controller.MContext)
}
