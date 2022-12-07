package main

import (
	"context"
	"log"
	"time"

	pb "github.com/waiyan93/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(grpcClient pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked.")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		Name: "Jone Doe",
	}
	res, err := grpcClient.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("This is not gRPC error: %v\n", err)
		}
	}

	log.Printf("Greet with deadeline %v\n", res.Result)
}
