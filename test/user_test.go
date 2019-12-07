package tests

import (
	"testing"
)



func TestUser(t *testing.T) {
	_ = t.Run("RegisterLogin", srv.HandleTestWithOutError(testUserRegisterLogin)) &&
		t.Run("Get", testUserGet) &&
		t.Run("List", testUserList) &&
		t.Run("Buy", srv.HandleTestWithOutError(testUserBuy)) &&
		t.Run("Sell", srv.HandleTestWithOutError(testUserSell)) &&
		t.Run("ConfirmBuy", srv.HandleTestWithOutError(testUserConfirmBuy)) &&
		t.Run("ConfirmSell", srv.HandleTestWithOutError(testUserConfirmSell)) &&
		t.Run("ChangePassword", srv.HandleTestWithOutError(testUserChangePassword)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testUserDelete))
}
