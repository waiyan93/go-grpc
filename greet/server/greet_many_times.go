package main

import (
	"fmt"
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func (grpcServer *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked with %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.Name, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
