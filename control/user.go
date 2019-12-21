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
	//         + `id uint`: id of the user
	//
	//         + `nick_name string`: nick-name of the user
	//
	//         + `phone string` phone number of the user
	//
	// The following is a description of the returns
	//
	//     + `identity array[string]`: the groups user currently in
	//
	//     + `id uint`: the unique user id in database
	//
	//     + `phone string`: the phone of user
	//
	//     + `nick_name string`: the unique user nick name in this website
	//
	//     + `name string`: the true name of user
	//
	//     + `token string`: the jwt token identifies an user, which must set in every auth api request's header
	//
	//     + `refresh_token string`: the jwt token used to refresh token without requirement of user inputting password to login again
	Login(c controller.MContext)

	// @Title Register
	// @Description Register an user in market server
	// The following is a description of the parameters
	//     + `name string`: name of the user
	//
	//     + `nick_name string`: nick-name of the user, must be unique
	//
	//     + `phone string` phone number of the user, must be unique
	//
	//     + `password string` password number of the user, which must pass the [password test](https://github.com/Myriad-Dreamin/market/blob/master/service/user/change-password.go).
	//
	//     + `register_city string`: register city of the user
	//
	// The following is a description of the returns
	//
	//     + `id uint`: the unique user id in database
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

	// @Title GetCities
	// @Description GetCities Needs
	GetCities(c controller.MContext)
}
