package user

import (
	"encoding/json"
	"errors"
	"github.com/Myriad-Dreamin/market/server"
	userservice "github.com/Myriad-Dreamin/market/service/user"
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

type ErrorObject struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

func (t *TesterContext) NoErr(resp server.ResponseI) bool {
	t.t.Helper()
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
	m.Run()
}

func TestRegister(t *testing.T) {
	a := tester.Context(t)
	a.Equal(0, 0, "a")
	resp := tester.Post("/v1/user", userservice.RegisterRequest{
		Name:         "123",
		Password:     "123",
		NickName:     "123",
		Phone:        "123",
		RegisterCity: "123",
	})
	a.NoErr(resp)

}
