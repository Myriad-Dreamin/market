package main

import (
	"flag"
	"github.com/Myriad-Dreamin/market/docs"
	"github.com/Myriad-Dreamin/market/lib/parser"
	"github.com/Myriad-Dreamin/market/lib/sugar"
	"os"
	"path/filepath"
	"text/template"
)

const (
	defaultSourceValue = "__default__"
)

var (
	source = flag.String("source", defaultSourceValue,"Input Go source package")
	destination = flag.String("destination", defaultSourceValue,"Output test files")
	structName = flag.String("i", "Service", "for generate specified struct's test")
	//docRoot = flag.Bool("doc_root", false, "is package the docs root")
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	currentPath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}

	TargetPackageName, err := parser.ParsePackageName(currentPath)
	if err != nil {
		panic(err)
	}

	TargetPackage, err := filepath.Rel(docs.RootPath, currentPath)
	if err != nil {
		panic(err)
	}
	TargetPackage = filepath.Join(docs.RootPackage, TargetPackage)

	sugar.WithWriteFile(func(outputFile *os.File) {
		err = resultTemplate.Execute(outputFile, struct {
			PackageName, Package, DocsPackage, CurrentPath string
		}{
			PackageName:TargetPackageName,
			Package:TargetPackage,
			DocsPackage:docs.DocsPackage,
			CurrentPath:currentPath,
		})
	}, currentPath+"/register-docs.go")
}


var resultTemplate = template.New("resultTemplate")

func init() {
	var err error
	resultTemplate, err = resultTemplate.Parse(`// Code generated minimum-attach-file.go DO NOT EDIT
package {{.PackageName}}

import docs "{{.DocsPackage}}"

func init() {
	docs.Register("{{.Package}}", "{{.CurrentPath}}")
}

`)
	if err != nil {
		panic(err)
	}
}

