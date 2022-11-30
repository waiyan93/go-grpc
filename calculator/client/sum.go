package main

import (
	"context"
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func doSum(grpcClient pb.CalculatorServiceClient) {
	log.Println("doSum was invoked.")
	res, err := grpcClient.Sum(context.Background(), &pb.SumRequest{
		NumberOne: 2,
		NumberTwo: 3,
	})

	if err != nil {
		log.Fatalf("Error while calling doSum %v\n", err)
	}

	log.Printf("Sum %d\n", res.Result)
}
