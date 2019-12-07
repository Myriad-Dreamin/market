package tests

import (
	"testing"
)



func TestUser(t *testing.T) {
	_ = t.Run("TestUserRegisterLogin", srv.HandleTestWithOutError(testUserRegisterLogin)) &&
		t.Run("TestUserGet", testUserGet) &&
		t.Run("TestUserList", testUserList) &&
		t.Run("TestUserBuy", srv.HandleTestWithOutError(testUserBuy)) &&
		t.Run("TestUserSell", srv.HandleTestWithOutError(testUserSell)) &&
		t.Run("TestUserChangePassword", srv.HandleTestWithOutError(testUserChangePassword)) &&
		t.Run("TestUserDelete", srv.HandleTestWithOutError(testUserDelete))
}
