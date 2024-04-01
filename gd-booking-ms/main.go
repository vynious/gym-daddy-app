package main

import (
	"github.com/joho/godotenv"
	"github.com/syahmimscs/gd-booking-ms/db"                     // Assuming this is the correct path to your db package
	"github.com/syahmimscs/gd-booking-ms/pb/proto_files/booking" // Adjust the path to where the booking proto Go files are generated
	"github.com/syahmimscs/gd-booking-ms/rpc"                    // Correct the path according to your actual rpc package
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or error loading .env file: %v", err)
		// You can decide whether to continue or exit based on the need for .env variables
		// os.Exit(1)
	}

	// Initialize database repository
	repository, err := db.SpawnRepository(db.LoadDatabaseConfig())
	if err != nil {
		log.Fatalf("failed to initialize database repository: %v", err)
	}
	defer repository.CloseConnection() // Defer the clean-up to ensure it's called before exiting main

	// Initialize gRPC server
	server := rpc.NewBookingServer(repository)

	lis, err := net.Listen("tcp", ":6000") // Choose the appropriate port for your service
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


}
