package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
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

func testGoodsFilters(t *tester.TesterContext) {
	srv := srv.Context(t).AssertNoError(true)

	resp := srv.Get("/v1/goods-list?page=1&page_size=2")

	reply := srv.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)

	resp = srv.Get("/v1/goods-list?page=2&page_size=2")

	reply = srv.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)
}
