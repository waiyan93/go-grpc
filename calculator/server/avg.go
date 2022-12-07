package main

import (
	"io"
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func (grpcServer *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg was invoked!")
	count := int64(0)
	total := int64(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(total) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while receiving stream %v\n", err)
		}

		total += req.Number
		count++
	}
}
