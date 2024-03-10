package main

import (
	"github.com/joho/godotenv"
	"gd-booking-ms/db" // Assuming this is the correct path to your db package
	"gd-booking-ms/rpc" // Correct the path according to your actual rpc package
	"gd-booking-ms/pb/proto_files/booking" // Adjust the path to where the booking proto Go files are generated
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env files: %v", err)
	}

	// Initialize database repository
	repository, err := db.SpawnRepository(db.LoadDatabaseConfig())
	if err != nil {
		log.Fatalf("failed to initialize database repository: %v", err)
	}

	// Initialize gRPC server
	server := rpc.NewBookingServer(repository)

	lis, err := net.Listen("tcp", ":8000") // Choose the appropriate port for your service
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register your booking service with the gRPC server
	booking.RegisterBookingServiceServer(s, server)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// This is assumed to be a cleanup function you've defined that closes any open connections
	// like database connections, Kafka producers, etc. It needs to be called before the program exits
	// or when an unrecoverable error occurs.
	defer repository.CloseConnection()
}
