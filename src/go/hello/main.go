package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/datravis/go-meetup/src/go/protogen"
)

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, in *protogen.HelloRequest) (*protogen.HelloResponse, error) {
	fmt.Println(in.Message)
	return &protogen.HelloResponse{
		Message: "What's up",
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	log.Println("Starting Hello server")
	grpcServer := grpc.NewServer(opts...)
	protogen.RegisterHelloServiceServer(grpcServer, &HelloService{})
	grpcServer.Serve(lis)
}
