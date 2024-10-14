package main

import (
	pb "example/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func run_server(port *string, node *node) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, node) //client가 사용할 수 있도록 등록
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
