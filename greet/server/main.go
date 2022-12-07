package main

import (
	"log"
	"net"

	pb "github.com/waiyan93/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
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

	options := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates %v\n", err)
		}
		options = append(options, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(options...)

	pb.RegisterGreetServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
