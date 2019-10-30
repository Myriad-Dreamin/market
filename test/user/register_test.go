package user

import (
	"encoding/json"
	"errors"
	"github.com/Myriad-Dreamin/market/lib/tracer"
	"github.com/Myriad-Dreamin/market/server"
	userservice "github.com/Myriad-Dreamin/market/service/user"
	"log"
	"testing"
)

type Tester struct {
	*server.Mocker
}

type TesterContext struct {
	*server.MockerContext
	t *testing.T
}

func StartTester() (tester *Tester) {
	tester = new(Tester)
	tester.Mocker = server.Mock()
	if tester.Mocker == nil {
		panic(errors.New("req mocker error"))
	}
	return tester
}

func (t *Tester) Context(tt *testing.T) (s *TesterContext) {
	return &TesterContext{
		MockerContext: t.Mocker.Context(tt),
		t: tt,
	}
}

func (t *Tester) NoErr(resp server.ResponseI) bool {
	if resp.Code() != 200 {
		log.Fatal("resp has bad code ", resp.Code(), resp.Body().String())
		return false
	}
	body := resp.Body()
	var obj ErrorObject
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		log.Fatal(err)
		return false
	}
	if len(obj.Error) != 0 || obj.Code != 0 {
		log.Fatalf("Code, Error (%v, %v)", obj.Code, obj.Error)
	}
	return true
	//if gjson
}

type ErrorObject struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

func (t *TesterContext) NoErr(resp server.ResponseI) bool {
	t.t.Helper()
	if resp.Code() != 200 {
		t.t.Error("resp has bad code ", resp.Code())
		return false
	}
	body := resp.Body()
	var obj ErrorObject
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		t.t.Error(err)
		return false
	}
	if len(obj.Error) != 0 || obj.Code != 0 {
		t.t.Errorf("Code, Error (%v, %v)", obj.Code, obj.Error)
	}
	return true
	//if gjson
}


func (t *Tester) Release() {
	tester.Mocker.ReleaseMock()
}

func (t *Tester) MakeAdminContext() bool {
	resp := tester.Post("/v1/user", userservice.RegisterRequest{
		Name:         "admin_context",
		Password:     "admin",
		NickName:     "admin_context",
		Phone:        "1234567891011",
		RegisterCity: "Qing Dao S.D.",
	})
	if !tester.NoErr(resp) {
		return false
	}

	var r userservice.RegisterReply
	err := resp.JSON(&r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	resp = tester.Post("/v1/login",
		userservice.LoginRequest{
			ID:         r.ID,
			Password:     "admin",
		})
	if !tester.NoErr(resp) {
		return false
	}

	var r2 userservice.LoginReply
	err = resp.JSON(&r2)
	if err != nil {
		log.Fatal(err)
		return false
	}

	//fmt.Println(r2)
	//r2.RefreshToken
	tester.UseToken(r2.Token)
	return true
}

var tester = StartTester()

func TestMain(m *testing.M) {
	defer func() {
		if err := recover(); err != nil {
			tracer.PrintStack()
			tester.Logger.Error("panic", "error", err)
		}
		tester.Release()
		print("here")
	}()
	if !tester.MakeAdminContext() {
		return
	}
	m.Run()
}

func TestRegister(t *testing.T) {
	a := tester.Context(t)
	a.Equal(0, 0, "a")
	resp := tester.Post("/v1/user", userservice.RegisterRequest{
		Name:         "234",
		Password:     "234",
		NickName:     "234",
		Phone:        "234",
		RegisterCity: "234",
	})
	a.NoErr(resp)

}

func TestRegister2(t *testing.T) {

}

