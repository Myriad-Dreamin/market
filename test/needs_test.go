package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestNeeds(t *testing.T) {

	_ = t.Run("Post", testNeedsPost) &&
		t.Run("PostWithError", testNeedsPostWithError) &&
		t.Run("Get", testNeedsGet) &&
		t.Run("Put", srv.HandleTestWithOutError(testNeedsPut)) &&
		t.Run("GetList", testNeedsGetList) &&
		t.Run("Filters", srv.HandleTestWithOutError(testNeedsFilters)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testNeedsDelete))
}

func testNeedsPut(t *tester.TesterContext) {

}