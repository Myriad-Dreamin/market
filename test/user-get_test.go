package tests

import (
	"fmt"
	userservice "github.com/Myriad-Dreamin/market/service/user"
	"testing"
)

func testGet(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	srv.PrintRequest(true)
	resp := srv.Get("/v1/user/1")
	reply := srv.DecodeJSON(resp.Body(), new(userservice.GetReply)).(*userservice.GetReply)
	fmt.Println(reply, reply.User)
}
