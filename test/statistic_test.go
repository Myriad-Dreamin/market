package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestStatistic(t *testing.T) {

	_ = t.Run("TestGoodsStatisticCountXY", srv.HandleTestWithOutError(testGoodsStatisticCountXY)) &&
		t.Run("TestGoodsStatisticFeeXY", srv.HandleTestWithOutError(testGoodsStatisticFeeXY)) &&
		t.Run("TestGoodsStatisticFee", srv.HandleTestWithOutError(testGoodsStatisticFee)) &&
		t.Run("TestNeedsStatisticCountXY", srv.HandleTestWithOutError(testNeedsStatisticCountXY)) &&
		t.Run("TestNeedsStatisticFeeXY", srv.HandleTestWithOutError(testNeedsStatisticFeeXY)) &&
		t.Run("TestNeedsStatisticFee", srv.HandleTestWithOutError(testNeedsStatisticFee)) &&
		t.Run("TestTotalStatisticCountXY", srv.HandleTestWithOutError(testTotalStatisticCountXY)) &&
		t.Run("TestTotalStatisticFeeXY", srv.HandleTestWithOutError(testTotalStatisticFeeXY)) &&
		t.Run("TestTotalStatisticFee", srv.HandleTestWithOutError(testTotalStatisticFee))
}

func testTotalStatisticFee(t *tester.TesterContext) {

}

func testTotalStatisticFeeXY(t *tester.TesterContext) {

}

func testTotalStatisticCountXY(t *tester.TesterContext) {

}

func testNeedsStatisticFee(t *tester.TesterContext) {

}

func testNeedsStatisticFeeXY(t *tester.TesterContext) {

}

func testNeedsStatisticCountXY(t *tester.TesterContext) {

}

func testGoodsStatisticFee(t *tester.TesterContext) {

}

func testGoodsStatisticFeeXY(t *tester.TesterContext) {

}

func testGoodsStatisticCountXY(t *tester.TesterContext) {

}
