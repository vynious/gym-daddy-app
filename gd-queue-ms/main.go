package main

import (
	"log"
	"net"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	"github.com/vynious/gd-queue-ms/db"
	"github.com/vynious/gd-queue-ms/pb/proto_files/queue"
	"github.com/vynious/gd-queue-ms/queue_mgmt"
	"github.com/vynious/gd-queue-ms/rpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env variable: %v", err)
	}

	rdb, err := db.SpawnRedisDB("gym-daddy")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	qService := queue_mgmt.SpawnQueueService(rdb)
	rpcServer := rpc.SpawnServer(qService)

	grpcLis, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	queue.RegisterQueueServiceServer(grpcServer, rpcServer)

	// Start gRPC server in a goroutine
	go func() {
		log.Printf("gRPC server listening at %v", grpcLis.Addr())
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	metricsPort := ":9100"
	log.Printf("Prometheus metrics server listening at %v", metricsPort)
	if err := http.ListenAndServe(metricsPort, nil); err != nil {
		log.Fatalf("failed to start Prometheus metrics server: %v", err)
	}
}
