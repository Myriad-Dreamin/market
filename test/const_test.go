package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"testing"
)

func TestConst(t *testing.T) {
	_ = t.Run("ServiceCode", srv.HandleTestWithOutError(testConstServiceCode))
}

func testConstServiceCode(t *tester.TesterContext) {
	t.Get("v2/const/service-codes", mock.Comment("get service code mapping from int to string"))
}
