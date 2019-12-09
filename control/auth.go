package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

/* auth
 * refresh token GET: 刷新登陆用token
 */
type AuthService interface {
	AuthSignatureXXX() interface{}
	// /refresh-token GET
	RefreshToken(c controller.MContext)

	//AddGroup(c controller.MContext)
	//DeleteGroup(c controller.MContext)

	// grant
	//Grant(c controller.MContext)
	//Revoke(c controller.MContext)
	//Check(c controller.MContext)
	//
	//GrantGroup(c controller.MContext)
	//RevokeGroup(c controller.MContext)
	//CheckGroup(c controller.MContext)
	//
	//ReadGranter(entityInterface auth.ReadEntityInterface) controller.HandlerFunc
	//ReadRevoker(entityInterface auth.ReadEntityInterface) controller.HandlerFunc
	//ReadChecker(entityInterface auth.ReadEntityInterface) controller.HandlerFunc
	//
	//WriteGranter(entityInterface auth.WriteEntityInterface) controller.HandlerFunc
	//WriteRevoker(entityInterface auth.WriteEntityInterface) controller.HandlerFunc
	//WriteChecker(entityInterface auth.WriteEntityInterface) controller.HandlerFunc
	//
	//JustSimpleWriteGranter(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc
	//JustSimpleWriteRevoker(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc
	//JustSimpleWriteChecker(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc
}