package tests

import (
	doc_gen "github.com/Myriad-Dreamin/market/lib/generate/doc-gen"
	"github.com/Myriad-Dreamin/market/lib/sugar"
	"github.com/Myriad-Dreamin/market/server"
	"github.com/Myriad-Dreamin/market/test/tester"
	"os"
	"testing"
)

var srv *tester.Tester

func TestMain(m *testing.M) {
	sugar.WithFile(func(logFile *os.File) {
		var options = []server.Option{
			server.OptionRouterLoggerWriter{
				Writer: logFile,
			},
		}
		srv = tester.StartTester(options)
		srv.PrintRequest(true)
		srv.CollectResults(true)
		srv.MainM(m)
		err := doc_gen.FromGinResults(&doc_gen.GinInfo{
			Result: srv.DumpResults(),
			Host: "127.0.0.1",
			ApiDoc: "Minimum-Market",
		})
		if err != nil {
			panic(err)
		}
	}, "test.log")
}