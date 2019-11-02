package doc_gen

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/ast"
	"go/doc"
	"path/filepath"

	//"go/ast"
	//"go/doc"
	"go/parser"
	"go/token"
	"log"
	//"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

//type visitor struct {
//	r, f string
//	res interface{}
//}

//func (v *visitor) Visit(node ast.Node) (w visitor) {
//	switch n := node.(type) {
//	case *ast.Field:
//		//fmt.Println(n.)
//	case  *ast.FuncType:
//	}
//	return
//
//}

func GetPackage(fileName string) *ast.Package {
	fset := token.NewFileSet()

	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	//_ = ast.Print(fset, parsedAst)

	//ast.Walk(&visitor{r: recvName, f:funcName}, parsedAst)

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst
	return pkg
}

func GetDoc(fileName string) *doc.Package {
	importPath, _ := filepath.Abs("/")
	return doc.New(GetPackage(fileName), importPath, doc.AllDecls)
}

// Get description of a func
func FuncDescription(f interface{}) string {
	fpc := runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	prf := strings.Split(filepath.Base(fpc.Name()), ".")
	var recvName, funcName string
	if len(prf) == 3 {
		recvName = prf[len(prf)-2]
		funcName = prf[len(prf)-1]
	} else {
		funcName = prf[len(prf)-1]
	}
	for len(recvName) > 0 && recvName[0] == '*' {
		recvName = recvName[1:]
	}
	for len(funcName) >= 3 && funcName[len(funcName)-3:len(funcName)] == "-fm" {
		funcName = funcName[:len(funcName)-3]
	}

	filename, _ := fpc.FileLine(0)
	myDoc := GetDoc(filename)
	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return theFunc.Doc
		}
		if theFunc.Name == recvName {
			fmt.Println(theFunc.Name)
		}
	}

	for _, obj := range myDoc.Types {
		if obj.Name != recvName {
			continue
		}
		for _, m := range obj.Methods {
			if m.Recv == recvName && m.Name == funcName {
				return m.Doc
			}
		}
		for _, m := range obj.Decl.Specs {
			switch mm := m.(type) {
			case *ast.TypeSpec:
				switch obj := mm.Type.(type) {
				case *ast.InterfaceType:
					for _, m := range obj.Methods.List {
						if m.Names[0].Name == funcName {

							if m.Doc != nil {
								var docs = make([]string, len(m.Doc.List))
								for i := range m.Doc.List {
									docs[i] = m.Doc.List[i].Text
								}
								return strings.Join(docs, "\n")
							} else if m.Comment != nil {
								var docs = make([]string, len(m.Comment.List))
								for i := range m.Comment.List {
									docs[i] = m.Comment.List[i].Text
								}
								return strings.Join(docs, "\n")
							}
							return ""
						}
					}
				}
			}
		}
	}

	for _, obj := range myDoc.Vars {
		fmt.Println(obj.Names)
	}

	for _, obj := range myDoc.Consts {
		fmt.Println(obj.Names)
	}
	return ""
}

// @category Accounts
// @description 12345
type RawService struct {

}


// @titLE Get AccountsA
//          @Description get accounts
//   @Success 200 {array} model.Account
func (RawService) TestHandler(ctx *gin.Context) {

}

func InterfaceDescription(i interface{}) string {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Func:
		return FuncDescription(i)
	case reflect.Ptr:
		return FuncDescription(reflect.ValueOf(i).Elem())
	case reflect.Interface:
		fmt.Println(t.NumMethod())
		return ""
	case reflect.Struct:
		//v := reflect.ValueOf(i)
		//if v.NumMethod() > 0 {
		//
		//	//myDoc := GetDoc(fileName)
		//	fmt.Println(fileName)
		//}

		return ""
	default:
		fmt.Println(reflect.TypeOf(i).PkgPath())
		//pkg, _ := importer.Default().Import("log")
		fmt.Println(reflect.TypeOf(i))
		return ""
	}
}
