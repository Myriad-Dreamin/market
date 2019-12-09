package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

/* goods
 * post: 登记商品
 * put: 修改商品信息
 * delete: 删除商品信息
 * force delete: 强制删除商品信息（只有管理员能调用）
 * get: 获取商品信息
 * list: 根据filter得到商品列表
 */
type GoodsService interface {
	GoodsSignatureXXX() interface{}

	// @Title Goods Post
	// @Description Post a Goods you want to sell to others.
	// The following is a description of the parameters
	//     + `g_type uint16`: the type of goods, optional values refer to the link [market types](https://github.com/Myriad-Dreamin/market/blob/master/types/goods_type.go):
	//
	//     + `name string`: the name of your goods.
	//
	//     + `min_price uint64`: the lowest price required to sell the goods, if you want to sell the goods at least $10.55 then `min_price` will equal to 1055 to avoid the loss of floating precision (10.55 * 100)
	//
	//     + `is_fixed bool`: If `is_fixed` is false, then the item will be auctioned. And after the end time, you can confirm and sell at the highest price. Otherwise, then the item will be sold on a first-come-first-get basis. In this case, you should also confirm the transaction.
	//
	//     + `description string`: a short description of the goods.
	Post(c controller.MContext)

	// @Title Goods Put
	// @Description Put Goods
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
}

