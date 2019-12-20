package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	userservice "github.com/Myriad-Dreamin/market/service/user"
	"github.com/Myriad-Dreamin/market/test/tester"
	"reflect"
	"strconv"
)

func testUserConfirmBuy(t *tester.TesterContext) {

	id := reflect.ValueOf(srv.Get(goodsBuyIdK)).Convert(intType).Interface().(int)

	resp := t.Post("/v1/user/1/goods/"+strconv.Itoa(int(id))+"/confirm-buy", userservice.ConfirmBuyRequest{
		Confirm: true})
	fmt.Println(resp)

	resp = t.Get("/v1/goods/" + strconv.Itoa(int(id)) + "/inspect")
	reply := t.DecodeJSON(resp.Body(), new(goodsservice.InspectReply)).(*goodsservice.InspectReply)
	fmt.Println(reply.Goods.Buyer, reply.Goods.Seller, reply.Goods.Status)
	t.Delete("/v1/goods/" + strconv.Itoa(id) + "/force")
}
