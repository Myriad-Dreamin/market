package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
)

func testUserCities(t *tester.TesterContext) {
	_ = t.Get("/v2/const/user-cities", mock.Comment("return an map from id to city object"))
	//fmt.Println(resp)
}
