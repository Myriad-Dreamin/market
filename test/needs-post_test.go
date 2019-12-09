package tests

import (
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"strconv"
	"testing"
	"time"
)


func testNeedsPost(t *testing.T) {
	ctx := srv.Context(t).AssertNoError(true)
	var (
		end_at  = time.Now().Add(time.Hour * 24)
		tp = types.GoodsTypeElectronic
		name = "es000"
		min_price uint64 = 100
		max_price uint64 = 100
	)
	_ = ctx.Post("/v1/needs", needsservice.PostRequest{
		EndAt:       end_at,
		Type:        tp,
		Name:        name,
		MinPrice:    min_price,
		MaxPrice:    max_price,
		Description: "",
	})
	for i := range RangeInt(0, 10) {
		resp := ctx.Post("/v1/needs", needsservice.PostRequest{
			EndAt:       end_at,
			Type:        tp,
			Name:        name + strconv.Itoa(i),
			MinPrice:    min_price,
			MaxPrice:     max_price,
			Description: "",
		})
		var x needsservice.PostReply
		ctx.HandlerError0(resp.JSON(&x))
		if i == 1 {
			srv.Set(needsEsDeleteIdK, x.Needs.ID)
		} else if i == 2 {
			srv.Set(needsEsIdK, x.Needs.ID)
		}
	}
}

func testNeedsPostWithError(t *testing.T) {
	ctx := srv.Context(t).AssertNoError(false)
	var (
		end_at  = time.Now().Add(time.Hour * 24)
		tp = types.GoodsTypeElectronic
		min_price uint64 = 100
		max_price uint64 = 100
	)
	resp := ctx.Post("/v1/needs", needsservice.PostRequest{
		EndAt:       end_at,
		Type:        types.GoodsTypeUnknown,
		Name:        "es0",
		MinPrice:    min_price,
		MaxPrice:     max_price,
		Description: "",
	}, mock.Comment("type invalid"))
	err := ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error("code not be 200: ", err)
	}

	resp = ctx.Post("/v1/needs", needsservice.PostRequest{
		EndAt:       end_at,
		Type:        tp,
		Name:        "",
		MinPrice:    min_price,
		MaxPrice:     max_price,
		Description: "",
	}, mock.Comment("name required"))
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/needs", map[string]interface{}{
		"end_at": end_at,
		"g_type": tp,
		"name": "es1",
		"min_price": -1,
		"max_price": 1,
	}, mock.Comment("min price negative"))
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/needs", map[string]interface{}{
		"end_at": end_at,
		"g_type": tp,
		"name": "es1",
		"min_price": 1,
		"max_price": -1,
	}, mock.Comment("max price negative"))
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/needs", map[string]interface{}{
		"end_at": end_at,
		"g_type": tp,
		"name": "es1",
		"min_price": 2,
		"max_price": 1,
	}, mock.Comment("min price bigger than max price"))
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}

	resp = ctx.Post("/v1/needs", needsservice.PostRequest{
		EndAt:       time.Now(),
		Type:        tp,
		Name:        "es0",
		MinPrice:    min_price,
		MaxPrice:     max_price,
		Description: "",
	}, mock.Comment("short alive duration"))
	err = ctx.FetchError(resp)
	if err.RespCode != 200 || err.Code != types.CodeInvalidParameters {
		t.Error(err)
	}
}
