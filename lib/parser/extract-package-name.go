package parser

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func ParsePackageName(filePath string) (string, error) {
	fset := token.NewFileSet()
	if s, err := os.Stat(filePath); err != nil {
		return "", err
	} else if s.IsDir() {
		fs, err := parser.ParseDir(fset, filePath, nil, parser.PackageClauseOnly)
		if err != nil {
			return "", err
		}
		if len(fs) != 1 {
			return "", errors.New("ambiguous/bad package name")
		}
		var f *ast.Package
		for _, ft := range fs {
			f = ft
		}
		if f == nil {
			return "", errors.New("nil package")
		}
		return f.Name, nil
	} else {
		f, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
		if err != nil {
			return "", err
		}
		return f.Name.Name, nil
	}
}
