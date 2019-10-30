package user

import (
	userservice "github.com/Myriad-Dreamin/market/service/user"
	tester "github.com/Myriad-Dreamin/market/test"
	"testing"
)

var srv = tester.StartTester()

func TestMain(m *testing.M) {
	srv.PrintRequest(true)
	srv.Main(m)
}

func TestRegisterLogin(t *testing.T) {
	a := srv.Context(t)
	var (
		name  = "chan tan"
		nick  = "tan chan"
		phone = "10086111"
		pswd  = "11122222"
	)
	resp := srv.Post("/v1/user", userservice.RegisterRequest{
		Name:         name,
		Password:     pswd,
		NickName:     nick,
		Phone:        phone,
		RegisterCity: "tan arch",
	})
	a.NoErr(resp)
	id := a.DecodeJSON(resp.Body(),
		new(userservice.RegisterReply)).(*userservice.RegisterReply).ID
	resp = srv.Post("/v1/login", userservice.LoginRequest{
		ID:       id,
		Password: pswd,
	})
	a.NoErr(resp)
	resp = srv.Post("/v1/login", userservice.LoginRequest{
		Name:     name,
		Password: pswd,
	})
	a.NoErr(resp)
	resp = srv.Post("/v1/login", userservice.LoginRequest{
		Phone:    phone,
		Password: pswd,
	})
	a.NoErr(resp)
}

func TestRegister2(t *testing.T) {

}
