package main

import (
	"context"
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
)

func (grpcServer *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.NumberOne + in.NumberTwo,
	}, nil
}
