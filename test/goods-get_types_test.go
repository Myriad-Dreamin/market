package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testGoodsGetTypes(t *tester.TesterContext) {
	resp := t.Get("/v2/const/goods-types")
	fmt.Println(resp)
}
