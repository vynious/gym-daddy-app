from flask import Flask, request, jsonify
import grpc
import requests
from os import environ

# Import the generated protobuf code.
from pb.booking import booking_service_pb2 as booking_pb2
from pb.booking import booking_service_pb2_grpc as booking_pb2_grpc
from pb.booking import booking_message_pb2 as booking_message_pb2
from pb.booking import (
    booking_message_pb2_grpc as booking_message_pb2_grpc,
)  # nothing inside

from pb.notification import notification_service_pb2 as notification_pb2
from pb.notification import notification_service_pb2_grpc as notification_pb2_grpc
from pb.notification import notification_message_pb2 as notification_message_pb2
from pb.notification import (
    notification_message_pb2_grpc as notification_message_pb2_grpc,
)

from pb.queue import queue_message_pb2 as queue_pb2
from pb.queue import queue_message_pb2_grpc as queue_message_pb2_grpc


app = Flask(__name__)

FLASK_CLASS_SERVER = environ.get("FLASK_CLASS_SERVER") or "http://localhost:5200"
# Assuming the gRPC server is running on localhost:6000
BOOKING_GRPC_SERVER = environ.get("BOOKING_GRPC_SERVER") or "localhost:6000"
# Create a gRPC channel and client stub for the booking server
channel = grpc.insecure_channel(BOOKING_GRPC_SERVER)
booking_grpc_client = booking_pb2_grpc.BookingServiceStub(channel)

NOTIFICATION_GRPC_SERVER = environ.get("NOTIFICATION_GRPC_SERVER") or "localhost:3000"
# Create a gRPC channel and client stub for the notification server
channel = grpc.insecure_channel(NOTIFICATION_GRPC_SERVER)
notification_grpc_client = notification_pb2_grpc.NotificationServiceStub(channel)


def authenticate(f):
    def wrapper(*args, **kwargs):
        token = request.headers.get('Authorisation')
        if not token:
            return jsonify({"error": "Missing Authorisation header"}), 401

        validate_jwt_url = "http://user-ms:3005/api/users/validatejwt/default"
        headers = {"Authorisation": token}
        response = requests.get(validate_jwt_url, headers=headers)

        if response.status_code != 200:
            return jsonify({"error": "Unauthorised"}), 401

        return f(*args, **kwargs)
    wrapper.__name__ = f.__name__
    return wrapper


@app.route("/api/booking", methods=["POST"])
@authenticate
def create_booking():
    print("start creating booking", flush=True)

    # Extract necessary info from the request
    user_id = request.json.get("user_id")
    class_id = request.json.get("class_id")
    res = requests.patch(
        f"http://{FLASK_CLASS_SERVER}/classes/{int(class_id)}", json={"action": "book"}
    )

    if res.status_code != 201:
        print("checking for error", flush=True)
        print(res.json(), flush=True)
        return jsonify({"message": "Failed to create booking", "error": str(res)}), 500

    # Create a new booking request message
    booking_request = booking_message_pb2.CreateBookingRequest(
        user_id=user_id, class_id=class_id
    )

    print("completed booking", flush=True)

    # Make a gRPC call to the server to create a booking
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.CreateBooking(booking_request)
        print(f"Received response from booking gRPC server: {response}")

    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({"error": str(e)}), 500

    # Create a new notification request message
    try:
        # Create a new notification request message

        notification_request = notification_pb2.CreateNotificationRequest(
            notification_type="Booking-Confirmation", user_id=user_id
        )

        notification_response = notification_grpc_client.CreateNotification(
            notification_request
        )

        print(
            f"Received response from notification gRPC server: {notification_response}",
            flush=True,
        )
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}", flush=True)
        return jsonify({"error": str(e)}), 500

    return (
        jsonify(
            {
                "booking_id": response.booking.id,
                "message": "Booking created and notification sent successfully",
            }
        ),
        201,
    )


