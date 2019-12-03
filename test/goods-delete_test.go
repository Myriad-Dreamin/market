package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
	"reflect"
	"strconv"
)

var intT = 1
var intType = reflect.TypeOf(intT)

func testGoodsDelete(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(goodsEsDeleteIdK)).Convert(intType).Interface().(int)

	resp := t.Delete("/v1/goods/" + strconv.Itoa(id))
	t.AssertNoError(false)
	resp = t.Get("/v1/goods/" + strconv.Itoa(id))
	fmt.Println(resp)
}