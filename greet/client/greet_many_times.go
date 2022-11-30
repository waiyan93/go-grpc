package main

import (
	"context"
	"io"
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func doGreetManyTimes(grpcClient pb.GreetServiceClient) {
	log.Println("GreetManyTimes was invoked.")

	req := &pb.GreetRequest{
		Name: "Wai Yan",
	}

	stream, err := grpcClient.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error when calling GreetManyTimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}

}
