// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: plugin/plugin.proto

package pluginpb

import (
	_ "github.com/mercari/grpc-federation/grpc/federation"
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

type Regexp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ptr uint64 `protobuf:"varint,1,opt,name=ptr,proto3" json:"ptr,omitempty"` // store raw pointer value.
}

func (x *Regexp) Reset() {
	*x = Regexp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regexp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regexp) ProtoMessage() {}

func (x *Regexp) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regexp.ProtoReflect.Descriptor instead.
func (*Regexp) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Regexp) GetPtr() uint64 {
	if x != nil {
		return x.Ptr
	}
	return 0
}

var File_plugin_plugin_proto protoreflect.FileDescriptor

var file_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x70, 0x1a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x74, 0x72, 0x42,
	0xe6, 0x04, 0x82, 0x97, 0x22, 0xce, 0x03, 0x0a, 0xcb, 0x03, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x70, 0x1a, 0xea, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x3d, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x70, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61,
	0x20, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xa0, 0x01, 0x0a,
	0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x55, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x20, 0x77, 0x68, 0x65, 0x74, 0x68, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20,
	0x61, 0x6e, 0x79, 0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x0a, 0x01, 0x73, 0x12, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x20, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x20, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x69, 0x66, 0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x2c, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x72, 0x75, 0x65, 0x22,
	0xaf, 0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x12, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x20, 0x70, 0x61, 0x72, 0x73, 0x65, 0x73, 0x20, 0x61, 0x20, 0x72, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2c, 0x20, 0x69, 0x66,
	0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x2c, 0x20, 0x61, 0x20, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x70, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x20,
	0x61, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x20, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x29, 0x0a, 0x04,
	0x65, 0x78, 0x70, 0x72, 0x12, 0x19, 0x61, 0x20, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x20,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x65, 0x78, 0x74, 0x1a,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x08, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x70, 0x2a, 0x22, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x13, 0x61, 0x20, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x42, 0x0b, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x17, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70,
	0x62, 0xa2, 0x02, 0x03, 0x45, 0x52, 0x58, 0xaa, 0x02, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0xca, 0x02, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0xe2, 0x02, 0x1a, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5c, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x3a, 0x3a, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_plugin_proto_rawDescOnce sync.Once
	file_plugin_plugin_proto_rawDescData = file_plugin_plugin_proto_rawDesc
)

func file_plugin_plugin_proto_rawDescGZIP() []byte {
	file_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_plugin_proto_rawDescData)
	})
	return file_plugin_plugin_proto_rawDescData
}

var file_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_plugin_plugin_proto_goTypes = []interface{}{
	(*Regexp)(nil), // 0: example.regexp.Regexp
}
var file_plugin_plugin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_plugin_proto_init() }
func file_plugin_plugin_proto_init() {
	if File_plugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regexp); i {
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
			RawDescriptor: file_plugin_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_plugin_proto_msgTypes,
	}.Build()
	File_plugin_plugin_proto = out.File
	file_plugin_plugin_proto_rawDesc = nil
	file_plugin_plugin_proto_goTypes = nil
	file_plugin_plugin_proto_depIdxs = nil
}
