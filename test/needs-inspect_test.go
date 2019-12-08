package tests

import (
	"fmt"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"github.com/Myriad-Dreamin/market/test/tester"
)

func testNeedsInspect(t *tester.TesterContext) {
	resp := t.Get("/v1/needs/1/inspect")
	reply := t.DecodeJSON(resp.Body(), new(needsservice.GetReply)).(*needsservice.GetReply)
	fmt.Println(reply)
}
