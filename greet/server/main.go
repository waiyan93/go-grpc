package main

import (
	"log"
	"net"

	pb "github.com/waiyan93/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

var address string = "0.0.0.0:3000"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on%v\n", err)
	}

	log.Printf("Listening on: %s\n", address)

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &Server{})

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
