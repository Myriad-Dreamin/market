package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

func TestUser(t *testing.T) {
	_ = t.Run("TestUserRegisterLogin", testUserRegisterLogin) &&
		t.Run("TestUserGet", testUserGet) &&
		t.Run("TestUserList", testUserList) &&
		t.Run("TestUserBuy", srv.HandleTestWithOutError(testUserBuy)) &&
		t.Run("TestUserSell", srv.HandleTestWithOutError(testUserSell)) &&
		t.Run("TestUserChangePassword", srv.HandleTestWithOutError(testUserChangePassword))
		t.Run("TestUserDelete", srv.HandleTestWithOutError(testUserDelete))
}

func testUserDelete(t *tester.TesterContext) {

}

func testUserChangePassword(t *tester.TesterContext) {

}

func testUserSell(t *tester.TesterContext) {

}

func testUserBuy(t *tester.TesterContext) {

}
