package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

// UserService defines the interface of user service

type UserService interface {
	UserSignatureXXX() interface{}

	// @Title Login
	// @Description Login to access more api in backend.
	// The following is a description of the parameters
	//     + `password string`: password of the user
	// uid, one of following three parameter must be provided:
	//     + `id uint`: id of the user
	//
	//     + `nick_name string`: nick-name of the user
	//
	//     + `phone string` phone number of the user
	//
	Login(c controller.MContext)

	// @Title Register
	// @Description Register an user in market server
	// The following is a description of the parameters
	//     + `name string`: name of the user
	//
	//     + `nick_name string`: nick-name of the user
	//
	//     + `phone string` phone number of the user
	//
	//     + `password string` password number of the user, which must pass the [password test](https://github.com/Myriad-Dreamin/market/blob/master/service/user/change-password.go).
	//
	//     + `register_city string`: register city of the user
	//
	Register(c controller.MContext)

	// @Title Put
	// @Description Update User Information. only the phone number can be modified
	// The following is a description of the parameters
	//     + `phone string` phone number of the user
	Put(c controller.MContext)

	// @Title ChangePassword
	// @Description change password of user
	// The following is a description of the parameters
	//     + `old_password string` old password number of the user, must match the string in database after encrypted
	//
	//     + `new_password string` new password number of the user, which must pass the [password test](https://github.com/Myriad-Dreamin/market/blob/master/service/user/change-password.go).
	//
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
