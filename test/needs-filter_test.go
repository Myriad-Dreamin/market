package tests

import (
	"fmt"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
)

func testNeedsFilters(t *tester.TesterContext) {
	t.PrintRequest(true)
	resp := t.Get("/v1/needs-list?page=1&page_size=20&name=es0001")

	reply := t.DecodeJSON(resp.Body(), new(needsservice.ListReply)).(*needsservice.ListReply)
	fmt.Println(reply)

	queryWithType := "/v1/needs-list?page=1&page_size=20&type=1"
	resp = t.Get(queryWithType, mock.Comment(`query the needs that with type 1`))

	reply = t.DecodeJSON(resp.Body(), new(needsservice.ListReply)).(*needsservice.ListReply)
	fmt.Println(reply)

	resp = t.Get("/v1/needs-list?page=2&page_size=20")

	reply = t.DecodeJSON(resp.Body(), new(needsservice.ListReply)).(*needsservice.ListReply)
	fmt.Println(reply)
	t.PrintRequest(false)
}
