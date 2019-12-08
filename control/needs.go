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

	// @Title Post
	// @Description Post Needs
	Post(c *gin.Context)

	// @Title Put
	// @Description Put Needs
	Put(c *gin.Context)

	// @Title Delete
	// @Description Delete Needs
	Delete(c *gin.Context)

	// @Title Force Delete
	// @Description Force Delete Needs
	ForceDelete(c *gin.Context)

	// @Title Get
	// @Description Get Needs
	Get(c *gin.Context)

	// @Title List
	// @Description List Needs
	List(c *gin.Context)


	// @Title Post Picture
	// @Description Post Picture of Goods
	PostPicture(c *gin.Context)
}

