package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mohammadinasab-dev/grpctask/grpcserver"
	"mohammadinasab-dev/grpctask/logwrapper"
	"mohammadinasab-dev/grpctask/protos"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
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
		runGrpcClient()
	}

}

func runGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := protos.NewProductServicesClient(conn)
	var Pid int
	var currency string
	fmt.Println("Enter product id")
	fmt.Scanln(&Pid)
	fmt.Println("Enter your currency")
	fmt.Scanln(&currency)
	product, err := client.GetProduct(context.Background(), &protos.ProductRequest{Id: int32(Pid), Currency: currency})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(product)

}
