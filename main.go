package main

import (
	pb "example/proto"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	Follower = iota
	Candidate
	Leader
)

type node struct {
	pb.UnimplementedGreeterServer //이 부분은 안하면 에러가 발생한다. protobuf generate시 생성됨
	id                            int
	client_cnt                    int
	status                        int
	is_voted                      bool
	election_timeout              time.Duration
	term                          int

	conns             [5]*grpc.ClientConn
	clients           [5]pb.GreeterClient
	heartbeat_channel chan bool
	mu                sync.Mutex
}

func main() {
	node := &node{}
	port := flag.String("port", "12345", "The server port")
	clients := flag.String("client", "12346", "The client port")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	go run_server(port, node)

	client_ports := strings.Split(*clients, ",")
	var err error

	// initialize node
	init_node(node, port, client_ports)

	// 모두 연결될 때까지 연결 시도
	for index, client_port := range client_ports {
		log.Print(index)
		node.conns[index], err = grpc.Dial(fmt.Sprintf("127.0.0.1:%s", client_port), grpc.WithInsecure(), grpc.WithBlock())
		node.clients[index] = pb.NewGreeterClient(node.conns[index]) //서버의 method를 사용할 수 있게 해줌
		defer node.conns[index].Close()
	}

	log.Print("Connected to all nodes")

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	for {
		switch node.status {
		case Follower:
			follower_behavior(node)
		case Candidate:
			candidate_behavior(node)
		case Leader:
			leader_behavior(node)
		}
	}
}
