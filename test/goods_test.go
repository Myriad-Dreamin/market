package tests

import "testing"


func TestGoods(t *testing.T) {
	_ = t.Run("TestGoodsPost", testGoodsPost) &&
		t.Run("TestGoodsPostWithError", testGoodsPostWithError)
}

