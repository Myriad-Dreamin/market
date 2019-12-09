package tests

import (
	"fmt"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"testing"
)

func testNeedsGet(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	resp := srv.Get("/v1/needs/1")
	reply := srv.DecodeJSON(resp.Body(), new(needsservice.GetReply)).(*needsservice.GetReply)
	fmt.Println(reply)
}
