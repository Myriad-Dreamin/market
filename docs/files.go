package docs

import (
	"path/filepath"
	"reflect"
	"runtime"
)

var RootPackage, DocsPackage string
var RootPath, DocsPath string
type __DOCS___ struct {}
var __DOCS__ __DOCS___

func init() {
	DocsPackage = reflect.TypeOf(__DOCS__).PkgPath()
	RootPackage = filepath.Dir(DocsPackage)
	_, DocsPath, _, _ = runtime.Caller(0)
	DocsPath = filepath.Dir(DocsPath)
	RootPath = filepath.Dir(DocsPath)
}