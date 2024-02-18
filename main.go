package main

import (
	"github.com/joho/godotenv"
	"github.com/vynious/gym-daddy/db"
	"github.com/vynious/gym-daddy/kafka"
	"github.com/vynious/gym-daddy/pb/proto_files/notification"
	"github.com/vynious/gym-daddy/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env files")
	}

	// init database
	repository, err := db.SpawnRepository(db.LoadDatabaseConfig())
	if err != nil {
		log.Fatalf(err.Error())
	}

	// init kafka producer
	prod := kafka.SpawnKafkaProducer(kafka.LoadKafkaConfigurations())

	// embed database service and kafka producer into rpc client
	server, err := rpc.SpawnGrpcServer(repository, prod)
	if err != nil {
		log.Fatalf(err.Error())
	}

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	notification.RegisterNotificationServiceServer(s, server)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer server.CloseConnections()
}
