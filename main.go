package main

import (
	"flag"
	"log"
	"mohammadinasab-dev/grpctask/grpcclient"
	"mohammadinasab-dev/grpctask/grpcserver"
	"mohammadinasab-dev/grpctask/logwrapper"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//create an instances of log and pass it throw RunGrpcServer()
	STDLog, err := logwrapper.NewSTDLog()
	if err != nil {
		log.Fatalln("the logger dose'nt set")
	}

	op := flag.String("op", "s", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		grpcserver.RunGrpcServer(STDLog)
	case "c":
		grpcclient.RunGrpcClient()
	}

}
