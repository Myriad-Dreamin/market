package tests

import (
	"github.com/Myriad-Dreamin/market/server"
	"github.com/Myriad-Dreamin/market/test/tester"
	"log"
	"os"
	"testing"
)

var srv *tester.Tester

func TestMain(m *testing.M) {
	logFile, err := os.OpenFile("test.log", os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("open log file error, err:", err)
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			log.Print("close log file error, err:", err)
		}
	}()
	var options = []server.Option{
		server.OptionRouterLoggerWriter{
			Writer: logFile,
		},
	}
	srv = tester.StartTester(options)
	srv.PrintRequest(true)
	srv.CollectResults(true)
	srv.MainM(m)
}
