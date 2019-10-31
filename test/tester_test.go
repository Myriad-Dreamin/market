package tests

import (
	doc_gen "github.com/Myriad-Dreamin/market/lib/generate/doc-gen"
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

var srv = tester.StartTester()

func TestMain(m *testing.M) {
	srv.PrintRequest(true)
	srv.CollectResults(true)
	srv.MainM(m)
	doc_gen.FromResults(srv.DumpResults())
}

