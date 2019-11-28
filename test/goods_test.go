package tests

import "testing"


func TestGoods(t *testing.T) {
	_ = t.Run("TestGoodsPost", testGoodsPost) &&
		t.Run("TestGoodsPostWithError", testGoodsPostWithError) &&
		t.Run("TestGoodsGet", testGoodsGet) &&
		t.Run("TestGoodsDelete", testGoodsDelete) &&
		t.Run("TestGoodsList", testGoodsList)
}