package tests

import (
	"fmt"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/market/types"
	"strconv"
	"time"
)

func testUserSell(t *tester.TesterContext) {
	resp := t.Post("/v1/needs", needsservice.PostRequest{
		EndAt:       time.Now().Add(time.Hour),
		Type:        types.GoodsTypeElectronic,
		Name:        "es000000000000000",
		MaxPrice:     2,
		Description: "",
	})
	var x needsservice.PostReply
	t.HandlerError0(resp.JSON(&x))
	id := x.Needs.ID


	resp = t.Post("/v1/user/1/needs/" + strconv.Itoa(int(id)) + "/sell")

	fmt.Println(resp)

	resp = t.Get("/v1/needs/" + strconv.Itoa(int(id)) + "/inspect")
	reply := t.DecodeJSON(resp.Body(), new(needsservice.InspectReply)).(*needsservice.InspectReply)
	fmt.Println(reply.Needs.Buyer, reply.Needs.Seller, reply.Needs.Status)

	//_ = t.Delete("/v1/needs/" + strconv.Itoa(int(id)))
	srv.Set(needsSellIdK, id)
}
