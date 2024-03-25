# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: booking/booking_service.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from pb.booking import booking_message_pb2 as booking_dot_booking__message__pb2



DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1d\x62ooking/booking_service.proto\x12\x07\x62ooking\x1a\x1d\x62ooking/booking_message.proto2\xf9\x03\n\x0e\x42ookingService\x12P\n\rCreateBooking\x12\x1d.booking.CreateBookingRequest\x1a\x1e.booking.CreateBookingResponse\"\x00\x12M\n\x0cListBookings\x12\x1c.booking.ListBookingsRequest\x1a\x1d.booking.ListBookingsResponse\"\x00\x12G\n\nGetBooking\x12\x1a.booking.GetBookingRequest\x1a\x1b.booking.GetBookingResponse\"\x00\x12Y\n\x10GetBookingByUser\x12 .booking.GetBookingByUserRequest\x1a!.booking.GetBookingByUserResponse\"\x00\x12P\n\rUpdateBooking\x12\x1d.booking.UpdateBookingRequest\x1a\x1e.booking.UpdateBookingResponse\"\x00\x12P\n\rCancelBooking\x12\x1d.booking.CancelBookingRequest\x1a\x1e.booking.CancelBookingResponse\"\x00\x42\x16Z\x14/proto_files/bookingb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'booking.booking_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\024/proto_files/booking'
  _globals['_BOOKINGSERVICE']._serialized_start=74
  _globals['_BOOKINGSERVICE']._serialized_end=579
# @@protoc_insertion_point(module_scope)
