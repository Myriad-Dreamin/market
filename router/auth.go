package router

import "github.com/Myriad-Dreamin/market/router/auth"

func ApplyAuth(router *RootRouter) {
	var uig = router.UserRouter.IDRouter
	uig.ChangePassword.Use(uig.Auth.Build(auth.UserEntity.Write()))
	uig.Put.Use(uig.Auth.Build(auth.UserEntity.Write()))
	uig.Delete.Use(uig.Auth.AdminOnly())

	var gig = router.GoodsRouter.IDRouter
	gig.Put.Use(gig.Auth.Build(auth.GoodsEntity.Write()))
	gig.Delete.Use(gig.Auth.Build(auth.GoodsEntity.Write()))
	//
	var nig = router.NeedsRouter.IDRouter
	nig.Put.Use(nig.Auth.Build(auth.NeedsEntity.Write()))
	nig.Delete.Use(nig.Auth.Build(auth.NeedsEntity.Write()))

}

