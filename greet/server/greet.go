package main

import (
	"context"
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func (grpcServer *Server) Greet(c context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello, " + in.Name,
	}, nil
}
