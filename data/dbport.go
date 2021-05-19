package data

import "mohammadinasab-dev/grpctask/protos"

type DBProductCrudinterface interface {
	DBGetProduct(in *protos.ProductRequest) (*Product, error)
}
