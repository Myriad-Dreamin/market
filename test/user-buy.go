package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/market/types"
	"strconv"
	"time"
)

func testUserBuy(t *tester.TesterContext) {
	resp := t.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       time.Now().Add(time.Hour),
		Type:        types.GoodsTypeElectronic,
		Name:        "es000000000000000",
		MinPrice:    1,
		IsFixed:     false,
		Description: "",
	})
	var x goodsservice.PostReply
	t.HandlerError0(resp.JSON(&x))


	resp = t.Post("/v1/user/1/goods/" + strconv.Itoa(int(x.Goods.ID)) + "/buy")
	fmt.Println(resp)

	resp = t.Get("/v1/goods/" + strconv.Itoa(int(x.Goods.ID)))
	reply := t.DecodeJSON(resp.Body(), new(goodsservice.GetReply)).(*goodsservice.GetReply)
	fmt.Println(reply.Goods.Buyer, reply.Goods.Seller, reply.Goods.Status)

	//_ = t.Delete("/v1/goods/" + strconv.Itoa(int(x.Goods.ID)))
}

