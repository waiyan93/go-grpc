package main

import (
	"context"
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func doGreet(grpcClient pb.GreetServiceClient) {
	log.Println("Do greet was invoked.")
	res, err := grpcClient.Greet(context.Background(), &pb.GreetRequest{
		Name: "Wai Yan",
	})

	if err != nil {
		log.Printf("Could not send response %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
