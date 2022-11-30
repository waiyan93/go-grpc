package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func (grpcServer *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked.")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while leading client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s\n", req.Name)
	}
}
