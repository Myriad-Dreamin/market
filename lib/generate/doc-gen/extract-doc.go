package doc_gen

import (
	"fmt"
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

// Get description of a func
func FuncDescription(f interface{}) string {
	fpc := runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	fileName, _ := fpc.FileLine(0)
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
	fset := token.NewFileSet()

	fmt.Println(fileName, recvName, funcName, fpc.Name())

	// Parse src
	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	//_ = ast.Print(fset, parsedAst)

	//ast.Walk(&visitor{r: recvName, f:funcName}, parsedAst)

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst

	importPath, _ := filepath.Abs("/")
	myDoc := doc.New(pkg, importPath, doc.AllDecls)
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
