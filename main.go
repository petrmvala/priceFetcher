package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/petrmvala/priceFetcher/client"
	"github.com/petrmvala/priceFetcher/proto"
)

func main() {
	// client := client.New("http://localhost:3000")
	// price, err := client.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)

	// return
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the JSON transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the GRPC transport")
		svc      = NewLoggingService(NewMetricService(&priceFetcher{}))
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
