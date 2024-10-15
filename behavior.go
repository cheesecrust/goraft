package main

import (
	"context"
	pb "example/proto"
	"log"
	"time"
)

func follower_behavior(node *node) {
	log.Print("Follower")
	log.Println(node.election_timeout)

	for {
		select {
		// wait for clock time
		case <-time.After(node.election_timeout):
			change_status(node, Candidate)
			return
		case receive := <-node.heartbeat_channel:
			log.Printf("receive heartbeat: %v\n", receive)
			reset_timeout(node)
			continue
		}
	}
}

func candidate_behavior(node *node) {
	log.Println("Candidate")

	total_vote := 1
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < node.client_cnt; i++ {
		select {
		case <-node.heartbeat_channel:
			change_status(node, Follower)
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
		change_status(node, Leader)
		log.Println("candidate -> leader")
	} else {
		change_status(node, Follower)
		log.Println("candidate -> follower")
	}
}

func leader_behavior(node *node) {
	log.Println("Leader")
	node.term++
	for {
		for i := 0; i < node.client_cnt; i++ {
			log.Printf("send heartbeat %v\n", i)
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			node.clients[i].HeartBeat(ctx, &pb.HeartBeatRequest{Term: int32(node.term)})
		}
	}
}
