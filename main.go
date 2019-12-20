//go:generate package-attach-to-path -generate_init
package main

import (
	"flag"
	"github.com/Myriad-Dreamin/market/server"
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/types"
	_ "net/http/pprof"
)

var (
	port    = flag.String("port", ":23336", "serve on port")
	isDebug = flag.Bool("debug", false, "serve with debug mode")
)

func main() {
	types.SetCityObjectMap(config.Cities)
	srv := server.New("./config")
	if srv == nil {
		return
	}

	// srv.Inject(myPlugins...)

	if *isDebug {
		srv.ServeWithPProf(*port)
	} else {
		srv.Serve(*port)
	}

}
