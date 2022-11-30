package main

import (
	"context"
	"log"
	"time"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func doLongGreet(grpcClient pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked.")

	reqs := []*pb.GreetRequest{
		{Name: "Test"},
		{Name: "John"},
		{Name: "Doe"},
	}

	stream, err := grpcClient.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while closing LongGreet %v\n", err)
	}

	log.Printf("LongGreet %v\n", res.Result)
}
