# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: notification/notification_service.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from pb.notification import notification_message_pb2 as notification_dot_notification__message__pb2
from pb.queue import queue_message_pb2 as queue_dot_queue__message__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'notification/notification_service.proto\x12\x0cnotification\x1a\'notification/notification_message.proto\x1a\x19queue/queue_message.proto\x1a\x1cgoogle/api/annotations.proto\"\xbc\x01\n\x19\x43reateNotificationRequest\x12\'\n\x0buser_ticket\x18\x01 \x01(\x0b\x32\r.queue.TicketH\x00\x88\x01\x01\x12!\n\x14\x63urrent_queue_number\x18\x02 \x01(\x03H\x01\x88\x01\x01\x12\x0f\n\x07user_id\x18\x04 \x01(\t\x12\x19\n\x11notification_type\x18\x05 \x01(\tB\x0e\n\x0c_user_ticketB\x17\n\x15_current_queue_number\"N\n\x1a\x43reateNotificationResponse\x12\x30\n\x0cnotification\x18\x01 \x01(\x0b\x32\x1a.notification.Notification2\x9d\x01\n\x13NotificationService\x12\x85\x01\n\x12\x43reateNotification\x12\'.notification.CreateNotificationRequest\x1a(.notification.CreateNotificationResponse\"\x1c\x82\xd3\xe4\x93\x02\x16*\x11/api/notification:\x01*B\x1bZ\x19/proto_files/notificationb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'notification.notification_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\031/proto_files/notification'
  _globals['_NOTIFICATIONSERVICE'].methods_by_name['CreateNotification']._options = None
  _globals['_NOTIFICATIONSERVICE'].methods_by_name['CreateNotification']._serialized_options = b'\202\323\344\223\002\026*\021/api/notification:\001*'
  _globals['_CREATENOTIFICATIONREQUEST']._serialized_start=156
  _globals['_CREATENOTIFICATIONREQUEST']._serialized_end=344
  _globals['_CREATENOTIFICATIONRESPONSE']._serialized_start=346
  _globals['_CREATENOTIFICATIONRESPONSE']._serialized_end=424
  _globals['_NOTIFICATIONSERVICE']._serialized_start=427
  _globals['_NOTIFICATIONSERVICE']._serialized_end=584
# @@protoc_insertion_point(module_scope)
