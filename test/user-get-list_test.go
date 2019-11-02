package tests

import (
	"fmt"
	userservice "github.com/Myriad-Dreamin/market/service/user"
	"testing"
)

func testList(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)

	resp := srv.Get("/v1/user-list?page=1&page_size=1")

	reply := srv.DecodeJSON(resp.Body(), new(userservice.ListReply)).(*userservice.ListReply)
	fmt.Println(reply)
}
