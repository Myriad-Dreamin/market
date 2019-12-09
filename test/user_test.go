package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)



func TestUser(t *testing.T) {
	_ = t.Run("RegisterLogin", srv.HandleTestWithOutError(testUserRegisterLogin)) &&
		t.Run("Get", testUserGet) &&
		t.Run("Inspect", srv.HandleTestWithOutError(testUserInspect)) &&
		t.Run("List", testUserList) &&
		t.Run("Buy", srv.HandleTestWithOutError(testUserBuy)) &&
		t.Run("Sell", srv.HandleTestWithOutError(testUserSell)) &&
		t.Run("ConfirmBuy", srv.HandleTestWithOutError(testUserConfirmBuy)) &&
		t.Run("ConfirmSell", srv.HandleTestWithOutError(testUserConfirmSell)) &&
		t.Run("ChangePassword", srv.HandleTestWithOutError(testUserChangePassword)) &&
		t.Run("Put", srv.HandleTestWithOutError(testUserPut)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testUserDelete))
}

func testUserInspect(t *tester.TesterContext) {

}

func testUserPut(t *tester.TesterContext) {

}
