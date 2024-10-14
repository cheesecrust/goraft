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
	node.heartbeat_channel = make(chan bool)
	reset_timeout(node)
}

func change_status(node *node, status int) {
	node.mu.Lock()
	node.status = status
	if status == Follower {
		node.is_voted = false
		reset_timeout(node)
	} else if status == Candidate {
		node.is_voted = true
	}
	node.mu.Unlock()
}

func reset_timeout(node *node) {
	node.mu.Lock()
	node.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
	node.mu.Unlock()
}