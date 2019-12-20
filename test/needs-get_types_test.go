package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testNeedsGetTypes(t *tester.TesterContext) {
	resp := t.Get("/v1/needs-types")
	fmt.Println(resp)
}