@app.route("/api/booking/<string:booking_id>", methods=["DELETE"])
@authenticate
def cancel_booking(booking_id):
    booking_res, err = get_booking(booking_id=booking_id)
    if err == 500:
        return jsonify({"message": "Failed to cancel booking", "error": str(err)}), 500

    booking_res = booking_res.get_json()
    class_id = booking_res["booking"]["class_id"]

    res = requests.patch(
        f"http://{FLASK_CLASS_SERVER}/classes/{int(class_id)}",
        json={"action": "cancel"},
    )

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
        return jsonify({"error": str(e)}), 500

    return jsonify({"message": "Booking cancelled successfully"})


@app.route("/api/booking/<string:booking_id>", methods=["GET"])
@authenticate
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
        return jsonify({"error": str(e)}), 500
    class_response = requests.get(
        f"{FLASK_CLASS_SERVER}/classes/{int(response.booking.class_id)}"
    )
    if class_response.status_code != 200:
        return (
            jsonify({"message": "Failed to get class", "error": str(class_response)}),
            500,
        )

    booking_info = {
        "booking": {
            "id": response.booking.id,
            "user_id": response.booking.user_id,
            "class_id": response.booking.class_id,
            "created_at": response.booking.created_at.ToJsonString(),
            "class": class_response.json()["data"],
        }
    }

    return jsonify(booking_info), 200


@app.route("/api/booking", methods=["GET"])
@authenticate
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
        return jsonify({"error": str(e)}), 500

    bookings = []
    for booking in response.bookings:
        current_class = requests.get(
            f"http://{FLASK_CLASS_SERVER}/classes/{int(booking.class_id)}"
        )
        bookings.append(
            {
                "id": booking.id,
                "user_id": booking.user_id,
                "class_id": booking.class_id,
                "created_at": booking.created_at.ToJsonString(),
                "class": current_class.json()["data"],
            }
        )

    return jsonify({"bookings": bookings})


@app.route("/api/booking/user/<string:user_id>", methods=["GET"])
@authenticate
def get_booking_by_user(user_id):
    # Create a new get booking by user request message
    get_booking_by_user_request = booking_message_pb2.GetBookingByUserRequest(
        user_id=user_id
    )

    # Make a gRPC call to the server to get bookings by user
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.GetBookingByUser(get_booking_by_user_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({"error": str(e)}), 500

    bookings = []
    for booking in response.bookings:
        current_class = requests.get(
            f"http://{FLASK_CLASS_SERVER}/classes/{int(booking.class_id)}"
        )
        bookings.append(
            {
                "id": booking.id,
                "user_id": booking.user_id,
                "class_id": booking.class_id,
                "created_at": booking.created_at.ToJsonString(),
                "class": current_class.json()["data"],
            }
        )

    return jsonify({"bookings": bookings})


@app.route("/api/booking/<string:booking_id>", methods=["PUT"])
@authenticate
def update_booking(booking_id):
    # Extract necessary info from the request
    class_id = request.json.get("class_id")

    # Create a new update booking request message
    update_booking_request = booking_message_pb2.UpdateBookingRequest(
        id=booking_id, class_id=class_id
    )

    # Make a gRPC call to the server to update a booking
    try:
        print("Sending request to gRPC server...")
        response = booking_grpc_client.UpdateBooking(update_booking_request)
        print(f"Received response from gRPC server: {response}")
    except grpc.RpcError as e:
        print(f"gRPC call failed: {e}")
        return jsonify({"error": str(e)}), 500
    class_response = requests.get(
        f"http://{FLASK_CLASS_SERVER}/classes/{int(response.booking.class_id)}"
    )
    return jsonify(
        {
            "booking": {
                "id": response.booking.id,
                "user_id": response.booking.user_id,
                "class_id": response.booking.class_id,
                "created_at": response.booking.created_at.ToJsonString(),
                "class": class_response.json()["data"],
            }
        }
    )


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5002, debug=True)
