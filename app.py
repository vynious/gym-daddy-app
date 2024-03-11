from flask import Flask, request, jsonify
import grpc
# Import the generated protobuf code.
from pb.booking import booking_service_pb2 as booking_pb2
from pb.booking import booking_service_pb2_grpc as booking_pb2_grpc

app = Flask(__name__)

# Assuming the gRPC server is running on localhost:6000
GRPC_SERVER = 'localhost:6000'

@app.route('/create-booking', methods=['POST'])
def create_booking():
    # Extract necessary info from the request
    user_id = request.json.get('user_id')
    class_id = request.json.get('class_id')

    # Create a gRPC channel and client
    channel = grpc.insecure_channel(GRPC_SERVER)
    grpc_client = booking_pb2_grpc.BookingServiceStub(channel)

    # Create a new booking request message
    booking_request = booking_pb2.CreateBookingRequest(user_id=user_id, class_id=class_id)

    # Make a gRPC call to the server to create a booking
    try:
        print("Sending request to gRPC server...")
        response = grpc_client.CreateBooking(booking_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500

    return jsonify({'booking_id': response.booking.id})

if __name__ == '__main__':
    app.run(debug=True)
