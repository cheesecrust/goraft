package main

import (
	"context"
	pb "example/proto"
	"log"
	"math/rand"
	"time"
)

func follower_behavior(server *server) {
	log.Print(Follower)
	log.Println(server.election_timeout)

	for {
		select {
		// wait for clock time
		case <-time.After(server.election_timeout):
			server.mu.Lock()
			server.status = Candidate
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
	log.Println(Candidate)

	total_vote := 1
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < server.client_cnt; i++ {
		select {
		case <-server.heartbeat_channel:
			server.mu.Lock()
			server.status = Follower
			server.is_voted = false
			server.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
			server.mu.Unlock()
			return
		default:
			log.Println(i)
			reply, err := server.clients[i].RequestVote(ctx, &pb.VoteRequest{Port: int32(server.id), Term: int32(server.term)})
			if err != nil {
				log.Printf("Error in RequestVote to client %v: %v", i, err)
				continue
			}
			if reply.Granted {
				total_vote++
			}
		}
	}
	log.Printf("total_vote: %v\n", total_vote)
	if total_vote > (server.client_cnt+1)/2 {
		server.mu.Lock()
		server.status = Leader
		log.Println("candidate -> leader")
		server.mu.Unlock()
	} else {
		server.mu.Lock()
		server.status = Follower
		server.is_voted = false
		log.Println("candidate -> follower")
		server.mu.Unlock()
	}
}

func leader_behavior(server *server) {
	log.Println("leader")
	server.term++
	for {
		for i := 0; i < server.client_cnt; i++ {
			log.Printf("send heartbeat %v\n", i)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			server.clients[i].HeartBeat(ctx, &pb.HeartBeatRequest{Term: int32(server.term)})
			defer cancel()
		}
	}
}
