package main

import (
	"context"
	"io"
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func doPrime(grpcClient pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked.")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := grpcClient.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while invoking doPrime %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving %v\n", err)
		}

		log.Printf("Prime %d\n", msg.Result)
	}
}
