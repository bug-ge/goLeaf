// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: leaf-grpc/v1/error_reason.proto

package v1

import (
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

type ErrorReason int32

const (
	ErrorReason_SEG_SERVER_UNSPECIFIED   ErrorReason = 0
	ErrorReason_BIZ_TAG_NOT_FOUND        ErrorReason = 1
	ErrorReason_DB_OPERATE               ErrorReason = 2
	ErrorReason_ID_CACHE_INIT_FALSE      ErrorReason = 3
	ErrorReason_ID_TWO_SEGMENTS_ARE_NULL ErrorReason = 4
	// 时钟回拨超过5ms
	ErrorReason_SNOWFLAKE_TIME_EXCEPTION ErrorReason = 5
	ErrorReason_SNOWFLAKE_ID_ILLEGAL     ErrorReason = 6
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "SEG_SERVER_UNSPECIFIED",
		1: "BIZ_TAG_NOT_FOUND",
		2: "DB_OPERATE",
		3: "ID_CACHE_INIT_FALSE",
		4: "ID_TWO_SEGMENTS_ARE_NULL",
		5: "SNOWFLAKE_TIME_EXCEPTION",
		6: "SNOWFLAKE_ID_ILLEGAL",
	}
	ErrorReason_value = map[string]int32{
		"SEG_SERVER_UNSPECIFIED":   0,
		"BIZ_TAG_NOT_FOUND":        1,
		"DB_OPERATE":               2,
		"ID_CACHE_INIT_FALSE":      3,
		"ID_TWO_SEGMENTS_ARE_NULL": 4,
		"SNOWFLAKE_TIME_EXCEPTION": 5,
		"SNOWFLAKE_ID_ILLEGAL":     6,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_leaf_grpc_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_leaf_grpc_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_leaf_grpc_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_leaf_grpc_v1_error_reason_proto protoreflect.FileDescriptor

var file_leaf_grpc_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x65, 0x61, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6c, 0x65, 0x61, 0x66, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2a, 0xbf,
	0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x16, 0x53, 0x45, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x49,
	0x5a, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x42, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x49, 0x4e,
	0x49, 0x54, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x44,
	0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x52,
	0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x4e, 0x4f, 0x57,
	0x46, 0x4c, 0x41, 0x4b, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x4e, 0x4f, 0x57, 0x46, 0x4c,
	0x41, 0x4b, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x06,
	0x42, 0x21, 0x5a, 0x1f, 0x73, 0x65, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaf_grpc_v1_error_reason_proto_rawDescOnce sync.Once
	file_leaf_grpc_v1_error_reason_proto_rawDescData = file_leaf_grpc_v1_error_reason_proto_rawDesc
)

func file_leaf_grpc_v1_error_reason_proto_rawDescGZIP() []byte {
	file_leaf_grpc_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_leaf_grpc_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaf_grpc_v1_error_reason_proto_rawDescData)
	})
	return file_leaf_grpc_v1_error_reason_proto_rawDescData
}

var file_leaf_grpc_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_leaf_grpc_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: leafgrpc.v1.ErrorReason
}
var file_leaf_grpc_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_leaf_grpc_v1_error_reason_proto_init() }
func file_leaf_grpc_v1_error_reason_proto_init() {
	if File_leaf_grpc_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_leaf_grpc_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_leaf_grpc_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_leaf_grpc_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_leaf_grpc_v1_error_reason_proto_enumTypes,
	}.Build()
	File_leaf_grpc_v1_error_reason_proto = out.File
	file_leaf_grpc_v1_error_reason_proto_rawDesc = nil
	file_leaf_grpc_v1_error_reason_proto_goTypes = nil
	file_leaf_grpc_v1_error_reason_proto_depIdxs = nil
}
