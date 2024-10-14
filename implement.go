package main

import (
	"context"
	pb "example/proto"
	"log"
)

// SayHello implements helloworld.GreeterServer
func (server *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v, %v", in.Num1, in.Num2)
	return &pb.HelloReply{Res1: in.Num1 + in.Num2, Res2: in.Num1 - in.Num2}, nil
}

func (server *server) RequestVote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {
	log.Printf("Received: %v", in.Port)
	if server.term > int(in.Term) || server.is_voted {
		return &pb.VoteReply{Granted: false}, nil
	}
	server.mu.Lock()
	server.is_voted = true
	server.mu.Unlock()
	return &pb.VoteReply{Granted: true}, nil
}

func (server *server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (*pb.HeartBeatReply, error) {
	log.Println("Received heartbeat")
	server.heartbeat_channel <- true
	return &pb.HeartBeatReply{Res: true}, nil
}
