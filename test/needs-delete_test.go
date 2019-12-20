package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
	"reflect"
	"strconv"
)

func testNeedsDelete(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(needsEsDeleteIdK)).Convert(intType).Interface().(int)

	resp := t.Delete("/v1/needs/" + strconv.Itoa(id))
	t.AssertNoError(false)
	resp = t.Get("/v1/needs/" + strconv.Itoa(id))
	NotFound := t.FetchError(resp)
	t.Equal(NotFound.RespCode, http.StatusOK, "must result in 200")
	t.Equal(NotFound.Code, types.CodeNotFound, "not be the ErrNotFound")

}
