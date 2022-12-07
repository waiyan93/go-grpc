package main

import (
	"log"

	pb "github.com/waiyan93/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address string = "localhost:3000"

func main() {
	tls := true
	options := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error loading CA trust certificate: %v\n", err)
		}

		options = append(options, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(address, options...)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	grpcClient := pb.NewGreetServiceClient(conn)

	doGreet(grpcClient)
	// doGreetManyTimes(grpcClient)
	// doLongGreet(grpcClient)
	// doGreetEveryone(grpcClient)
	// doGreetWithDeadline(grpcClient, 1*time.Second)
}
