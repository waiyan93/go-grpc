package main

import (
	"context"
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(grpcClient pb.CalculatorServiceClient, n int64) {
	log.Println("doSqrt was invoked.")

	res, err := grpcClient.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("Provided with negative number.")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt %f\n", res.Result)
}
