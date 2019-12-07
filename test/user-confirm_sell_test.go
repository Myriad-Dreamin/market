package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"reflect"
	"strconv"
)

func testUserConfirmSell(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(needsSellIdK)).Convert(intType).Interface().(int)
	t.Delete("/v1/needs/" + strconv.Itoa(id) + "/force")

}

