package control

import (
	"github.com/Myriad-Dreamin/market/control/auth"
	"github.com/Myriad-Dreamin/market/lib/controller"
)

var AuthCates []interface{}

// @Category Auth - Group Grant/Revoke/Check Api Group
// @Description Grant Api Group
// @Path /v1/auth/group/user/:id
type authGrantApiGroupCate interface {
}

// @Category Auth - Sugar - Group Admin Grant/Revoke/Check Api Group
// @Description Grant Api Group
// @Path /v1/auth/sugar/group/admin/user/:id
type authGrantAdminApiGroupCate interface {
}

func init() {
	var (
		a authGrantApiGroupCate      = 0
		b authGrantAdminApiGroupCate = 0
	)
	AuthCates = []interface{}{
		&a,
		&b,
	}
}

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
	Grant(c controller.MContext)
	Revoke(c controller.MContext)
	Check(c controller.MContext)

	// @Title Group Grants
	// @Description Group Grants
	GrantGroup(c controller.MContext)

	// @Title Group Revokes
	// @Description Group Revokes
	RevokeGroup(c controller.MContext)

	// @Title Group Checks
	// @Description Group Checks
	CheckGroup(c controller.MContext)

	// @Title Group Grants
	// @Description Group Grants
	GroupGranter(groupName string) controller.HandlerFunc

	// @Title Group Revokes
	// @Description Group Revokes
	GroupRevoker(groupName string) controller.HandlerFunc

	// @Title Group Checks
	// @Description Group Checks
	GroupChecker(groupName string) controller.HandlerFunc

	ReadGranter(entityInterface auth.ReadEntityInterface) controller.HandlerFunc
	ReadRevoker(entityInterface auth.ReadEntityInterface) controller.HandlerFunc
	ReadChecker(entityInterface auth.ReadEntityInterface) controller.HandlerFunc

	WriteGranter(entityInterface auth.WriteEntityInterface) controller.HandlerFunc
	WriteRevoker(entityInterface auth.WriteEntityInterface) controller.HandlerFunc
	WriteChecker(entityInterface auth.WriteEntityInterface) controller.HandlerFunc

	JustSimpleWriteGranter(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc
	JustSimpleWriteRevoker(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc
	JustSimpleWriteChecker(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc
}
