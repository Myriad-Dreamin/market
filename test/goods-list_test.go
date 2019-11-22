package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"testing"
)

func testGoodsGetList(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)

	resp := srv.Get("/v1/goods-list?page=1&page_size=2")

	reply := srv.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)


	resp = srv.Get("/v1/goods-list?page=2&page_size=2")

	reply = srv.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)
}
