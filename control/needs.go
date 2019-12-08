package control

import "github.com/gin-gonic/gin"


/* needs
 * post: 登记商品
 * put: 修改商品信息
 * delete: 删除商品信息
 * get: 获取商品信息
 * list: 根据filter得到商品列表
 */
type NeedsService interface {
	NeedsSignatureXXX() interface{}

	// @Title Needs Post
	// @Description Post a Needs you want to buy from others.
	// The following is a description of the parameters
	//     + `g_type uint16`: the type of needs, optional values refer to the link [market types](https://github.com/Myriad-Dreamin/market/blob/master/types/goods_type.go):
	//
	//     + `name string`: the name of your needs.
	//
	//     + `min_price uint64`: the lowest price required to buy the needs, if you want to buy the needs at least $10.55 then `min_price` will equal to 1055 to avoid the loss of floating precision (10.55 * 100)
	//
	//     + `max_price uint64`: the highest price accepted to buy the needs, if you want to buy the needs at least $10.55 then `min_price` will equal to 1055 to avoid the loss of floating precision (10.55 * 100)
	//
	//     + `description string`: a short description of the needs.
	Post(c *gin.Context)

	// @Title Needs Put
	// @Description Put Needs
	Put(c *gin.Context)

	// @Title Needs Delete
	// @Description This is the public interface for deleting needs. You cannot use this interface to delete a transaction that has already been completed (`needs.Status == types.GoodsStatusFinished`). the request url is `v1/needs/:nid.DELETE` where `:nid` is the corresponding needs number
	Delete(c *gin.Context)

	// @Title Needs Force Delete
	// @Description This is the private interface for deleting an needs. Only administrators can call this api. It can delete any needs, so when you call it, you should know what you are doing. the request url is `v1/needs/:nid/force.DELETE` where `:nid` is the corresponding needs number
	ForceDelete(c *gin.Context)

	// @Title Needs Get
	// @Description This is a public interface for getting needs information, and only returns limited information. If you want to get all information of specified needs, you can access the Inspect interface. the request url is `v1/needs/:nid.GET` where `:nid` is the corresponding needs number
	Get(c *gin.Context)

	// @Title Needs Inspect
	// @Description This is a public interface for getting needs information, and will returns all the information. the request url is `v1/needs/:nid/inspect.GET` where `:nid` is the corresponding needs number
	Inspect(c *gin.Context)

	// @Title Needs List
	// @Description List Needs
	List(c *gin.Context)


	// @Title Upload Needs' Picture
	// @Description This interface is used to upload pictures of needs. It is not merged into the Put interface just because the man who writes backend is too lazy. the request url is `v1/needs/:nid/picture.PUT` where `:nid` is the corresponding needs number
	PutPicture(c *gin.Context)
}

