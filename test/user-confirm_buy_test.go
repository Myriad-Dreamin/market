package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
	"reflect"
	"strconv"
)

func testUserConfirmBuy(t *tester.TesterContext) {

	id := reflect.ValueOf(srv.Get(goodsBuyIdK)).Convert(intType).Interface().(int)

	resp := t.Post("/v1/user/1/goods/" + strconv.Itoa(int(id)) + "/confirm-buy")
	fmt.Println(resp)

	resp = t.Get("/v1/goods/" + strconv.Itoa(int(id)))
	reply := t.DecodeJSON(resp.Body(), new(goodsservice.GetReply)).(*goodsservice.GetReply)
	fmt.Println(reply.Goods.Buyer, reply.Goods.Seller, reply.Goods.Status)
	t.Delete("/v1/goods/" + strconv.Itoa(id) + "/force")
}

