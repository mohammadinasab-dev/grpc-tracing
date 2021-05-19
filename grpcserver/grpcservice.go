package grpcserver

import (
	"context"
	"mohammadinasab-dev/grpctask/buisness"
	"mohammadinasab-dev/grpctask/data"
	"mohammadinasab-dev/grpctask/logwrapper"
	"mohammadinasab-dev/grpctask/protos"
)

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
	//log every things here
	product, err := server.dbhandler.DBGetProduct(in)
	if err != nil {
		server.STDLog.ErrorLogger.Println(err)
		return nil, err
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
