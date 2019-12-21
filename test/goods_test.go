package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestGoods(t *testing.T) {
	_ = t.Run("Post", testGoodsPost) &&
		t.Run("PostWithError", testGoodsPostWithError) &&
		t.Run("Get", testGoodsGet) &&
		t.Run("Inspect", srv.HandleTestWithOutError(testGoodsInspect)) &&
		t.Run("Put", srv.HandleTestWithOutError(testGoodsPut)) &&
		t.Run("GetList", testGoodsGetList) &&
		t.Run("Filters", srv.HandleTestWithOutError(testGoodsFilters)) &&
		t.Run("UploadPicture", srv.HandleTestWithOutError(testGoodsUploadPicture)) &&
		t.Run("DownloadPicture", srv.HandleTest(testGoodsDownloadPicture)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testGoodsDelete)) &&
		t.Run("GetGoodsTypes", srv.HandleTestWithOutError(testGoodsGetTypes))
}

func testGoodsDownloadPicture(t *tester.TesterContext) {
	fmt.Println(t.Get("/v1/goods-picture/4.png"))
}

func testGoodsPut(t *tester.TesterContext) {

}
