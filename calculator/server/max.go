package main

import (
	"io"
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func (grpcServer *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked.")

	var max int64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving request %v\n", req)
		}

		log.Printf("The request is %d\n", req.Number)

		if number := req.Number; number > max {
			max = number
			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})
			if err != nil {
				log.Fatalf("Error while sending response %v\n", err)
			}
		}
	}
}
