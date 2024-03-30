from flask import Flask, request, jsonify
from flask_cors import CORS
from os import environ
import os, sys
import grpc
import requests
from invokes import invoke_http
from google.protobuf.json_format import MessageToJson
import json

from pb.notification import notification_service_pb2 as notification_pb2
from pb.notification import notification_service_pb2_grpc as notification_pb2_grpc
from pb.notification import notification_message_pb2 as notification_message_pb2
from pb.notification import (
    notification_message_pb2_grpc as notification_message_pb2_grpc,
)

app = Flask(__name__)
CORS(app)

FLASK_CLASS_SERVER = environ.get("FLASK_CLASS_SERVER") or "localhost:5200" #class-ms:5200
NOTIFICATION_GRPC_SERVER = environ.get("NOTIFICATION_GRPC_SERVER") or "localhost:3000" #notification-ms:3000
REST_USER_SERVER = environ.get("REST_USER_SERVER") or "localhost:3005" #user-ms:3005
channel = grpc.insecure_channel(NOTIFICATION_GRPC_SERVER)
notification_grpc_client = notification_pb2_grpc.NotificationServiceStub(channel)

def authenticate(f):
    def wrapper(*args, **kwargs):
        token = request.headers.get('Authorisation')
        if not token:
            return jsonify({"error": "Missing Authorisation header"}), 401

        validate_jwt_url = "http://user-ms:3005/api/users/validatejwt"
        headers = {"Authorisation": token}
        response = requests.get(validate_jwt_url, headers=headers)

        if response.status_code != 200:
            return jsonify({"error": "Unauthorised"}), 401

        return f(*args, **kwargs)
    wrapper.__name__ = f.__name__
    return wrapper


@app.route("/api/class", methods=["POST"])
@authenticate
def create_class():
    if request.is_json:
        try:
            new_class = request.get_json()
            print("\nCreate class request in JSON:", new_class)

            # Send class info to class microservice
            result = process_create_class(
                new_class
            )  # Invokes function to fulfil create class process
            return jsonify(result), result["code"]

        except Exception as e:
            # Unexpected error in code
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            ex_str = (
                str(e)
                + " at "
                + str(exc_type)
                + ": "
                + fname
                + ": line "
                + str(exc_tb.tb_lineno)
            )
            print(ex_str)

            return (
                jsonify(
                    {
                        "code": 500,
                        "message": "create_class.py internal error: " + ex_str,
                    }
                ),
                500,
            )

    # Not a JSON request
    return (
        jsonify(
            {"code": 400, "message": "Invalid JSON input: " + str(request.get_data())}
        ),
        400,
    )

def process_create_class(new_class):
    # Send the class info
    # Invoke the class microservice
    print("\n-----Invoking class microservice-----")
    create_class_result = invoke_http(
        f"http://{FLASK_CLASS_SERVER}/classes", method="POST", json=new_class
    )
    print(create_class_result)
    if create_class_result["code"] != 201:
        return create_class_result

        # Retrieve all user IDs from user microservice
    print("\n-----Retrieving all user IDs-----")
    try:
        users_result = requests.get(f"http://{REST_USER_SERVER}/api/users/allusers", headers={"Authorisation": request.headers.get("Authorisation")})
        users_result.raise_for_status()  # This will raise an exception for HTTP errors
        user_ids = [user["user_id"] for user in users_result.json()]
    except requests.exceptions.HTTPError as e:
        print(f"Failed to retrieve user IDs: {e}")
        return {
            "code": e.response.status_code,
            "message": f"Failed to retrieve user IDs: {e.response.content}",
        }
    except Exception as e:
        print(f"An error occurred: {e}")
        return {"code": 500, "message": "Failed to retrieve user IDs"}

    # Send notification to users
    print("\n-----Invoking notification microservice-----")
    try:
        notification_requests = [notification_pb2.CreateNotificationRequest(
            user_id=user_id,
            notification_type="New-Class"
        ) for user_id in user_ids]
        notification_responses = []
        for notification_request in notification_requests:
            notification_response = notification_grpc_client.CreateNotification(
                notification_request
            )
        # Convert the Protobuf message to a JSON string, then parse it into a dictionary
            response_json = MessageToJson(notification_response)
            response_dict = json.loads(response_json)
            notification_responses.append(response_dict)
        print("Received responses from notification gRPC server:", flush=True)
    except Exception as e:
        print(f"Notification failed: {e}")
        return {
            "code": 500,
            "message": "Notification failed: " + str(e)
        }

    return {
        "code": 201,
        "data": {
            "create_class_result": create_class_result,
            "notification_responses": notification_responses,
        },
    }


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001, debug=True)
