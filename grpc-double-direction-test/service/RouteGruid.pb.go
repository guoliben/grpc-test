// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: RouteGruid.proto

package hello_grpc_proto

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

type RouteNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slice   string `protobuf:"bytes,1,opt,name=slice,proto3" json:"slice,omitempty"`
	Time    string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Block   string `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
	Detail  string `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RouteNoteRequest) Reset() {
	*x = RouteNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RouteGruid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNoteRequest) ProtoMessage() {}

func (x *RouteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_RouteGruid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNoteRequest.ProtoReflect.Descriptor instead.
func (*RouteNoteRequest) Descriptor() ([]byte, []int) {
	return file_RouteGruid_proto_rawDescGZIP(), []int{0}
}

func (x *RouteNoteRequest) GetSlice() string {
	if x != nil {
		return x.Slice
	}
	return ""
}

func (x *RouteNoteRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *RouteNoteRequest) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *RouteNoteRequest) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *RouteNoteRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RouteNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack  string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Next string `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *RouteNoteResponse) Reset() {
	*x = RouteNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RouteGruid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNoteResponse) ProtoMessage() {}

func (x *RouteNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_RouteGruid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNoteResponse.ProtoReflect.Descriptor instead.
func (*RouteNoteResponse) Descriptor() ([]byte, []int) {
	return file_RouteGruid_proto_rawDescGZIP(), []int{1}
}

func (x *RouteNoteResponse) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

func (x *RouteNoteResponse) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

var File_RouteGruid_proto protoreflect.FileDescriptor

var file_RouteGruid_proto_rawDesc = []byte{
	0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5f, 0x0a, 0x10, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0d, 0x0a, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0c,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0d, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0e, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0f, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x22, 0x2e, 0x0a, 0x11,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0b, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0c,
	0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x32, 0x55, 0x0a, 0x09,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RouteGruid_proto_rawDescOnce sync.Once
	file_RouteGruid_proto_rawDescData = file_RouteGruid_proto_rawDesc
)

func file_RouteGruid_proto_rawDescGZIP() []byte {
	file_RouteGruid_proto_rawDescOnce.Do(func() {
		file_RouteGruid_proto_rawDescData = protoimpl.X.CompressGZIP(file_RouteGruid_proto_rawDescData)
	})
	return file_RouteGruid_proto_rawDescData
}

var file_RouteGruid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RouteGruid_proto_goTypes = []interface{}{
	(*RouteNoteRequest)(nil),  // 0: service.RouteNoteRequest
	(*RouteNoteResponse)(nil), // 1: service.RouteNoteResponse
}
var file_RouteGruid_proto_depIdxs = []int32{
	0, // 0: service.RouteGuid.RouteChat:input_type -> service.RouteNoteRequest
	1, // 1: service.RouteGuid.RouteChat:output_type -> service.RouteNoteResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RouteGruid_proto_init() }
func file_RouteGruid_proto_init() {
	if File_RouteGruid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RouteGruid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNoteRequest); i {
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
		file_RouteGruid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNoteResponse); i {
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
			RawDescriptor: file_RouteGruid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RouteGruid_proto_goTypes,
		DependencyIndexes: file_RouteGruid_proto_depIdxs,
		MessageInfos:      file_RouteGruid_proto_msgTypes,
	}.Build()
	File_RouteGruid_proto = out.File
	file_RouteGruid_proto_rawDesc = nil
	file_RouteGruid_proto_goTypes = nil
	file_RouteGruid_proto_depIdxs = nil
}
