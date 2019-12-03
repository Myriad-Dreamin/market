package tests

import (
	"fmt"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"testing"
)

func testNeedsGetList(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)

	resp := srv.Get("/v1/needs-list?page=1&page_size=2")

	reply := srv.DecodeJSON(resp.Body(), new(needsservice.ListReply)).(*needsservice.ListReply)
	fmt.Println(reply)


	resp = srv.Get("/v1/needs-list?page=2&page_size=2")

	reply = srv.DecodeJSON(resp.Body(), new(needsservice.ListReply)).(*needsservice.ListReply)
	fmt.Println(reply)
}
