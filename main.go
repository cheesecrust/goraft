package main

import (
	pb "example/proto"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer //이 부분은 안하면 에러가 발생한다. protobuf generate시 생성됨
	id                            int
	client_cnt                    int
	status                        string
	is_voted                      bool
	election_timeout              time.Duration
	term                          int

	conns             [5]*grpc.ClientConn
	clients           [5]pb.GreeterClient
	heartbeat_channel chan bool
	mu                sync.Mutex
}

func main() {
	server := &server{}
	port := flag.String("port", "12345", "The server port")
	clients := flag.String("client", "12346", "The client port")
	flag.Parse()

	go run_server(port, server)

	client_ports := strings.Split(*clients, ",")
	var err error

	// initialize server
	port_int, err := strconv.Atoi(*port)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	server.id = port_int
	server.client_cnt = len(client_ports)
	server.status = "follower"
	server.is_voted = false
	server.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
	server.heartbeat_channel = make(chan bool)
	rand.Seed(time.Now().UnixNano())

	// 모두 연결될 때까지 연결 시도
	for index, client_port := range client_ports {
		log.Print(index)
		server.conns[index], err = grpc.Dial(fmt.Sprintf("127.0.0.1:%s", client_port), grpc.WithInsecure(), grpc.WithBlock())
		server.clients[index] = pb.NewGreeterClient(server.conns[index]) //서버의 method를 사용할 수 있게 해줌
		defer server.conns[index].Close()
	}

	log.Print("Connected to all servers")

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	for {
		switch server.status {
		case "follower":
			follower_behavior(server)
		case "candidate":
			candidate_behavior(server)
		case "leader":
			leader_behavior(server)
		}
	}
}
