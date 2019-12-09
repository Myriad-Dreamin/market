package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

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
	Login(c controller.MContext)

	// @Title Register
	// @Description Register
	Register(c controller.MContext)

	// @Title Put
	// @Description Put User
	Put(c controller.MContext)

	// @Title ChangePassword
	// @Description change password of user
	ChangePassword(c controller.MContext)

	//// /:id/grant PUT
	//Grant(c mcontext.MContext)

	// @Title Get
	// @Description Get User
	Get(c controller.MContext)

	// @Title Delete
	// @Description Delete User
	Delete(c controller.MContext)

	// @Title List
	// @Description List User
	List(c controller.MContext)

	// @Title Buy
	// @Description Buy Goods
	Buy(c controller.MContext)

	// @Title Sell
	// @Description Sell Needs
	Sell(c controller.MContext)

	// @Title ConfirmBuy
	// @Description ConfirmBuy Goods
	ConfirmBuy(c controller.MContext)

	// @Title ConfirmSell
	// @Description ConfirmSell Needs
	ConfirmSell(c controller.MContext)
}