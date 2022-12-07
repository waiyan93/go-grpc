package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func doGreetEveryone(grpcClient pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked.")

	stream, err := grpcClient.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling stream %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{Name: "John"},
		{Name: "Doe"},
		{Name: "Aung"},
	}

	waitChannel := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving response %v\n", err)
				break
			}
			log.Printf("Received %v\n", res.Result)
		}
		close(waitChannel)
	}()

	<-waitChannel
}
