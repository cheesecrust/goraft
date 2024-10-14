package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init_node(node *node, port *string, client_ports []string) {
	// initialize server
	port_int, err := strconv.Atoi(*port)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	node.id = port_int
	node.client_cnt = len(client_ports)
	node.status = Follower
	node.is_voted = false
	node.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
	node.heartbeat_channel = make(chan bool)

}
