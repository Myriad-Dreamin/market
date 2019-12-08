package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
	"net/url"
	"testing"
	"time"
)

func TestStatistic(t *testing.T) {

	_ = t.Run("CountXY", srv.HandleTestWithOutError(testGoodsStatisticCountXY)) &&
		t.Run("FeeXY", srv.HandleTestWithOutError(testGoodsStatisticFeeXY))
}


func testGoodsStatisticFeeXY(t *tester.TesterContext) {
	A := time.Now()
	resp := t.Get(fmt.Sprintf(
		`/v1/statistic/fee?le=%s&ge=%s`, url.QueryEscape(A.Format(time.RFC3339)), url.QueryEscape(A.Format(time.RFC3339))))
	fmt.Println(resp)
}

func testGoodsStatisticCountXY(t *tester.TesterContext) {
	A := time.Now()
	resp := t.Get(fmt.Sprintf(
		`/v1/statistic/count?le=%s&ge=%s`, url.QueryEscape(A.Format(time.RFC3339)), url.QueryEscape(A.Format(time.RFC3339))))
	fmt.Println(resp)
}
