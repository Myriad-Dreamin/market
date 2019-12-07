package control

import "github.com/gin-gonic/gin"

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
	UserSignatureXXX() interface{}

	// @Title Login
	// @Description Login
	Login(c *gin.Context)

	// @Title Register
	// @Description Register
	Register(c *gin.Context)

	// @Title Put
	// @Description Put User
	Put(c *gin.Context)

	// @Title ChangePassword
	// @Description change password of user
	ChangePassword(c *gin.Context)

	//// /:id/grant PUT
	//Grant(c *gin.Context)

	// @Title Get
	// @Description Get User
	Get(c *gin.Context)

	// @Title Delete
	// @Description Delete User
	Delete(c *gin.Context)

	// @Title List
	// @Description List User
	List(c *gin.Context)

	// @Title Buy
	// @Description Buy Goods
	Buy(c *gin.Context)

	// @Title Sell
	// @Description Sell Needs
	Sell(c *gin.Context)

	// @Title ConfirmBuy
	// @Description ConfirmBuy Goods
	ConfirmBuy(c *gin.Context)

	// @Title ConfirmSell
	// @Description ConfirmSell Needs
	ConfirmSell(c *gin.Context)
}