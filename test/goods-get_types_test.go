package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testGoodsGetTypes(t *tester.TesterContext) {
	resp := t.Get("/v1/goods-types")
	fmt.Println(resp)
}
