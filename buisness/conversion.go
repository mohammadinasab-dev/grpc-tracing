package buisness

import (
	"mohammadinasab-dev/grpctask/data"
	"mohammadinasab-dev/grpctask/logwrapper"
)

type Buisness struct {
	STDLog  *logwrapper.STDLog
	Product *data.Product
}

func NewBuisness(STDLog *logwrapper.STDLog, product *data.Product) (*Buisness, error) { //check if error requierd here?
	return &Buisness{
		STDLog:  STDLog,
		Product: product,
	}, nil
}

//check lowercase method that has an access throw the instance

func getRate(base, dest string) float32 {

	//do or call sth to get the rate
	//clinet to other grpc microservice
	rate := 1.5
	return float32(rate)

}

func (bi Buisness) Convert(dest string) float32 {
	bi.STDLog.InfoLogger.Println("i'm here in the convert function")

	rate := getRate(bi.Product.Currency, dest)
	price := bi.Product.Price
	finalprice := price * rate
	//fmt.Println("final price is: ", finalprice)
	return finalprice
}
