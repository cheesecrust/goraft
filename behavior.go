package main

import (
	"context"
	pb "example/proto"
	"log"
	"math/rand"
	"time"
)

func follower_behavior(node *node) {
	log.Print(Follower)
	log.Println(node.election_timeout)

	for {
		select {
		// wait for clock time
		case <-time.After(node.election_timeout):
			node.mu.Lock()
			node.status = Candidate
			node.is_voted = true
			node.mu.Unlock()
			return
		case receive := <-node.heartbeat_channel:
			log.Printf("receive heartbeat: %v\n", receive)
			node.mu.Lock()
			node.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
			node.mu.Unlock()
			continue
		}
	}
}

func candidate_behavior(node *node) {
	log.Println(Candidate)

	total_vote := 1
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < node.client_cnt; i++ {
		select {
		case <-node.heartbeat_channel:
			node.mu.Lock()
			node.status = Follower
			node.is_voted = false
			node.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
			node.mu.Unlock()
			return
		default:
			log.Println(i)
			reply, err := node.clients[i].RequestVote(ctx, &pb.VoteRequest{Port: int32(node.id), Term: int32(node.term)})
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
	if total_vote > (node.client_cnt+1)/2 {
		node.mu.Lock()
		node.status = Leader
		log.Println("candidate -> leader")
		node.mu.Unlock()
	} else {
		node.mu.Lock()
		node.status = Follower
		node.is_voted = false
		node.election_timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
		log.Println("candidate -> follower")
		node.mu.Unlock()
	}
}

func leader_behavior(node *node) {
	log.Println("leader")
	node.term++
	for {
		for i := 0; i < node.client_cnt; i++ {
			log.Printf("send heartbeat %v\n", i)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			node.clients[i].HeartBeat(ctx, &pb.HeartBeatRequest{Term: int32(node.term)})
			defer cancel()
		}
	}
}
