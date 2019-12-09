package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

//go:generate package-attach-to-path -generate_register_map


/* auth
 * refresh token GET: 刷新登陆用token
 */
type Auth interface {
	// /refresh-token GET
	RefreshToken(c controller.MContext)
}

type ObjectService interface {
	ObjectSignatureXXX() interface{}

	// @Title Post
	// @Description Post Object
	Post(c controller.MContext)

	// @Title Put
	// @Description Put Object
	Put(c controller.MContext)

	// @Title Delete
	// @Description Delete Object
	Delete(c controller.MContext)

	// @Title Get
	// @Description Get Object
	Get(c controller.MContext)

	// @Title List
	// @Description List Object
	List(c controller.MContext)

}
