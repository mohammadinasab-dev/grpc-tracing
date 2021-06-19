package grpcserver

import (
	"context"
	"mohammadinasab-dev/grpctask/buisness"
	"mohammadinasab-dev/grpctask/data"
	"mohammadinasab-dev/grpctask/logwrapper"
	"mohammadinasab-dev/grpctask/protos"

	"github.com/mohammadinasab-dev/grpctask/data"
	"github.com/opentracing/opentracing-go"
)

const service_name_db_query_product = "db query product"

type GrpcServer struct {
	STDLog    *logwrapper.STDLog
	dbhandler data.DBProductCrudinterface
}

func NewGrpcServer(STDLog *logwrapper.STDLog, db data.DBProductCrudinterface) (*GrpcServer, error) { //"user:1234@/people"
	return &GrpcServer{
		STDLog:    STDLog,
		dbhandler: db,
	}, nil
}

func (server *GrpcServer) GetProduct(ctx context.Context, in *protos.ProductRequest) (*protos.Product, error) {
	var product *data.Product
	var err error
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer := opentracing.GlobalTracer(); tracer != nil {
			mysqlSpan := tracer.StartSpan(service_name_db_query_product, opentracing.ChildOf(pctx))
			defer mysqlSpan.Finish()
			product, err = server.dbhandler.DBGetProduct(in)
			if err != nil {
				server.STDLog.ErrorLogger.Println(err)
				return nil, err
			}
		}
	}

	//ConvertCurrency(in.currency, product.currency)(product)
	buisness, err := buisness.NewBuisness(server.STDLog, product)
	if err != nil {
		server.STDLog.ErrorLogger.Println(err)
		return nil, err
	}
	product.Price = buisness.Convert(in.Currency) //***********

	pproduct, err := convertToGrpcProduct(*product)
	return pproduct, nil
}

func convertToGrpcProduct(product data.Product) (*protos.Product, error) {
	return &protos.Product{
		PID:      int32(product.PID),
		Name:     product.Name,
		Currency: product.Currency,
		Price:    product.Price,
	}, nil
}
