package tests

import (
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/types"
	"strconv"
	"testing"
	"time"
)


func RangeInt(l,r int) []int {
	var x = make([]int, r-l)
	for i := l; i < r; i++ {
		x[i-l] = i
	}
	return x
}

func testGoodsPost(t *testing.T) {
	ctx := srv.Context(t).AssertNoError(true)
	var (
		end_at  = time.Now().Add(time.Hour * 24)
		tp = types.GoodsTypeElectronic
		name = "es000"
		min_price uint64 = 100
		is_fixed = false
	)
	_ = ctx.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       end_at,
		Type:        tp,
		Name:        name,
		MinPrice:    min_price,
		IsFixed:     is_fixed,
		Description: "",
	})
	for i := range RangeInt(0, 10) {
		resp := ctx.Post("/v1/goods", goodsservice.PostRequest{
			EndAt:       end_at,
			Type:        tp,
			Name:        name + strconv.Itoa(i),
			MinPrice:    min_price,
			IsFixed:     is_fixed,
			Description: "",
		})
		var x goodsservice.PostReply
		ctx.HandlerError0(resp.JSON(&x))
		if i == 1 {
			srv.Set(goodsEsDeleteIdK, x.Goods.ID)
		} else if i == 2 {
			srv.Set(goodsEsIdK, x.Goods.ID)
		}
	}

}

func testGoodsPostWithError(t *testing.T) {
	ctx := srv.Context(t).AssertNoError(false)
	ctx.Cfg.BaseParametersConfig.GoodsMinimumEndDuration = time.Second * 10
	var (
		end_at  = time.Now().Add(time.Hour * 24)
		tp = types.GoodsTypeElectronic
		min_price uint64 = 100
		is_fixed = false
	)
	resp := ctx.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       end_at,
		Type:        types.GoodsTypeUnknown,
		Name:        "es0",
		MinPrice:    min_price,
		IsFixed:     is_fixed,
		Description: "",
	})
	err := ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error("code not be 200: ", err)
	}
	resp = ctx.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       end_at,
		Type:        tp,
		Name:        "",
		MinPrice:    min_price,
		IsFixed:     is_fixed,
		Description: "",
	})
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       end_at,
		Type:        tp,
		Name:        "",
		MinPrice:    min_price,
		IsFixed:     is_fixed,
		Description: "",
	})
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/goods", map[string]interface{}{
		"end_at": end_at,
		"g_type": tp,
		"name": "es1",
		"min_price": -1,
		"is_fixed": false,
	})
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/goods", goodsservice.PostRequest{
		EndAt:       time.Now(),
		Type:        tp,
		Name:        "es0",
		MinPrice:    min_price,
		IsFixed:     is_fixed,
		Description: "",
	})
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}
}
