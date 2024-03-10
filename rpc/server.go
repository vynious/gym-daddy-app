// rpc/server.go

package rpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"gd-booking-ms/pb/proto_files/booking"
	"gd-booking-ms/db"
)

type BookingServer struct {
	repo *db.Repository
	// Embed the unimplemented server as recommended by the gRPC development team for forward compatibility
	booking.UnimplementedBookingServiceServer
}

func NewBookingServer(r *db.Repository) *BookingServer {
	return &BookingServer{repo: r}
}

// Implement the gRPC service methods using the Repository, such as CreateBooking, etc.

func (s *BookingServer) Start() error {
	lis, err := net.Listen("tcp", ":50051") // Port can be changed as needed
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	booking.RegisterBookingServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

// Example implementation
func (s *BookingServer) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest) (*booking.CreateBookingResponse, error) {
	// Logic to create booking using s.repo
	return &booking.CreateBookingResponse{}, nil
}

// And so on for the rest of the gRPC methods...
