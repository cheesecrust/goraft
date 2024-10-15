package main

import (
	pb "example/proto"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init_node(node *node, port *string, client_ports []string) {
	// initialize server
	port_int, err := strconv.Atoi(*port)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	node.id = port_int
	node.client_cnt = len(client_ports)
	node.heartbeat_channel = make(chan bool)
	change_status(node, Follower)

	for index, client_port := range client_ports {
		node.conns[index], _ = grpc.NewClient(fmt.Sprintf("127.0.0.1:%s", client_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		node.clients[index] = pb.NewGreeterClient(node.conns[index]) //서버의 method를 사용할 수 있게 해줌
	}
}

func change_status(node *node, status int) {
	println("status changed to ", status)
	node.status = status
	if status == Follower {
		node.mu.Lock()
		node.is_voted = false
		node.mu.Unlock()
		reset_timeout(node)
	} else if status == Candidate {
		node.mu.Lock()
		node.is_voted = true
		node.mu.Unlock()
	}
}

func reset_timeout(node *node) {
	node.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
}
