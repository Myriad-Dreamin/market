package tests

import (
	"fmt"
	authservice "github.com/Myriad-Dreamin/market/service/auth"
	"github.com/Myriad-Dreamin/market/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"testing"
)

func TestAuth(t *testing.T) {
	_ =
		t.Run("CheckGroup", srv.HandleTestWithOutError(testAuthCheckGroup)) &&
			t.Run("GrantGroup", srv.HandleTestWithOutError(testAuthGrantGroup)) &&
			t.Run("RevokeGroup", srv.HandleTestWithOutError(testAuthRevokeGroup)) &&
			t.Run("GeneratedCheckApiGroup", srv.HandleTestWithOutError(testAuthGeneratedCheckApiGroup)) &&
			t.Run("GeneratedGrantApiGroup", srv.HandleTestWithOutError(testAuthGeneratedGrantApiGroup)) &&
			t.Run("GeneratedRevokeApiGroup", srv.HandleTestWithOutError(testAuthGeneratedRevokeApiGroup))
}

func testAuthGeneratedRevokeApiGroup(t *tester.TesterContext) {
	resp := t.Get("/v1/auth/sugar/group/admin/user/1")
	fmt.Println(resp)
}

func testAuthGeneratedGrantApiGroup(t *tester.TesterContext) {

}

func testAuthGeneratedCheckApiGroup(t *tester.TesterContext) {

}

func testAuthCheckGroup(t *tester.TesterContext) {
	resp := t.Get("/v1/auth/group/user/1?group=admin")
	fmt.Println(resp)
}

func testAuthRevokeGroup(t *tester.TesterContext) {
	t.Delete("/v1/auth/group/user/1", authservice.GroupRequest{GroupName: "__test_group__"}, mock.Comment("remove the user from group with name"))
	t.Equal(false, t.DecodeJSON(t.Get("/v1/auth/group/user/1?group=__test_group__", mock.AbortRecord(true)).Body(), new(authservice.CheckReply)).(*authservice.CheckReply).Has, "must be removed")
}

func testAuthGrantGroup(t *tester.TesterContext) {
	t.Put("/v1/auth/group/user/1", authservice.GroupRequest{GroupName: "__test_group__"}, mock.Comment("add the user to group with name"))
	t.Equal(true, t.DecodeJSON(t.Get("/v1/auth/group/user/1?group=__test_group__", mock.AbortRecord(true)).Body(), new(authservice.CheckReply)).(*authservice.CheckReply).Has, "must be added")
}
