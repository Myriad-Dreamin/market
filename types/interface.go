package types

import "github.com/gin-gonic/gin"

/* auth
 * refresh token GET: 刷新登陆用token
 */
type Auth interface {
	// /refresh-token GET
	RefreshToken(c *gin.Context)
}

/* user
 * login POST 登陆
 * register POST 注册
 * put PUT 更新用户信息
 * change password PUT 修改用户密码
 * grant PUT 授予用户身份或权限
 * get GET 获取用户信息
 * delete DELETE 删除用户
 */
type UserService interface {
	// /:id POST
	Login(c *gin.Context)
	// / POST
	Register(c *gin.Context)
	// /:id PUT
	Put(c *gin.Context)
	// /:id/password PUT
	ChangePassword(c *gin.Context)
	//// /:id/grant PUT
	//Grant(c *gin.Context)
	// /:id GET
	Get(c *gin.Context)
	// /:id DELETE
	Delete(c *gin.Context)

	// user-list GET
	List(c *gin.Context)
}

/* goods
 * post: 登记商品
 * put: 修改商品信息
 * delete: 删除商品信息
 * get: 获取商品信息
 * list: 根据filter得到商品列表
 */
type GoodsService interface {
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

	// Buy(c *gin.Context)
}

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


/*
查询一定条件下当前已经成交物品的累计中介费收益信息
统计分析功能模块：显示起始年月、终止年月、某个地域不同物品类型成交中介费的明细，并按月以折线图的方式展示每月累计成交笔数、中介费金额的变化趋势
*/

/* statistic
 * show GET: 统计信息
 */
type StatisticService interface {
	// /show GET
	FilterGoods(c *gin.Context)
}
