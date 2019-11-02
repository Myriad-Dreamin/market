package parser

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


// @docname Maximum Document
// @category Accounts
// @description 12345
type RawService struct {

}


// @titLE Get AccountsA
//          @Description get accounts
//   @Success 200 {array} model.Account
func (RawService) TestHandler(ctx *gin.Context) {

}

// b
type subService interface {

}

// this is RawServiceI
type RawServiceI interface {
	// a
	subService
	// @titLE Get AccountsB
	//          @Description get accounts
	//   @Success 200 {array} model.Account
	/* @b
	 * @a
	 */
	TestHandler(ctx *gin.Context)
}

var (
	testHandler1 = func() {
		testHandler2 := func() {

		}
		// bad
		fmt.Println(FuncDescription(testHandler2))
	}
)


func TestFuncDescription(t *testing.T) {
	var testHandler0 = func() {
		testHandler3 := func() {
			testHandler4 := func() {
			}
			// bad
			fmt.Println(FuncDescription(testHandler4))
		}
		// bad
		fmt.Println(FuncDescription(testHandler3))
		testHandler3()
	}

	// ok
	fmt.Println(FuncDescription(inlineFunction))
	// ok
	fmt.Println(FuncDescription(multilineFunction))
	r := RawService{}
	var ri RawServiceI = r
	// ok
	fmt.Println(FuncDescription(r.TestHandler))
	// ok
	fmt.Println(FuncDescription(ri.TestHandler))

	// bad
	fmt.Println(FuncDescription(testHandler0))

	// bad
	fmt.Println(FuncDescription(testHandler1))

	testHandler1()
	testHandler0()
}


func TestInterfaceDescription(t *testing.T) {
	fmt.Println(InterfaceDescription(&RawService{}))
	var ri RawServiceI = RawService{}
	fmt.Println(InterfaceDescription(&ri))
}