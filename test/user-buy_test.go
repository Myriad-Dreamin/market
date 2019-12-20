package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func testUserBuy(t *tester.TesterContext) {
	resp := t.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       time.Now().Add(time.Hour),
		Type:        types.GoodsTypeElectronic,
		Name:        "es000000000000000",
		MinPrice:    100000,
		IsFixed:     true,
		Description: " ",
	})
	var x goodsservice.PostReply
	t.HandlerError0(resp.JSON(&x))
	id := x.Goods.ID

	resp = t.Post("/v1/user/1/goods/"+strconv.Itoa(int(id))+"/buy", gin.H{
		"fixed": true,
	})
	fmt.Println(resp)

	resp = t.Get("/v1/goods/" + strconv.Itoa(int(id)))

	_ = t.DecodeJSON(resp.Body(), new(goodsservice.GetReply)).(*goodsservice.GetReply)
	//fmt.Println(reply.Buyer, reply.Seller, reply.Status)

	//_ = t.Delete("/v1/goods/" + strconv.Itoa(int(id)))
	srv.Set(goodsBuyIdK, id)

	resp = t.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       time.Now().Add(time.Hour),
		Type:        types.GoodsTypeElectronic,
		Name:        "es00000000000000000000",
		MinPrice:    100000,
		IsFixed:     false,
		Description: " ",
	})
	t.HandlerError0(resp.JSON(&x))
	id = x.Goods.ID

	resp = t.Post("/v1/user/1/goods/"+strconv.Itoa(int(id))+"/buy", gin.H{
		"price": 200000,
	})
	fmt.Println(resp)
}
