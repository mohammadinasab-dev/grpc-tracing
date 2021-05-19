package grpcserver

import (
	"mohammadinasab-dev/grpctask/configuration"
	"mohammadinasab-dev/grpctask/data"
	"mohammadinasab-dev/grpctask/logwrapper"
	"mohammadinasab-dev/grpctask/protos"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func RunGrpcServer(STDLog *logwrapper.STDLog) {

	config, err := configuration.LoadConfig(".")
	if err != nil { //check this if else statement out in a function
		STDLog.ErrorLogger.Fatal(err)
	}
	var db data.DBProductCrudinterface

	//if config.environment == "product" or "debug"
	if config.DBDriver == "mysql" {
		db, err = data.CreateMySQLDBConnection(config, STDLog)
		if err != nil {
			STDLog.ErrorLogger.Fatal(err)
		}
	}

	// if config.DBDriver == "other"{
	// 	db = data.CreateMyOtherBConnection(config)
	// }

	grpclog.Println("starting server...")
	ls, err := net.Listen("tcp", ":8282")
	if err != nil {
		STDLog.WarningLogger.Fatalf("failed to listen to thev port %v", err)
	}
	grpclog.Println("listenning established...")
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...) //change the order of lines

	productserver, err := NewGrpcServer(STDLog, db)
	if err != nil {
		STDLog.WarningLogger.Fatal(err)
	}
	protos.RegisterProductServicesServer(server, productserver)
	err = server.Serve(ls)
	if err != nil {
		STDLog.WarningLogger.Fatal(err)
	}
}
