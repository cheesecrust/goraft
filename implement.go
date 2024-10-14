package main

import (
	"context"
	pb "example/proto"
	"log"
)

// SayHello implements helloworld.GreeterServer
func (node *node) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v, %v", in.Num1, in.Num2)
	return &pb.HelloReply{Res1: in.Num1 + in.Num2, Res2: in.Num1 - in.Num2}, nil
}

func (node *node) RequestVote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {
	log.Printf("Received: %v", in.Port)
	if node.term > int(in.Term) || node.is_voted {
		return &pb.VoteReply{Granted: false}, nil
	}
	node.mu.Lock()
	node.is_voted = true
	node.mu.Unlock()
	return &pb.VoteReply{Granted: true}, nil
}

func (node *node) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (*pb.HeartBeatReply, error) {
	log.Println("Received heartbeat")
	node.heartbeat_channel <- true
	return &pb.HeartBeatReply{Res: true}, nil
}
