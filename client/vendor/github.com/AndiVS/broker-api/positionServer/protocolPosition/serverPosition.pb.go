// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protocolPosition/serverPosition.proto

package protocolPosition

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

type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyName   string  `protobuf:"bytes,1,opt,name=currencyName,proto3" json:"currencyName,omitempty"`
	CurrencyAmount int64   `protobuf:"varint,2,opt,name=currencyAmount,proto3" json:"currencyAmount,omitempty"`
	Price          float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocolPosition_serverPosition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocolPosition_serverPosition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_protocolPosition_serverPosition_proto_rawDescGZIP(), []int{0}
}

func (x *OpenRequest) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

func (x *OpenRequest) GetCurrencyAmount() int64 {
	if x != nil {
		return x.CurrencyAmount
	}
	return 0
}

func (x *OpenRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type OpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionID string `protobuf:"bytes,1,opt,name=positionID,proto3" json:"positionID,omitempty"`
}

func (x *OpenResponse) Reset() {
	*x = OpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocolPosition_serverPosition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenResponse) ProtoMessage() {}

func (x *OpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocolPosition_serverPosition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenResponse.ProtoReflect.Descriptor instead.
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return file_protocolPosition_serverPosition_proto_rawDescGZIP(), []int{1}
}

func (x *OpenResponse) GetPositionID() string {
	if x != nil {
		return x.PositionID
	}
	return ""
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionID   string `protobuf:"bytes,1,opt,name=positionID,proto3" json:"positionID,omitempty"`
	CurrencyName string `protobuf:"bytes,2,opt,name=currencyName,proto3" json:"currencyName,omitempty"`
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocolPosition_serverPosition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocolPosition_serverPosition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_protocolPosition_serverPosition_proto_rawDescGZIP(), []int{2}
}

func (x *CloseRequest) GetPositionID() string {
	if x != nil {
		return x.PositionID
	}
	return ""
}

func (x *CloseRequest) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocolPosition_serverPosition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocolPosition_serverPosition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_protocolPosition_serverPosition_proto_rawDescGZIP(), []int{3}
}

func (x *CloseResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protocolPosition_serverPosition_proto protoreflect.FileDescriptor

var file_protocolPosition_serverPosition_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x6f, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0d,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0xae, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protocolPosition_serverPosition_proto_rawDescOnce sync.Once
	file_protocolPosition_serverPosition_proto_rawDescData = file_protocolPosition_serverPosition_proto_rawDesc
)

func file_protocolPosition_serverPosition_proto_rawDescGZIP() []byte {
	file_protocolPosition_serverPosition_proto_rawDescOnce.Do(func() {
		file_protocolPosition_serverPosition_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocolPosition_serverPosition_proto_rawDescData)
	})
	return file_protocolPosition_serverPosition_proto_rawDescData
}

var file_protocolPosition_serverPosition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocolPosition_serverPosition_proto_goTypes = []interface{}{
	(*OpenRequest)(nil),   // 0: protocolBroker.OpenRequest
	(*OpenResponse)(nil),  // 1: protocolBroker.OpenResponse
	(*CloseRequest)(nil),  // 2: protocolBroker.CloseRequest
	(*CloseResponse)(nil), // 3: protocolBroker.CloseResponse
}
var file_protocolPosition_serverPosition_proto_depIdxs = []int32{
	0, // 0: protocolBroker.PositionService.OpenPosition:input_type -> protocolBroker.OpenRequest
	2, // 1: protocolBroker.PositionService.ClosePosition:input_type -> protocolBroker.CloseRequest
	1, // 2: protocolBroker.PositionService.OpenPosition:output_type -> protocolBroker.OpenResponse
	3, // 3: protocolBroker.PositionService.ClosePosition:output_type -> protocolBroker.CloseResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocolPosition_serverPosition_proto_init() }
func file_protocolPosition_serverPosition_proto_init() {
	if File_protocolPosition_serverPosition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocolPosition_serverPosition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRequest); i {
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
		file_protocolPosition_serverPosition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenResponse); i {
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
		file_protocolPosition_serverPosition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_protocolPosition_serverPosition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
			RawDescriptor: file_protocolPosition_serverPosition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocolPosition_serverPosition_proto_goTypes,
		DependencyIndexes: file_protocolPosition_serverPosition_proto_depIdxs,
		MessageInfos:      file_protocolPosition_serverPosition_proto_msgTypes,
	}.Build()
	File_protocolPosition_serverPosition_proto = out.File
	file_protocolPosition_serverPosition_proto_rawDesc = nil
	file_protocolPosition_serverPosition_proto_goTypes = nil
	file_protocolPosition_serverPosition_proto_depIdxs = nil
}
