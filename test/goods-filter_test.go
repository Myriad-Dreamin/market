package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testGoodsFilters(t *tester.TesterContext) {
	t.PrintRequest(true)
	resp := t.Get("/v1/goods-list?page=1&page_size=20&name=es0001")

	reply := t.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)

	resp = t.Get("/v1/goods-list?page=2&page_size=20")

	reply = t.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)
	t.PrintRequest(false)
}

