package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
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
		t.Run("Delete", srv.HandleTestWithOutError(testNeedsDelete)) &&
		t.Run("GetTypes", srv.HandleTestWithOutError(testNeedsGetTypes))
}

func testNeedsPut(t *tester.TesterContext) {

}
