// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: map.proto

package world

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

type GetMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMapRequest) Reset() {
	*x = GetMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapRequest) ProtoMessage() {}

func (x *GetMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapRequest.ProtoReflect.Descriptor instead.
func (*GetMapRequest) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{0}
}

type GetMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMapResponse) Reset() {
	*x = GetMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapResponse) ProtoMessage() {}

func (x *GetMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapResponse.ProtoReflect.Descriptor instead.
func (*GetMapResponse) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{1}
}

var File_map_proto protoreflect.FileDescriptor

var file_map_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x6f, 0x78,
	0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x43, 0x0a, 0x03, 0x4d, 0x61, 0x70,
	0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x78, 0x2e, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x78, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x2e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x78, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x3b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_map_proto_rawDescOnce sync.Once
	file_map_proto_rawDescData = file_map_proto_rawDesc
)

func file_map_proto_rawDescGZIP() []byte {
	file_map_proto_rawDescOnce.Do(func() {
		file_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_map_proto_rawDescData)
	})
	return file_map_proto_rawDescData
}

var file_map_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_map_proto_goTypes = []interface{}{
	(*GetMapRequest)(nil),  // 0: box.world.GetMapRequest
	(*GetMapResponse)(nil), // 1: box.world.GetMapResponse
}
var file_map_proto_depIdxs = []int32{
	0, // 0: box.world.Map.Get:input_type -> box.world.GetMapRequest
	1, // 1: box.world.Map.Get:output_type -> box.world.GetMapResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_map_proto_init() }
func file_map_proto_init() {
	if File_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapRequest); i {
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
		file_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapResponse); i {
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
			RawDescriptor: file_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_map_proto_goTypes,
		DependencyIndexes: file_map_proto_depIdxs,
		MessageInfos:      file_map_proto_msgTypes,
	}.Build()
	File_map_proto = out.File
	file_map_proto_rawDesc = nil
	file_map_proto_goTypes = nil
	file_map_proto_depIdxs = nil
}
