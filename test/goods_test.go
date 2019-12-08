package tests

import (
	"bytes"
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"mime/multipart"
	"reflect"
	"strconv"
	"testing"
)

func TestGoods(t *testing.T) {

	_ = t.Run("Post", testGoodsPost) &&
		t.Run("PostWithError", testGoodsPostWithError) &&
		t.Run("Get", testGoodsGet) &&
		t.Run("Put", srv.HandleTestWithOutError(testGoodsPut)) &&
		t.Run("GetList", testGoodsGetList) &&
		t.Run("Filters", srv.HandleTestWithOutError(testGoodsFilters)) &&
		t.Run("UploadPicture", srv.HandleTestWithOutError(testGoodsUploadPicture)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testGoodsDelete))
}

func testGoodsUploadPicture(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(goodsEsIdK)).Convert(intType).Interface().(int)
	putURL := "/v1/goods/" + strconv.Itoa(id) + "/picture"
	w := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(w)
	t.HandlerError(writer.CreateFormFile("upload", "1.png"))
	w.Write([]byte("22222"))
	t.HandlerError0(writer.Close())

	t.Put(putURL, mock.NewNamedReader(writer.FormDataContentType(), w),
		mock.Comment(fmt.Sprintf(
			"url(%s), put picture to server", putURL)))
}

func testGoodsPut(t *tester.TesterContext) {

}




