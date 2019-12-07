package control

import "github.com/gin-gonic/gin"


/* goods
 * post: 登记商品
 * put: 修改商品信息
 * delete: 删除商品信息
 * force delete: 强制删除商品信息（只有管理员能调用）
 * get: 获取商品信息
 * list: 根据filter得到商品列表
 */
type GoodsService interface {

	// @Title Post
	// @Description Post Goods
	Post(c *gin.Context)

	// @Title Put
	// @Description Put Goods
	Put(c *gin.Context)

	// @Title Delete
	// @Description Delete Goods
	Delete(c *gin.Context)

	// @Title Force Delete
	// @Description Force Delete Goods
	ForceDelete(c *gin.Context)

	// @Title Get
	// @Description Get Goods
	Get(c *gin.Context)

	// @Title List
	// @Description List Goods
	List(c *gin.Context)

	// Buy(c *gin.Context)
}

