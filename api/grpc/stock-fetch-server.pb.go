// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/grpc/stock-fetch-server.proto

package grpcapi

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

type StockConfigUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockName    string `protobuf:"bytes,1,opt,name=stock_name,json=stockName,proto3" json:"stock_name,omitempty"`
	OptionType   int32  `protobuf:"varint,2,opt,name=option_type,json=optionType,proto3" json:"option_type,omitempty"`
	OptionStatus bool   `protobuf:"varint,3,opt,name=option_status,json=optionStatus,proto3" json:"option_status,omitempty"`
}

func (x *StockConfigUpdateRequest) Reset() {
	*x = StockConfigUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_stock_fetch_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockConfigUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockConfigUpdateRequest) ProtoMessage() {}

func (x *StockConfigUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_stock_fetch_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockConfigUpdateRequest.ProtoReflect.Descriptor instead.
func (*StockConfigUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_stock_fetch_server_proto_rawDescGZIP(), []int{0}
}

func (x *StockConfigUpdateRequest) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

func (x *StockConfigUpdateRequest) GetOptionType() int32 {
	if x != nil {
		return x.OptionType
	}
	return 0
}

func (x *StockConfigUpdateRequest) GetOptionStatus() bool {
	if x != nil {
		return x.OptionStatus
	}
	return false
}

type StockConfigUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StockConfigUpdateResponse) Reset() {
	*x = StockConfigUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_stock_fetch_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockConfigUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockConfigUpdateResponse) ProtoMessage() {}

func (x *StockConfigUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_stock_fetch_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockConfigUpdateResponse.ProtoReflect.Descriptor instead.
func (*StockConfigUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_stock_fetch_server_proto_rawDescGZIP(), []int{1}
}

var File_api_grpc_stock_fetch_server_proto protoreflect.FileDescriptor

var file_api_grpc_stock_fetch_server_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2d, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x7f, 0x0a, 0x18, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x70, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_stock_fetch_server_proto_rawDescOnce sync.Once
	file_api_grpc_stock_fetch_server_proto_rawDescData = file_api_grpc_stock_fetch_server_proto_rawDesc
)

func file_api_grpc_stock_fetch_server_proto_rawDescGZIP() []byte {
	file_api_grpc_stock_fetch_server_proto_rawDescOnce.Do(func() {
		file_api_grpc_stock_fetch_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_stock_fetch_server_proto_rawDescData)
	})
	return file_api_grpc_stock_fetch_server_proto_rawDescData
}

var file_api_grpc_stock_fetch_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_stock_fetch_server_proto_goTypes = []interface{}{
	(*StockConfigUpdateRequest)(nil),  // 0: api.StockConfigUpdateRequest
	(*StockConfigUpdateResponse)(nil), // 1: api.StockConfigUpdateResponse
}
var file_api_grpc_stock_fetch_server_proto_depIdxs = []int32{
	0, // 0: api.StockConfigurator.UpdateStockConfiguration:input_type -> api.StockConfigUpdateRequest
	1, // 1: api.StockConfigurator.UpdateStockConfiguration:output_type -> api.StockConfigUpdateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_stock_fetch_server_proto_init() }
func file_api_grpc_stock_fetch_server_proto_init() {
	if File_api_grpc_stock_fetch_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_stock_fetch_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockConfigUpdateRequest); i {
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
		file_api_grpc_stock_fetch_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockConfigUpdateResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_grpc_stock_fetch_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_stock_fetch_server_proto_goTypes,
		DependencyIndexes: file_api_grpc_stock_fetch_server_proto_depIdxs,
		MessageInfos:      file_api_grpc_stock_fetch_server_proto_msgTypes,
	}.Build()
	File_api_grpc_stock_fetch_server_proto = out.File
	file_api_grpc_stock_fetch_server_proto_rawDesc = nil
	file_api_grpc_stock_fetch_server_proto_goTypes = nil
	file_api_grpc_stock_fetch_server_proto_depIdxs = nil
}