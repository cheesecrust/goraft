package main

import (
	"context"
	pb "example/proto"
	"log"
	"math/rand"
	"time"
)

func follower_behavior(server *server) {
	log.Print("follower")
	log.Println(server.election_timeout)

	for {
		log.Print("follower loop")
		select {
		// wait for clock time
		case <-time.After(server.election_timeout):
			server.mu.Lock()
			server.status = "candidate"
			server.is_voted = true
			server.mu.Unlock()
			return
		case receive := <-server.heartbeat_channel:
			log.Printf("receive heartbeat: %v\n", receive)
			server.mu.Lock()
			server.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
			server.mu.Unlock()
			continue
		}
	}
}

func candidate_behavior(server *server) {
	log.Println("candidate")

	total_vote := 0
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < server.client_cnt; i++ {
		reply, err := server.clients[i].RequestVote(ctx, &pb.VoteRequest{Port: int32(server.id)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if reply.Granted {
			total_vote++
		}
	}

	if total_vote > server.client_cnt/2 {
		server.mu.Lock()
		log.Println("candidate -> leader")
		server.status = "leader"
		server.mu.Unlock()
	} else {
		server.mu.Lock()
		server.status = "follower"
		log.Println("candidate -> follower")
		server.is_voted = false
		server.mu.Unlock()
	}
}

func leader_behavior(server *server) {
	log.Println("leader")

	for {
		for i := 0; i < server.client_cnt; i++ {
			log.Printf("send heartbeat %v\n", i)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			server.clients[i].HeartBeat(ctx, &pb.HeartBeatRequest{})
			defer cancel()
		}
	}
}
