package tests

import (
	"testing"
)

const goodsEsIdK = "goods/es/id"
const goodsEsDeleteIdK = "goods/es/id"

func TestGoods(t *testing.T) {

	_ = t.Run("TestGoodsPost", testGoodsPost) &&
		t.Run("TestGoodsPostWithError", testGoodsPostWithError) &&
		t.Run("TestGoodsGet", testGoodsGet) &&
		t.Run("TestGoodsGetList", testGoodsGetList) &&
		t.Run("TestGoodsFilters", srv.HandleTestWithOutError(testGoodsFilters)) &&
		t.Run("TestGoodsDelete", srv.HandleTestWithOutError(testGoodsDelete))
}
