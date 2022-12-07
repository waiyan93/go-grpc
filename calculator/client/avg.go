package main

import (
	"context"
	"log"
	"time"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func doAvg(grpcClient pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked.")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := grpcClient.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling stream %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error when receiving response %v\n", err)
	}

	log.Printf("Avg is %f\n", res.Result)
}
