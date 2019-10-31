package tests

import (
	"github.com/Myriad-Dreamin/market/test/tester"
	"testing"
)

var srv = tester.StartTester()

func TestMain(m *testing.M) {
	srv.PrintRequest(true)
	srv.CollectResults(true)
	srv.MainM(m)
}

