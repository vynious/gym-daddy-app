// rpc/server.go

package rpc

import (
	"context"
	"github.com/syahmimscs/gd-booking-ms/db"
	"github.com/syahmimscs/gd-booking-ms/pb/proto_files/booking"
	"net"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *BookingServer) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest) (*booking.CreateBookingResponse, error) {
	bookingEntry, err := s.repo.CreateBooking(req.UserId, req.ClassId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating booking: %v", err)
	}

	return &booking.CreateBookingResponse{
		Booking: &booking.Booking{
			Id:        bookingEntry.ID,
			UserId:    bookingEntry.UserID,
			ClassId:   bookingEntry.ClassID,
			CreatedAt: timestamppb.New(bookingEntry.CreatedAt),
		},
	}, nil
}

func (s *BookingServer) ListBookings(ctx context.Context, req *booking.ListBookingsRequest) (*booking.ListBookingsResponse, error) {
	bookings, err := s.repo.ListBookings()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing bookings: %v", err)
	}

	bookingList := make([]*booking.Booking, 0, len(bookings))
	for _, b := range bookings {
		bookingList = append(bookingList, &booking.Booking{
			Id:        b.ID,
			UserId:    b.UserID,
			ClassId:   b.ClassID,
			CreatedAt: timestamppb.New(b.CreatedAt),
		})
	}

	return &booking.ListBookingsResponse{Bookings: bookingList}, nil
}

func (s *BookingServer) GetBookingByUser(ctx context.Context, req *booking.GetBookingByUserRequest) (*booking.GetBookingByUserResponse, error) {
	bookings, err := s.repo.GetBookingByUserId(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting bookings: %v", err)
	}

	bookingList := make([]*booking.Booking, 0, len(bookings))
	for _, b := range bookings {
		bookingList = append(bookingList, &booking.Booking{
			Id:        b.ID,
			UserId:    b.UserID,
			ClassId:   b.ClassID,
			CreatedAt: timestamppb.New(b.CreatedAt),
		})
	}
	return &booking.GetBookingByUserResponse{Bookings: bookingList}, nil
}

func (s *BookingServer) GetBooking(ctx context.Context, req *booking.GetBookingRequest) (*booking.GetBookingResponse, error) {
	bookingEntry, err := s.repo.GetBooking(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting booking: %v", err)
	}

	return &booking.GetBookingResponse{
		Booking: &booking.Booking{
			Id:        bookingEntry.ID,
			UserId:    bookingEntry.UserID,
			ClassId:   bookingEntry.ClassID,
			CreatedAt: timestamppb.New(bookingEntry.CreatedAt),
		},
	}, nil
}

func (s *BookingServer) UpdateBooking(ctx context.Context, req *booking.UpdateBookingRequest) (*booking.UpdateBookingResponse, error) {
	// Validate request
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Booking ID is required")
	}
	if req.ClassId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Class ID is required")
	}

	// Call the UpdateBooking method on the repository to update the booking entry in the database.
	updatedBooking, err := s.repo.UpdateBooking(req.Id, req.ClassId)
	if err != nil {
		// Handle any errors that occur during the update, e.g. not found, validation errors, etc.
		return nil, status.Errorf(codes.Internal, "Error updating booking: %v", err)
	}

	// Create and return an UpdateBookingResponse with the updated booking information.
	return &booking.UpdateBookingResponse{
		Booking: &booking.Booking{
			Id:        updatedBooking.ID,
			UserId:    updatedBooking.UserID,
			ClassId:   updatedBooking.ClassID,
			CreatedAt: timestamppb.New(updatedBooking.CreatedAt),
		},
	}, nil
}

func (s *BookingServer) CancelBooking(ctx context.Context, req *booking.CancelBookingRequest) (*booking.CancelBookingResponse, error) {
	err := s.repo.CancelBooking(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error cancelling booking: %v", err)
	}

	return &booking.CancelBookingResponse{}, nil
}
