package router

import (
	"github.com/Myriad-Dreamin/market/control/auth"
)

func ApplyAuth(router *RootRouter) {
	var agi = router.AuthApiRouter.Group
	agi.RevokeGroup.Use(agi.Auth.AdminOnly())
	agi.GrantGroup.Use(agi.Auth.AdminOnly())

	var aggMap = router.AuthApiRouter.Sugar.DynamicList
	for _, agg := range aggMap {
		agg.Revoke.Use(agg.Auth.AdminOnly())
		agg.Grant.Use(agg.Auth.AdminOnly())
	}
	//agi.CheckGroup.Use(agi.Auth.AdminOnly())

	var uig = router.UserRouter.IDRouter
	uig.ChangePassword.Use(uig.Auth.Build(auth.UserEntity.Write()))
	uig.Put.Use(uig.Auth.Build(auth.UserEntity.Write()))
	uig.Delete.Use(uig.Auth.AdminOnly())

	var gig = router.GoodsRouter.IDRouter
	gig.Inspect.Use(gig.Auth.Build(auth.GoodsEntity.Read()))
	gig.Put.Use(gig.Auth.Build(auth.GoodsEntity.Write()))
	gig.Delete.Use(gig.Auth.Build(auth.GoodsEntity.Write()))
	gig.ForceDelete.Use(gig.Auth.AdminOnly())
	gig.PutPicture.Use(gig.Auth.Build(auth.GoodsEntity.Write()))

	//
	var nig = router.NeedsRouter.IDRouter
	nig.Inspect.Use(nig.Auth.Build(auth.NeedsEntity.Read()))
	nig.Put.Use(nig.Auth.Build(auth.NeedsEntity.Write()))
	nig.Delete.Use(nig.Auth.Build(auth.NeedsEntity.Write()))
	nig.ForceDelete.Use(nig.Auth.AdminOnly())
	nig.PutPicture.Use(nig.Auth.Build(auth.NeedsEntity.Write()))

}
