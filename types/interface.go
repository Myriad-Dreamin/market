package types



/* auth
 * refresh token GET: 刷新登陆用token
 */
type Auth interface {
	// /refresh-token GET
	RefreshToken()
}

/* user
 * login POST 登陆
 * register POST 注册
 * put PUT 更新用户信息
 * change password PUT 修改用户密码
 * grant PUT 授予用户身份或权限
 * get GET 获取用户信息
 * delete DELETE 删除用户
 */
type User interface {
	// /:id POST
	Login()
	// / POST
	Register()
	// /:id PUT
	Put()
	// /:id/password PUT
	ChangePassword()
	// /:id/grant PUT
	Grant()
	// /:id GET
	Get()
	// /:id DELETE
	Delete()
	// /:id POST
}

/* goods
 * post: 登记商品
 * put: 修改商品信息
 * delete: 删除商品信息
 * get: 获取商品信息
 * list: 根据filter得到商品列表
 */
type Goods interface {
	// / POST
	Post()
	// / PUT
	Put()
	// /:id DELETE
	Delete()
	// /:id GET
	Get()
	// / GET
	List()
}

/* statistic
 * show GET: 统计信息
 */
type Statistic interface {
	// /show GET
	Show()
}




