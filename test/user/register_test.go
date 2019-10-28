package user

import (
	"encoding/json"
	"errors"
	"fmt"
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

var tester = StartTester()

func TestMain(m *testing.M) {
	defer tester.Release()
	resp := tester.Post("/v1/user", userservice.RegisterRequest{
		Name:         "123",
		Password:     "123",
		NickName:     "123",
		Phone:        "123",
		RegisterCity: "123",
	})
	if !tester.NoErr(resp) {
		return
	}

	var r userservice.RegisterReply
	err := resp.JSON(&r)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp = tester.Post("/v1/login",
		userservice.LoginRequest{
		ID:         r.ID,
		Password:     "123",
	})
	if !tester.NoErr(resp) {
		return
	}

	var r2 userservice.LoginReply
	err = resp.JSON(&r2)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(r2)

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

