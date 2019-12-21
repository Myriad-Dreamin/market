package tests

import (
	"fmt"
	authservice "github.com/Myriad-Dreamin/market/service/auth"
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestAuth(t *testing.T) {
	_ =
		t.Run("CheckGroup", srv.HandleTestWithOutError(testAuthCheckGroup)) &&
		t.Run("GrantGroup", srv.HandleTestWithOutError(testAuthGrantGroup)) &&
		t.Run("RevokeGroup", srv.HandleTestWithOutError(testAuthRevokeGroup))
}

func testAuthCheckGroup(t *tester.TesterContext) {
	resp := t.Get("/v1/auth/group/user/1?group=admin")
	fmt.Println(resp)
}

func testAuthRevokeGroup(t *tester.TesterContext) {
	t.Delete("/v1/auth/group/user/1", authservice.GroupRequest{GroupName:"__test_group__"})
	t.Equal(false, t.DecodeJSON(t.Get("/v1/auth/group/user/1?group=__test_group__").Body(), new(authservice.CheckReply)).
	(*authservice.CheckReply).Has, "must be removed")
}

func testAuthGrantGroup(t *tester.TesterContext) {
	t.Put("/v1/auth/group/user/1", authservice.GroupRequest{GroupName:"__test_group__"})
	t.Equal(true, t.DecodeJSON(t.Get("/v1/auth/group/user/1?group=__test_group__").Body(), new(authservice.CheckReply)).
	(*authservice.CheckReply).Has, "must be added")
}



