// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: message/service/v1/message_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type MessageErrorReason int32

const (
	MessageErrorReason_UNKNOWN_ERROR                   MessageErrorReason = 0
	MessageErrorReason_ACCESS_USER_MEDAL_FAILED        MessageErrorReason = 1
	MessageErrorReason_GET_MAILBOX_LAST_TIME_FAILED    MessageErrorReason = 2
	MessageErrorReason_GET_MESSAGE_NOTIFICATION_FAILED MessageErrorReason = 3
	MessageErrorReason_SET_MAILBOX_LAST_TIME_FAILED    MessageErrorReason = 4
	MessageErrorReason_REMOVE_MAILBOX_COMMENT_FAILED   MessageErrorReason = 5
)

// Enum value maps for MessageErrorReason.
var (
	MessageErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "ACCESS_USER_MEDAL_FAILED",
		2: "GET_MAILBOX_LAST_TIME_FAILED",
		3: "GET_MESSAGE_NOTIFICATION_FAILED",
		4: "SET_MAILBOX_LAST_TIME_FAILED",
		5: "REMOVE_MAILBOX_COMMENT_FAILED",
	}
	MessageErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":                   0,
		"ACCESS_USER_MEDAL_FAILED":        1,
		"GET_MAILBOX_LAST_TIME_FAILED":    2,
		"GET_MESSAGE_NOTIFICATION_FAILED": 3,
		"SET_MAILBOX_LAST_TIME_FAILED":    4,
		"REMOVE_MAILBOX_COMMENT_FAILED":   5,
	}
)

func (x MessageErrorReason) Enum() *MessageErrorReason {
	p := new(MessageErrorReason)
	*p = x
	return p
}

func (x MessageErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_message_service_v1_message_error_proto_enumTypes[0].Descriptor()
}

func (MessageErrorReason) Type() protoreflect.EnumType {
	return &file_message_service_v1_message_error_proto_enumTypes[0]
}

func (x MessageErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageErrorReason.Descriptor instead.
func (MessageErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_message_service_v1_message_error_proto_rawDescGZIP(), []int{0}
}

var File_message_service_v1_message_error_proto protoreflect.FileDescriptor

var file_message_service_v1_message_error_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd7, 0x01, 0x0a, 0x12, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4d, 0x45, 0x44, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x20, 0x0a, 0x1c, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x49, 0x4c, 0x42, 0x4f, 0x58,
	0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x45, 0x54, 0x5f,
	0x4d, 0x41, 0x49, 0x4c, 0x42, 0x4f, 0x58, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45,
	0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x4d, 0x41, 0x49, 0x4c, 0x42, 0x4f, 0x58, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa0,
	0x45, 0xf4, 0x03, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_service_v1_message_error_proto_rawDescOnce sync.Once
	file_message_service_v1_message_error_proto_rawDescData = file_message_service_v1_message_error_proto_rawDesc
)

func file_message_service_v1_message_error_proto_rawDescGZIP() []byte {
	file_message_service_v1_message_error_proto_rawDescOnce.Do(func() {
		file_message_service_v1_message_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_service_v1_message_error_proto_rawDescData)
	})
	return file_message_service_v1_message_error_proto_rawDescData
}

var file_message_service_v1_message_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_service_v1_message_error_proto_goTypes = []interface{}{
	(MessageErrorReason)(0), // 0: message.v1.MessageErrorReason
}
var file_message_service_v1_message_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_service_v1_message_error_proto_init() }
func file_message_service_v1_message_error_proto_init() {
	if File_message_service_v1_message_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_service_v1_message_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_service_v1_message_error_proto_goTypes,
		DependencyIndexes: file_message_service_v1_message_error_proto_depIdxs,
		EnumInfos:         file_message_service_v1_message_error_proto_enumTypes,
	}.Build()
	File_message_service_v1_message_error_proto = out.File
	file_message_service_v1_message_error_proto_rawDesc = nil
	file_message_service_v1_message_error_proto_goTypes = nil
	file_message_service_v1_message_error_proto_depIdxs = nil
}
