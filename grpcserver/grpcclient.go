package grpcserver

import (
	"context"
	"fmt"
	"log"
	"mohammadinasab-dev/grpctask/protos"

	"google.golang.org/grpc"
)

func runGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := protos.NewProductServicesClient(conn)
	var id int32
	var currency string
	fmt.Println("enter ProductID and desired currency")
	fmt.Scanf("%d %s", &id, &currency)

	user, err := client.GetProduct(context.Background(), &protos.ProductRequest{Id: id, Currency: currency})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)

}
