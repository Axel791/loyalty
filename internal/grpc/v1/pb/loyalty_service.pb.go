// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: loyalty_service.proto

package pb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос c user_id
type CreateLoyaltyBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLoyaltyBalanceRequest) Reset() {
	*x = CreateLoyaltyBalanceRequest{}
	mi := &file_loyalty_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLoyaltyBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoyaltyBalanceRequest) ProtoMessage() {}

func (x *CreateLoyaltyBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loyalty_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoyaltyBalanceRequest.ProtoReflect.Descriptor instead.
func (*CreateLoyaltyBalanceRequest) Descriptor() ([]byte, []int) {
	return file_loyalty_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLoyaltyBalanceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateLoyaltyBalanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLoyaltyBalanceResponse) Reset() {
	*x = CreateLoyaltyBalanceResponse{}
	mi := &file_loyalty_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLoyaltyBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoyaltyBalanceResponse) ProtoMessage() {}

func (x *CreateLoyaltyBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loyalty_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoyaltyBalanceResponse.ProtoReflect.Descriptor instead.
func (*CreateLoyaltyBalanceResponse) Descriptor() ([]byte, []int) {
	return file_loyalty_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLoyaltyBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateLoyaltyBalanceResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_loyalty_service_proto protoreflect.FileDescriptor

var file_loyalty_service_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79,
	0x22, 0x36, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74,
	0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x75, 0x0a, 0x0e, 0x4c, 0x6f, 0x79, 0x61, 0x6c,
	0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x78, 0x65,
	0x6c, 0x37, 0x39, 0x31, 0x2f, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_loyalty_service_proto_rawDescOnce sync.Once
	file_loyalty_service_proto_rawDescData []byte
)

func file_loyalty_service_proto_rawDescGZIP() []byte {
	file_loyalty_service_proto_rawDescOnce.Do(func() {
		file_loyalty_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_loyalty_service_proto_rawDesc), len(file_loyalty_service_proto_rawDesc)))
	})
	return file_loyalty_service_proto_rawDescData
}

var file_loyalty_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loyalty_service_proto_goTypes = []any{
	(*CreateLoyaltyBalanceRequest)(nil),  // 0: loyalty.CreateLoyaltyBalanceRequest
	(*CreateLoyaltyBalanceResponse)(nil), // 1: loyalty.CreateLoyaltyBalanceResponse
}
var file_loyalty_service_proto_depIdxs = []int32{
	0, // 0: loyalty.LoyaltyService.CreateLoyaltyBalance:input_type -> loyalty.CreateLoyaltyBalanceRequest
	1, // 1: loyalty.LoyaltyService.CreateLoyaltyBalance:output_type -> loyalty.CreateLoyaltyBalanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loyalty_service_proto_init() }
func file_loyalty_service_proto_init() {
	if File_loyalty_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_loyalty_service_proto_rawDesc), len(file_loyalty_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loyalty_service_proto_goTypes,
		DependencyIndexes: file_loyalty_service_proto_depIdxs,
		MessageInfos:      file_loyalty_service_proto_msgTypes,
	}.Build()
	File_loyalty_service_proto = out.File
	file_loyalty_service_proto_goTypes = nil
	file_loyalty_service_proto_depIdxs = nil
}
