package crypto

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPassword(t *testing.T) {
	pswd, err := NewPassword("1")
	SetCost(20)
	fmt.Println(len(pswd), pswd, string(pswd))
	if err != nil {
		t.Error(err, reflect.TypeOf(err))
		return
	}
	checked, err := CheckPassword("1", pswd)
	if err != nil {
		t.Error(err, reflect.TypeOf(err))
		return
	}
	if !checked {
		t.Error("bad check")
	}
}

func TestCheckPassword(t *testing.T) {
	pswd := []byte("$2a$10$3FLs1pG2SzUuk/JscSH0GOuL0o2AXXYBwe4zgrIFLOTwLNs.Ll55O")
	checked, err := CheckPassword("1", pswd)
	if err != nil {
		t.Error(err, reflect.TypeOf(err))
		return
	}
	if !checked {
		t.Error("bad check")
	}
}

func TestCheckPasswordNotPassed(t *testing.T) {
	pswd := []byte("$2a$10$3FLs1pG2SzUuk/JscSH0GOuL0o2AXXYBwe4zgrIFLOTwLNs.Ll55O")
	checked, err := CheckPassword("2", pswd)
	if err != nil {
		t.Error(err, reflect.TypeOf(err))
		return
	}
	if checked {
		t.Error("bad check")
	}
}
