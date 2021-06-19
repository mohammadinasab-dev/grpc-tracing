package grpcserver

import (
	"context"
	"fmt"
	"log"
	"mohammadinasab-dev/grpctask/protos"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/mohammadinasab-dev/grpctask/protos"
	"github.com/opentracing/opentracing-go"

	openzipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	// zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	// openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

const service_name_call_get = "callGet"

func RunGrpcClient() {
	tracer, collector, err := newClientTracer()
	if err != nil {
		panic(err)
	}
	defer collector.Close()
	conn, err := grpc.Dial(host_url, grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := protos.NewProductServicesClient(conn)
	var id int32
	var currency string
	fmt.Println("enter ProductID and desired currency")
	fmt.Scanf("%d %s", &id, &currency)

	// product, err := client.GetProduct(context.Background(), &protos.ProductRequest{Id: id, Currency: currency})
	product, err := callGetProduct(client, &protos.ProductRequest{Id: id, Currency: currency})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(product)

}

func newClientTracer() (opentracing.Tracer, openzipkin.Collector, error) {
	collector, err := openzipkin.NewHTTPCollector(endpoint_url)
	if err != nil {
		return nil, nil, err
	}
	recorder := openzipkin.NewRecorder(collector, true, host_url, service_name)
	tracer, err := openzipkin.NewTracer(
		recorder,
		openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, collector, nil
}

func callGetProduct(client protos.ProductServicesClient, product *protos.ProductRequest) (*protos.Product, error) {
	span := opentracing.StartSpan(service_name_call_get)
	defer span.Finish()
	time.Sleep(5 * time.Millisecond)
	// Put root span in context so it will be used in our calls to the client.
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	//ctx := context.Background()
	return client.GetProduct(ctx, product)
}
