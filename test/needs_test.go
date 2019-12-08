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

func TestNeeds(t *testing.T) {

	_ = t.Run("Post", testNeedsPost) &&
		t.Run("PostWithError", testNeedsPostWithError) &&
		t.Run("Get", testNeedsGet) &&
		t.Run("Inspect", srv.HandleTestWithOutError(testNeedsInspect)) &&
		t.Run("Put", srv.HandleTestWithOutError(testNeedsPut)) &&
		t.Run("GetList", testNeedsGetList) &&
		t.Run("Filters", srv.HandleTestWithOutError(testNeedsFilters)) &&
		t.Run("UploadPicture", srv.HandleTestWithOutError(testNeedsUploadPicture)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testNeedsDelete))
}

func testNeedsUploadPicture(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(needsEsIdK)).Convert(intType).Interface().(int)
	putURL := "/v1/needs/" + strconv.Itoa(id) + "/picture"
	w := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(w)
	t.HandlerError(writer.CreateFormFile("upload", "1.png"))
	w.Write([]byte("456"))
	t.HandlerError0(writer.Close())

	t.Put(putURL, mock.NewNamedReader(writer.FormDataContentType(), w),
		mock.Comment(fmt.Sprintf(
			"url(%s), put picture to server", putURL)))

}

func testNeedsPut(t *tester.TesterContext) {

}