package main

import (
	"io"
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
)

func (grpcServer *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked.")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while invoking GreetEveryone %v\n", err)
		}

		err = stream.Send(&pb.GreetResponse{
			Result: "Greet" + req.Name + "!.",
		})

		if err != nil {
			log.Fatalf("Error while sending response %v\n", err)
		}
	}
}
