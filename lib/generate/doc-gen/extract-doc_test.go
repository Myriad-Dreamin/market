package doc_gen

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

// inline
// inline2
func inlineFunction() {}

/*
lines1
lines2
*/
func multilineFunction() {}
type rawService struct {

}

// @titLE Get AccountsA
//          @Description get accounts
//   @Success 200 {array} model.Account
func (rawService) testHandler(ctx *gin.Context) {

}

type rawServiceI interface {
	// @titLE Get AccountsB
	//          @Description get accounts
	//   @Success 200 {array} model.Account
	testHandler(ctx *gin.Context)
}

var (
	testHandler1 = func() {
		testHandler2 := func() {

		}
		fmt.Println(FuncDescription(testHandler2))
	}
)


func TestFuncDescription(t *testing.T) {
	var testHandler0 = func() {
		testHandler3 := func() {
			testHandler4 := func() {
			}
			fmt.Println(FuncDescription(testHandler4))
		}
		fmt.Println(FuncDescription(testHandler3))
		testHandler3()
	}


	fmt.Println(FuncDescription(inlineFunction))
	fmt.Println(FuncDescription(multilineFunction))
	r := rawService{}
	var ri rawServiceI = r
	fmt.Println(FuncDescription(r.testHandler))
	fmt.Println(FuncDescription(ri.testHandler))
	fmt.Println(FuncDescription(testHandler0))
	fmt.Println(FuncDescription(testHandler1))

	testHandler1()
	testHandler0()
}