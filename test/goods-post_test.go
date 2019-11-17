package tests

import (
	"fmt"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/types"
	"testing"
	"time"
)

func testGoodsPost(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	var (
		end_at  = time.Now().Add(time.Hour * 24)
		tp = types.GoodsTypeElectronic
		name = "es"
		min_price uint64 = 100
		is_fixed = false
	)
	resp := srv.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       end_at,
		Type:        tp,
		Name:        name,
		MinPrice:    min_price,
		IsFixed:     is_fixed,
		Description: "",
	})
	fmt.Println(resp)
}
