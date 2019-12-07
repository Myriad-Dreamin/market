package control

import "github.com/gin-gonic/gin"

//go:generate package-attach-to-path -generate_register_map


/* auth
 * refresh token GET: 刷新登陆用token
 */
type Auth interface {
	// /refresh-token GET
	RefreshToken(c *gin.Context)
}

type ObjectService interface {
	ObjectSignatureXXX() interface{}

	// @Title Post
	// @Description Post Object
	Post(c *gin.Context)

	// @Title Put
	// @Description Put Object
	Put(c *gin.Context)

	// @Title Delete
	// @Description Delete Object
	Delete(c *gin.Context)

	// @Title Get
	// @Description Get Object
	Get(c *gin.Context)

	// @Title List
	// @Description List Object
	List(c *gin.Context)

}
