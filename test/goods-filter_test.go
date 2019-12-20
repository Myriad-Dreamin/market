package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
)

func testGoodsFilters(t *tester.TesterContext) {
	t.PrintRequest(true)
	queryWithName := "/v1/goods-list?page=1&page_size=20&name=es0001"
	resp := t.Get(queryWithName, mock.Comment(
		fmt.Sprintf(`url(%s), query the goods that with name es0001`, queryWithName)))

	reply := t.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)

	resp = t.Get("/v1/goods-list?page=2&page_size=20")

	reply = t.DecodeJSON(resp.Body(), new(goodsservice.ListReply)).(*goodsservice.ListReply)
	fmt.Println(reply)
	t.PrintRequest(false)
}
