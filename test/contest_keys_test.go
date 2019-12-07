package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"reflect"
)

var srv *tester.Tester

const (
	goodsEsIdK       = "goods/es/id"
	goodsEsDeleteIdK = "goods/es/id"
	goodsBuyIdK       = "goods/buy/id"
	needsEsIdK       = "needs/es/id"
	needsEsDeleteIdK = "needs/es/id"
	needsSellIdK = "needs/sell/id"
	normalUserIdKey  = "user/normal/key"
	normalUserPassword = "11112222"
	normalUserNewPassword = "111122222"
)

var intT = 1
var intType = reflect.TypeOf(intT)
var uintType = reflect.TypeOf(uint(1))



func RangeInt(l,r int) []int {
	var x = make([]int, r-l)
	for i := l; i < r; i++ {
		x[i-l] = i
	}
	return x
}

