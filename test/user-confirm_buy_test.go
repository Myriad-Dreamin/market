package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"reflect"
	"strconv"
)

func testUserConfirmBuy(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(goodsBuyIdK)).Convert(intType).Interface().(int)
	t.Delete("/v1/goods/" + strconv.Itoa(id) + "/force")
}

