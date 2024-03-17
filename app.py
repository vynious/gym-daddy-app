from flask import Flask, request, jsonify
import grpc
import requests
# Import the generated protobuf code.
from pb.booking import booking_service_pb2 as booking_pb2
from pb.booking import booking_service_pb2_grpc as booking_pb2_grpc
from pb.booking import booking_message_pb2 as booking_message_pb2
from pb.booking import booking_message_pb2_grpc as booking_message_pb2_grpc # nothing inside

from pb.notification import notification_service_pb2 as notification_pb2
from pb.notification import notification_service_pb2_grpc as notification_pb2_grpc
from pb.notification import notification_message_pb2 as notification_message_pb2
from pb.notification import notification_message_pb2_grpc as notification_message_pb2_grpc




app = Flask(__name__)

FLASK_CLASS_SERVER = 'http://localhost:5200'
# Assuming the gRPC server is running on localhost:6000
GRPC_SERVER = 'localhost:6000'
# Create a gRPC channel and client stub for the booking server
channel = grpc.insecure_channel(GRPC_SERVER)
booking_grpc_client = booking_pb2_grpc.BookingServiceStub(channel)

NOTIFICATION_SERVER = 'http://localhost:3000'
# Create a gRPC channel and client stub for the notification server
channel = grpc.insecure_channel(NOTIFICATION_SERVER)
notification_grpc_client = notification_pb2_grpc.NotificationServiceStub(channel)



@app.route('/booking', methods=['POST'])
def create_booking():
    # Extract necessary info from the request
    user_id = request.json.get('user_id')
    class_id = request.json.get('class_id')
    res = requests.patch(f"{FLASK_CLASS_SERVER}/classes/{int(class_id)}", json={'action': 'book'})
    print(res)
    if res.status_code != 201:
        return jsonify({"message": "Failed to create booking", "error": str(res)}), 500
    

    # Create a new booking request message
    booking_request = booking_message_pb2.CreateBookingRequest(user_id=user_id, class_id=class_id)

    # Make a gRPC call to the server to create a booking
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.CreateBooking(booking_request)
        print(f"Received response from booking gRPC server: {response}")



    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500

    return jsonify({'booking_id': response.booking.id}), 201

@app.route('/booking/<booking_id>', methods=['DELETE'])
def cancel_booking(booking_id):
    booking_res = get_booking(booking_id=booking_id)
    # if err:
        # return jsonify({"message": "Failed to cancel booking", "error": str(err)}), 500
    booking_res = booking_res.get_json()
    class_id = booking_res['booking']['class_id']

    res = requests.patch(f"{FLASK_CLASS_SERVER}/classes/{int(class_id)}", json={'action': 'cancel'})

    if res.status_code != 201:
        return jsonify({"message": "Failed to cancel booking", "error": str(res)}), 500

    # Create a new cancel booking request message
    cancel_booking_request = booking_message_pb2.CancelBookingRequest(id=booking_id)

    # Make a gRPC call to the server to cancel a booking
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.CancelBooking(cancel_booking_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500

    return jsonify({'message': 'Booking cancelled successfully'})


@app.route('/booking/<booking_id>', methods=['GET'])
def get_booking(booking_id):
    # Create a new get booking request message
    get_booking_request = booking_message_pb2.GetBookingRequest(id=booking_id)

    # Make a gRPC call to the server to get a booking
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.GetBooking(get_booking_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500


    booking_info = {'booking': {
        'id': response.booking.id,
        'user_id': response.booking.user_id,
        'class_id': response.booking.class_id,
        'created_at': response.booking.created_at.ToJsonString()
    }}

    return jsonify(booking_info)

@app.route('/booking', methods=['GET'])
def list_bookings():

    # Create a new list bookings request message
    list_bookings_request = booking_message_pb2.ListBookingsRequest()

    # Make a gRPC call to the server to list bookings
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.ListBookings(list_bookings_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500

    bookings = []
    for booking in response.bookings:
        bookings.append({
            'id': booking.id,
            'user_id': booking.user_id,
            'class_id': booking.class_id,
            'created_at': booking.created_at.ToJsonString()
        })

    return jsonify({'bookings': bookings})

@app.route('/booking/user/<user_id>', methods=['GET'])
def get_booking_by_user(user_id):
    # Create a new get booking by user request message
    get_booking_by_user_request = booking_message_pb2.GetBookingByUserRequest(user_id=user_id)

    # Make a gRPC call to the server to get bookings by user
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.GetBookingByUser(get_booking_by_user_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500

    bookings = []
    for booking in response.bookings:
        bookings.append({
            'id': booking.id,
            'user_id': booking.user_id,
            'class_id': booking.class_id,
            'created_at': booking.created_at.ToJsonString()
        })

    return jsonify({'bookings': bookings})


@app.route('/booking/<booking_id>', methods=['PUT'])
def update_booking(booking_id):
    # Extract necessary info from the request
    class_id = request.json.get('class_id')

    # Create a new update booking request message
    update_booking_request = booking_message_pb2.UpdateBookingRequest(id=booking_id, class_id=class_id)

    # Make a gRPC call to the server to update a booking
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.UpdateBooking(update_booking_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({'error': str(e)}), 500

    return jsonify({'booking': {
        'id': response.booking.id,
        'user_id': response.booking.user_id,
        'class_id': response.booking.class_id,
        'created_at': response.booking.created_at.ToJsonString()
    }})




if __name__ == '__main__':
    app.run(debug=True)
