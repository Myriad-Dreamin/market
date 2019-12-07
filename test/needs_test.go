package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestNeeds(t *testing.T) {

	_ = t.Run("TestNeedsPost", testNeedsPost) &&
		t.Run("TestNeedsPostWithError", testNeedsPostWithError) &&
		t.Run("TestNeedsGet", testNeedsGet) &&
		t.Run("TestNeedsGetList", testNeedsGetList) &&
		t.Run("TestNeedsFilters", srv.HandleTestWithOutError(testNeedsFilters)) &&
		t.Run("TestNeedsDelete", srv.HandleTestWithOutError(testNeedsDelete))
}

func testNeedsFilters(t *tester.TesterContext) {

}
