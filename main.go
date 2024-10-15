package main

import (
	pb "example/proto"
	"flag"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const MAX_CONNECT = 100

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

	conns             [MAX_CONNECT]*grpc.ClientConn
	clients           [MAX_CONNECT]pb.GreeterClient
	heartbeat_channel chan bool
	mu                sync.Mutex
}

func main() {
	node := &node{}
	port := flag.String("port", "12345", "The server port")
	clients := flag.String("client", "12346", "The client port")
	flag.Parse()

	go run_server(port, node)

	client_ports := strings.Split(*clients, ",")

	// initialize node
	init_node(node, port, client_ports)
	defer func() {
		for _, conn := range node.conns {
			conn.Close()
		}
	}()

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
