package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestStatistic(t *testing.T) {

	_ = t.Run("TestGoodsStatisticCountXY", srv.HandleTestWithOutError(testGoodsStatisticCountXY)) &&
		t.Run("TestGoodsStatisticFeeXY", srv.HandleTestWithOutError(testGoodsStatisticFeeXY)) &&
		t.Run("TestGoodsStatisticFee", srv.HandleTestWithOutError(testGoodsStatisticFee))
}

func testGoodsStatisticFee(t *tester.TesterContext) {

}

func testGoodsStatisticFeeXY(t *tester.TesterContext) {

}

func testGoodsStatisticCountXY(t *tester.TesterContext) {

}
