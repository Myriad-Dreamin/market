package tests

import (
	"fmt"
	"testing"
)

func testGoodsGet(t *testing.T) {
	fmt.Println(srv.Get(goodsEsIdK))
}

