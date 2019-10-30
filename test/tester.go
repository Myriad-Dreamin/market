package tester

import (
	"encoding/json"
	"errors"
	"github.com/Myriad-Dreamin/market/lib/tracer"
	"github.com/Myriad-Dreamin/market/server"
	userservice "github.com/Myriad-Dreamin/market/service/user"
	"io"
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

func (t *TesterContext) DecodeJSON(body io.Reader, req interface{}) interface{} {
	if err := json.NewDecoder(body).Decode(req); err != nil {
		t.t.Fatal(err)
	}
	return req
}


func (t *Tester) Release() {
	t.Mocker.ReleaseMock()
}

func (t *Tester) MakeAdminContext() bool {
	resp := t.Post("/v1/user", userservice.RegisterRequest{
		Name:         "admin_context",
		Password:     "admin",
		NickName:     "admin_context",
		Phone:        "1234567891011",
		RegisterCity: "Qing Dao S.D.",
	})
	if !t.NoErr(resp) {
		return false
	}

	var r userservice.RegisterReply
	err := resp.JSON(&r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	resp = t.Post("/v1/login",
		userservice.LoginRequest{
			ID:         r.ID,
			Password:     "admin",
		})
	if !t.NoErr(resp) {
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
	t.UseToken(r2.Token)
	return true
}

func (t *Tester) Main(m *testing.M) {
	defer func() {
		if err := recover(); err != nil {
			tracer.PrintStack()
			t.Logger.Error("panic", "error", err)
		}
		t.Release()
		print("here")
	}()
	if !t.MakeAdminContext() {
		return
	}
	m.Run()
}
