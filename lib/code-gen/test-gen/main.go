package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/common/log"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)


const (
	defaultSourceValue = "__default__"
)
var (
	source = flag.String("source", defaultSourceValue,"Input Go source package")
	destination = flag.String("destination", defaultSourceValue,"Output test files")
	structName = flag.String("struct", "Service", "for generate specified struct's test")
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {

	fmt.Println("hello code gen", *source, *destination)
	if *source == defaultSourceValue {
		fmt.Println("must input source path")
		os.Exit(1)
	}
	fset := token.NewFileSet()
	res, err := parser.ParseDir(fset, *source, nil, 0)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	for packageName ,packageSet := range res {
		fmt.Println(packageName)
		fmt.Println(packageSet.Scope, packageSet.Imports)
		var funcs []*ast.FuncDecl
		var typeSpecs []*ast.TypeSpec
		for i := range packageSet.Files {
			file := packageSet.Files[i]
			for _, d := range file.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					funcs = append(funcs, fn)
				}
				if tyGenDecl, isGendecl := d.(*ast.GenDecl);
					isGendecl && tyGenDecl.Tok == token.TYPE {
					for idt := range tyGenDecl.Specs {
						if typeSpec, ok := tyGenDecl.Specs[idt].(*ast.TypeSpec); ok {
							typeSpecs = append(typeSpecs, typeSpec)
						}
					}
				}
			}
		}

		var maintype *ast.TypeSpec
		for _, tp := range typeSpecs {
			if tp.Name.Name == *structName {
				maintype = tp
			}
		}

		for _, fn := range funcs {
			fmt.Println(fn.Name)

			err = ast.Print(fset, fn)
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
		}

		fmt.Println(maintype)
		_ = ast.Print(fset, maintype)
	}
}


