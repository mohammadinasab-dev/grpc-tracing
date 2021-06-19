package grpcserver

import (
	"mohammadinasab-dev/grpctask/configuration"
	"mohammadinasab-dev/grpctask/data"
	"mohammadinasab-dev/grpctask/logwrapper"
	"mohammadinasab-dev/grpctask/protos"
	"net"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	openzipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	endpoint_url = "http://localhost:9411/api/v1/spans"
	host_url     = "127.0.0.1:8282"
	network      = "tcp"
	service_name = "product"
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

	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	grpclog.Println("starting server...")
	ls, err := net.Listen(network, host_url)
	if err != nil {
		STDLog.WarningLogger.Fatalf("failed to listen to thev port %v", err)
	}
	grpclog.Println("listenning established...")

	tracer, err := newServerTracer()
	if err != nil {
		panic(err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////
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

func newServerTracer() (opentracing.Tracer, error) {
	collector, err := openzipkin.NewHTTPCollector(endpoint_url)
	if err != nil {
		return nil, err
	}
	recorder := openzipkin.NewRecorder(collector, true, host_url, service_name)
	tracer, err := openzipkin.NewTracer(
		recorder,
		openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}
