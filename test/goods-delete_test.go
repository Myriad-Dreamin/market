package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"net/http"
	"reflect"
	"strconv"
)

func testGoodsDelete(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(goodsEsDeleteIdK)).Convert(intType).Interface().(int)
	deleteURL := "/v1/goods/" + strconv.Itoa(id)
	t.Delete(deleteURL,
		mock.Comment(fmt.Sprintf(
			"url(%s), delete the goods which have not be bought", deleteURL)))
	t.AssertNoError(false)
	resp := t.Get("/v1/goods/"+strconv.Itoa(id), mock.AbortRecord(true))
	NotFound := t.FetchError(resp)
	t.Equal(NotFound.RespCode, http.StatusOK, "must result in 200")
	t.Equal(NotFound.Code, types.CodeNotFound, "not be the ErrNotFound")

}
