// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: notification/notification_service.proto

package notification

import (
	queue "/proto_files/queue"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicket         *queue.Ticket `protobuf:"bytes,1,opt,name=user_ticket,json=userTicket,proto3" json:"user_ticket,omitempty"`
	CurrentQueueNumber *int64        `protobuf:"varint,2,opt,name=current_queue_number,json=currentQueueNumber,proto3,oneof" json:"current_queue_number,omitempty"`
	NotificationType   string        `protobuf:"bytes,3,opt,name=notification_type,json=notificationType,proto3" json:"notification_type,omitempty"`
}

func (x *CreateNotificationRequest) Reset() {
	*x = CreateNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationRequest) ProtoMessage() {}

func (x *CreateNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNotificationRequest) GetUserTicket() *queue.Ticket {
	if x != nil {
		return x.UserTicket
	}
	return nil
}

func (x *CreateNotificationRequest) GetCurrentQueueNumber() int64 {
	if x != nil && x.CurrentQueueNumber != nil {
		return *x.CurrentQueueNumber
	}
	return 0
}

func (x *CreateNotificationRequest) GetNotificationType() string {
	if x != nil {
		return x.NotificationType
	}
	return ""
}

type CreateNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *CreateNotificationResponse) Reset() {
	*x = CreateNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationResponse) ProtoMessage() {}

func (x *CreateNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationResponse.ProtoReflect.Descriptor instead.
func (*CreateNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNotificationResponse) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

var File_notification_notification_service_proto protoreflect.FileDescriptor

var file_notification_notification_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x27, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x14, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x2b, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x17, 0x0a,
	0x15, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x80, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_notification_service_proto_rawDescOnce sync.Once
	file_notification_notification_service_proto_rawDescData = file_notification_notification_service_proto_rawDesc
)

func file_notification_notification_service_proto_rawDescGZIP() []byte {
	file_notification_notification_service_proto_rawDescOnce.Do(func() {
		file_notification_notification_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_notification_service_proto_rawDescData)
	})
	return file_notification_notification_service_proto_rawDescData
}

var file_notification_notification_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notification_notification_service_proto_goTypes = []interface{}{
	(*CreateNotificationRequest)(nil),  // 0: notification.CreateNotificationRequest
	(*CreateNotificationResponse)(nil), // 1: notification.CreateNotificationResponse
	(*queue.Ticket)(nil),               // 2: queue.Ticket
	(*Notification)(nil),               // 3: notification.Notification
}
var file_notification_notification_service_proto_depIdxs = []int32{
	2, // 0: notification.CreateNotificationRequest.user_ticket:type_name -> queue.Ticket
	3, // 1: notification.CreateNotificationResponse.notification:type_name -> notification.Notification
	0, // 2: notification.NotificationService.CreateNotification:input_type -> notification.CreateNotificationRequest
	1, // 3: notification.NotificationService.CreateNotification:output_type -> notification.CreateNotificationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_notification_service_proto_init() }
func file_notification_notification_service_proto_init() {
	if File_notification_notification_service_proto != nil {
		return
	}
	file_notification_notification_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notification_notification_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_notification_notification_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_notification_notification_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notification_notification_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_notification_service_proto_goTypes,
		DependencyIndexes: file_notification_notification_service_proto_depIdxs,
		MessageInfos:      file_notification_notification_service_proto_msgTypes,
	}.Build()
	File_notification_notification_service_proto = out.File
	file_notification_notification_service_proto_rawDesc = nil
	file_notification_notification_service_proto_goTypes = nil
	file_notification_notification_service_proto_depIdxs = nil
}
