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
	// / POST
	Post(c *gin.Context)
	// / PUT
	Put(c *gin.Context)
	// /:id DELETE
	Delete(c *gin.Context)
	// /:id GET
	Get(c *gin.Context)
	// / GET
	List(c *gin.Context)

	// Sell(c *gin.Context)
}

