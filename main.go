package main

import (
	"github.com/joho/godotenv"
	"github.com/vynious/gd-queue-ms/db"
	"github.com/vynious/gd-queue-ms/pb/proto_files/queue"
	"github.com/vynious/gd-queue-ms/queue_mgmt"
	"github.com/vynious/gd-queue-ms/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env variable")
	}

	rdb, err := db.SpawnRedisDB("gym-daddy")
	if err != nil {
		log.Fatalf(err.Error())
	}
	qService := queue_mgmt.SpawnQueueService(rdb)
	rpcServer := rpc.SpawnServer(qService)

	lis, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	queue.RegisterQueueServiceServer(s, rpcServer)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
