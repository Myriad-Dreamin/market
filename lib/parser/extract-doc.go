package parser

import (
	"errors"
	"fmt"
	"github.com/Myriad-Dreamin/market/docs"
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

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst
	return pkg
}

func GetPackageDir(fileDir string) *ast.Package {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, fileDir, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if len(pkgs) != 1 {
		log.Fatal(errors.New("invalid package"))
		return nil
	}
	var pkg *ast.Package
	for _, p := range pkgs {
		pkg = p
	}
	return pkg
}

func GetDoc(fileName string) *doc.Package {
	importPath, _ := filepath.Abs("/")
	return doc.New(GetPackage(fileName), importPath, doc.AllDecls)
}

func GetDocDir(pkgDir string) *doc.Package {
	importPath, _ := filepath.Abs("/")
	return doc.New(GetPackageDir(pkgDir), importPath, doc.AllDecls)
}

func makeDoc(list []*ast.Comment) string{
	var d = make([]string, len(list))
	for i := range list {
		text := list[i].Text
		d[i] = strings.TrimPrefix(text, "//")
		if strings.HasPrefix(text, "/*") {
			d[i] = strings.TrimSuffix(strings.TrimPrefix(text, "/*"), "*/")
		}
	}
	return strings.Join(d, "\n")
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
								return makeDoc(m.Doc.List)
							} else if m.Comment != nil {
								return makeDoc(m.Comment.List)
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

func InterfaceDescription(i interface{}) string {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Ptr:
		return TypeInterfaceDescription(t.Elem())
	default:
		panic("not pointer type")
	}
}
func TypeInterfaceDescription(t reflect.Type) string {
	//fmt.Println(t)
	switch t.Kind() {
	case reflect.Ptr:
		return TypeInterfaceDescription(t.Elem())
	case reflect.Interface, reflect.Struct:
		typeName := t.Name()
		path := docs.Get(t.PkgPath())
		if len(path) == 0 {
			panic(fmt.Errorf("not registered package: %v", t.PkgPath()))
		}
		myDoc := GetDocDir(path)
		for _, obj := range myDoc.Types {
			if obj.Name == typeName {
				return obj.Doc
			}
		}
		return ""
	default:
		return ""
	}
}
