package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testUserConfirmSell(t *tester.TesterContext) {
	//id := reflect.ValueOf(srv.Get(needsSellIdK)).Convert(intType).Interface().(int)
	//
	//resp := t.Post("/v1/user/1/needs/" + strconv.Itoa(int(id)) + "/confirm-sell", userservice.ConfirmSellRequest{
	//	ConfirmOrCancel: true,
	//})
	//fmt.Println(resp)
	//
	//resp = t.Get("/v1/needs/" + strconv.Itoa(int(id)))
	//reply := t.DecodeJSON(resp.Body(), new(needsservice.GetReply)).(*needsservice.GetReply)
	//fmt.Println(reply.Needs.Buyer, reply.Needs.Seller, reply.Needs.Status)
	//t.Delete("/v1/needs/" + strconv.Itoa(id) + "/force")
}

