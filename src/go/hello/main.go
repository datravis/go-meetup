package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/datravis/go-meetup/src/go/protogen"
)

type HelloService struct{}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	log.Println("Starting Hello server")
	grpcServer := grpc.NewServer(opts...)
	protogen.RegisterHelloServiceServer(grpcServer, HelloService{})
	grpcServer.Serve(lis)
}
