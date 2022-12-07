package main

import (
	"log"

	pb "github.com/waiyan93/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:3000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	grpcClient := pb.NewCalculatorServiceClient(conn)

	// doSum(grpcClient)
	// doPrime(grpcClient)
	// doAvg(grpcClient)
	// doMax(grpcClient)
	doSqrt(grpcClient, -4)
}
