package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testGoodsInspect(t *tester.TesterContext) {
	resp := t.Get("/v1/goods/1/inspect")
	reply := t.DecodeJSON(resp.Body(), new(goodsservice.InspectReply)).(*goodsservice.InspectReply)
	fmt.Println(reply)
}
