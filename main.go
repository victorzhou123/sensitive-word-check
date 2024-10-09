package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/victorzhou123/sensitive-word-check/proto"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// sensitive word check server
	swSvc, err := NewServer()
	if err != nil {
		log.Fatalf("new grpc server failed, err: %v", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterCheckServer(s, swSvc)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
