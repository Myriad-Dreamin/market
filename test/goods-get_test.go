package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"testing"
)

func testGoodsGet(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	resp := srv.Get("/v1/goods/1")
	reply := srv.DecodeJSON(resp.Body(), new(goodsservice.GetReply)).(*goodsservice.GetReply)
	fmt.Println(reply, reply.Goods)
}
