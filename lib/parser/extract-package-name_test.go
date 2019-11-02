package parser

import (
	"fmt"
	"testing"
)

func TestParsePackageName(t *testing.T) {
	fmt.Println(ParsePackageName("./extract-doc.go"))
	fmt.Println(ParsePackageName("./"))
}