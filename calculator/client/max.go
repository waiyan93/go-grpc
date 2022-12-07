package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func doMax(grpcClient pb.CalculatorServiceClient) {
	log.Println("doMax was invoked.")

	stream, err := grpcClient.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while invoking doMax %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 10},
	}

	waitChannel := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request with %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			req, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving response %v\n", err)
				break
			}

			log.Printf("Max is %d\n", req.Result)
		}
		close(waitChannel)
	}()

	<-waitChannel

}
