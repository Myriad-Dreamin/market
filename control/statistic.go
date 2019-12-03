package control

import "github.com/gin-gonic/gin"

/*
查询一定条件下当前已经成交物品的累计中介费收益信息
统计分析功能模块：显示起始年月、终止年月、某个地域不同物品类型成交中介费的明细，并按月以折线图的方式展示每月累计成交笔数、中介费金额的变化趋势
*/

/* statistic
 * show GET: 统计信息
 */
type StatisticService interface {

	// /stat/goods/fee GET

	// @Title Statistic Goods Fee
	// @Description Statistic Goods Fee
	StatGoodsFee(c *gin.Context)

	// /stat/goods/fee-xy GET

	// @Title Statistic Goods Fee XY List
	// @Description Statistic Goods Fee XY List
	StatGoodsFeeXYList(c *gin.Context)

	// /stat/goods/fee GET

	// @Title Statistic Goods Count XY List
	// @Description Statistic Goods Count XY List
	StatGoodsCountXYList(c *gin.Context)
}

