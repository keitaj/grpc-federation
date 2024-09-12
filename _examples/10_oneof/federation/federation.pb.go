// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: federation/federation.proto

package federation

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{0}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetNoValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNoValueRequest) Reset() {
	*x = GetNoValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoValueRequest) ProtoMessage() {}

func (x *GetNoValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoValueRequest.ProtoReflect.Descriptor instead.
func (*GetNoValueRequest) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{2}
}

type GetNoValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoValue *M `protobuf:"bytes,1,opt,name=no_value,json=noValue,proto3" json:"no_value,omitempty"`
}

func (x *GetNoValueResponse) Reset() {
	*x = GetNoValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoValueResponse) ProtoMessage() {}

func (x *GetNoValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoValueResponse.ProtoReflect.Descriptor instead.
func (*GetNoValueResponse) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{3}
}

func (x *GetNoValueResponse) GetNoValue() *M {
	if x != nil {
		return x.NoValue
	}
	return nil
}

type UserSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to User:
	//
	//	*UserSelection_UserA
	//	*UserSelection_UserB
	//	*UserSelection_UserC
	User isUserSelection_User `protobuf_oneof:"user"`
}

func (x *UserSelection) Reset() {
	*x = UserSelection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSelection) ProtoMessage() {}

func (x *UserSelection) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSelection.ProtoReflect.Descriptor instead.
func (*UserSelection) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{4}
}

func (m *UserSelection) GetUser() isUserSelection_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (x *UserSelection) GetUserA() *User {
	if x, ok := x.GetUser().(*UserSelection_UserA); ok {
		return x.UserA
	}
	return nil
}

func (x *UserSelection) GetUserB() *User {
	if x, ok := x.GetUser().(*UserSelection_UserB); ok {
		return x.UserB
	}
	return nil
}

func (x *UserSelection) GetUserC() *User {
	if x, ok := x.GetUser().(*UserSelection_UserC); ok {
		return x.UserC
	}
	return nil
}

type isUserSelection_User interface {
	isUserSelection_User()
}

type UserSelection_UserA struct {
	UserA *User `protobuf:"bytes,1,opt,name=user_a,json=userA,proto3,oneof"`
}

type UserSelection_UserB struct {
	UserB *User `protobuf:"bytes,2,opt,name=user_b,json=userB,proto3,oneof"`
}

type UserSelection_UserC struct {
	UserC *User `protobuf:"bytes,3,opt,name=user_c,json=userC,proto3,oneof"`
}

func (*UserSelection_UserA) isUserSelection_User() {}

func (*UserSelection_UserB) isUserSelection_User() {}

func (*UserSelection_UserC) isUserSelection_User() {}

type MessageSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*MessageSelection_MsgA
	//	*MessageSelection_MsgB
	//	*MessageSelection_MsgC
	Message isMessageSelection_Message `protobuf_oneof:"message"`
}

func (x *MessageSelection) Reset() {
	*x = MessageSelection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSelection) ProtoMessage() {}

func (x *MessageSelection) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSelection.ProtoReflect.Descriptor instead.
func (*MessageSelection) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{5}
}

func (m *MessageSelection) GetMessage() isMessageSelection_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *MessageSelection) GetMsgA() string {
	if x, ok := x.GetMessage().(*MessageSelection_MsgA); ok {
		return x.MsgA
	}
	return ""
}

func (x *MessageSelection) GetMsgB() string {
	if x, ok := x.GetMessage().(*MessageSelection_MsgB); ok {
		return x.MsgB
	}
	return ""
}

func (x *MessageSelection) GetMsgC() string {
	if x, ok := x.GetMessage().(*MessageSelection_MsgC); ok {
		return x.MsgC
	}
	return ""
}

type isMessageSelection_Message interface {
	isMessageSelection_Message()
}

type MessageSelection_MsgA struct {
	MsgA string `protobuf:"bytes,1,opt,name=msg_a,json=msgA,proto3,oneof"`
}

type MessageSelection_MsgB struct {
	MsgB string `protobuf:"bytes,2,opt,name=msg_b,json=msgB,proto3,oneof"`
}

type MessageSelection_MsgC struct {
	MsgC string `protobuf:"bytes,3,opt,name=msg_c,json=msgC,proto3,oneof"`
}

func (*MessageSelection_MsgA) isMessageSelection_Message() {}

func (*MessageSelection_MsgB) isMessageSelection_Message() {}

func (*MessageSelection_MsgC) isMessageSelection_Message() {}

type NoValueSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NoValue:
	//
	//	*NoValueSelection_MA
	//	*NoValueSelection_MB
	NoValue isNoValueSelection_NoValue `protobuf_oneof:"no_value"`
}

func (x *NoValueSelection) Reset() {
	*x = NoValueSelection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoValueSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoValueSelection) ProtoMessage() {}

func (x *NoValueSelection) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoValueSelection.ProtoReflect.Descriptor instead.
func (*NoValueSelection) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{6}
}

func (m *NoValueSelection) GetNoValue() isNoValueSelection_NoValue {
	if m != nil {
		return m.NoValue
	}
	return nil
}

func (x *NoValueSelection) GetMA() *M {
	if x, ok := x.GetNoValue().(*NoValueSelection_MA); ok {
		return x.MA
	}
	return nil
}

func (x *NoValueSelection) GetMB() *M {
	if x, ok := x.GetNoValue().(*NoValueSelection_MB); ok {
		return x.MB
	}
	return nil
}

type isNoValueSelection_NoValue interface {
	isNoValueSelection_NoValue()
}

type NoValueSelection_MA struct {
	MA *M `protobuf:"bytes,1,opt,name=m_a,json=mA,proto3,oneof"`
}

type NoValueSelection_MB struct {
	MB *M `protobuf:"bytes,2,opt,name=m_b,json=mB,proto3,oneof"`
}

func (*NoValueSelection_MA) isNoValueSelection_NoValue() {}

func (*NoValueSelection_MB) isNoValueSelection_NoValue() {}

type M struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *M) Reset() {
	*x = M{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M) ProtoMessage() {}

func (x *M) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M.ProtoReflect.Descriptor instead.
func (*M) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{7}
}

func (x *M) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_federation_federation_proto protoreflect.FileDescriptor

var file_federation_federation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f,
	0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xba, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x0d, 0x9a, 0x4a, 0x0a, 0x12, 0x08, 0x73, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x14, 0x9a, 0x4a, 0x11, 0x12, 0x0f, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65,
	0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x4a,
	0x9a, 0x4a, 0x47, 0x0a, 0x26, 0x0a, 0x03, 0x73, 0x65, 0x6c, 0x6a, 0x1f, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0x27, 0x66, 0x6f, 0x6f, 0x27, 0x0a, 0x1d, 0x0a, 0x07, 0x6d,
	0x73, 0x67, 0x5f, 0x73, 0x65, 0x6c, 0x6a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x87, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x42, 0x1a, 0x9a, 0x4a, 0x17,
	0x12, 0x15, 0x6e, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x2e, 0x6e,
	0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x27, 0x9a, 0x4a, 0x24, 0x0a, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x73, 0x65, 0x6c, 0x6a, 0x12, 0x0a, 0x10, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x03, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x48, 0x9a, 0x4a, 0x45, 0x22, 0x43, 0x1a, 0x36, 0x0a, 0x02, 0x75, 0x61, 0x6a, 0x30,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x12, 0x03, 0x27, 0x61, 0x27, 0x12, 0x08, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x01, 0x30,
	0x12, 0x0e, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x12, 0x07, 0x27, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x27,
	0x22, 0x02, 0x75, 0x61, 0x0a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x12, 0x76, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x47, 0x9a, 0x4a, 0x44, 0x22,
	0x42, 0x1a, 0x36, 0x0a, 0x02, 0x75, 0x62, 0x6a, 0x30, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x03, 0x27, 0x62, 0x27, 0x12,
	0x08, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x01, 0x30, 0x12, 0x0e, 0x0a, 0x03, 0x62, 0x61, 0x72,
	0x12, 0x07, 0x27, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x27, 0x22, 0x02, 0x75, 0x62, 0x0a, 0x04, 0x74,
	0x72, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x42, 0x12, 0x76, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x47, 0x9a, 0x4a, 0x44, 0x22, 0x42, 0x1a, 0x3a, 0x0a, 0x02, 0x75, 0x63, 0x6a,
	0x34, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x12, 0x07, 0x24, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x08, 0x0a, 0x03, 0x66,
	0x6f, 0x6f, 0x12, 0x01, 0x30, 0x12, 0x0e, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x12, 0x07, 0x27, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x27, 0x22, 0x02, 0x75, 0x63, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x42, 0x06, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x9b, 0x01, 0x0a,
	0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x13, 0x9a, 0x4a, 0x10, 0x22, 0x0e, 0x22, 0x05, 0x27, 0x61, 0x61, 0x61, 0x27, 0x0a, 0x05,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x41, 0x12, 0x29, 0x0a,
	0x05, 0x6d, 0x73, 0x67, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x4a,
	0x0f, 0x22, 0x0d, 0x22, 0x05, 0x27, 0x62, 0x62, 0x62, 0x27, 0x0a, 0x04, 0x74, 0x72, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x42, 0x12, 0x25, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x5f,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x9a, 0x4a, 0x0b, 0x22, 0x09, 0x22, 0x05,
	0x27, 0x63, 0x63, 0x63, 0x27, 0x10, 0x01, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x43, 0x42,
	0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x10, 0x4e,
	0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x41, 0x0a, 0x03, 0x6d, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x42,
	0x1b, 0x9a, 0x4a, 0x18, 0x22, 0x16, 0x22, 0x0d, 0x4d, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x20, 0x27, 0x61, 0x27, 0x7d, 0x0a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x48, 0x00, 0x52, 0x02,
	0x6d, 0x41, 0x12, 0x41, 0x0a, 0x03, 0x6d, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x42, 0x1b, 0x9a, 0x4a, 0x18, 0x22, 0x16, 0x22, 0x0d, 0x4d, 0x7b, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x20, 0x27, 0x62, 0x27, 0x7d, 0x0a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x02, 0x6d, 0x42, 0x42, 0x0a, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x25, 0x0a, 0x01, 0x4d, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x9a, 0x4a, 0x07, 0x12, 0x05, 0x27, 0x66, 0x6f, 0x6f,
	0x27, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x9a,
	0x4a, 0x0b, 0x12, 0x09, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x3a, 0x67, 0x9a, 0x4a, 0x64, 0x0a, 0x62, 0x72, 0x60, 0x0a, 0x18, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x09, 0x24, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x05, 0x24, 0x2e,
	0x66, 0x6f, 0x6f, 0x1a, 0x0a, 0x24, 0x2e, 0x66, 0x6f, 0x6f, 0x20, 0x21, 0x3d, 0x20, 0x30, 0x12,
	0x19, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x12, 0x05, 0x24, 0x2e, 0x62, 0x61, 0x72, 0x1a, 0x0b, 0x24,
	0x2e, 0x62, 0x61, 0x72, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0x32, 0xb1, 0x01, 0x0a, 0x11, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x40, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x03, 0x9a, 0x4a, 0x00, 0x42, 0xb1,
	0x01, 0x9a, 0x4a, 0x11, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x3b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x4f, 0x46,
	0x58, 0xaa, 0x02, 0x0e, 0x4f, 0x72, 0x67, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0xca, 0x02, 0x0e, 0x4f, 0x72, 0x67, 0x5c, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x1a, 0x4f, 0x72, 0x67, 0x5c, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_federation_federation_proto_rawDescOnce sync.Once
	file_federation_federation_proto_rawDescData = file_federation_federation_proto_rawDesc
)

func file_federation_federation_proto_rawDescGZIP() []byte {
	file_federation_federation_proto_rawDescOnce.Do(func() {
		file_federation_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_federation_federation_proto_rawDescData)
	})
	return file_federation_federation_proto_rawDescData
}

var file_federation_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_federation_federation_proto_goTypes = []interface{}{
	(*GetRequest)(nil),         // 0: org.federation.GetRequest
	(*GetResponse)(nil),        // 1: org.federation.GetResponse
	(*GetNoValueRequest)(nil),  // 2: org.federation.GetNoValueRequest
	(*GetNoValueResponse)(nil), // 3: org.federation.GetNoValueResponse
	(*UserSelection)(nil),      // 4: org.federation.UserSelection
	(*MessageSelection)(nil),   // 5: org.federation.MessageSelection
	(*NoValueSelection)(nil),   // 6: org.federation.NoValueSelection
	(*M)(nil),                  // 7: org.federation.M
	(*User)(nil),               // 8: org.federation.User
}
var file_federation_federation_proto_depIdxs = []int32{
	8, // 0: org.federation.GetResponse.user:type_name -> org.federation.User
	7, // 1: org.federation.GetNoValueResponse.no_value:type_name -> org.federation.M
	8, // 2: org.federation.UserSelection.user_a:type_name -> org.federation.User
	8, // 3: org.federation.UserSelection.user_b:type_name -> org.federation.User
	8, // 4: org.federation.UserSelection.user_c:type_name -> org.federation.User
	7, // 5: org.federation.NoValueSelection.m_a:type_name -> org.federation.M
	7, // 6: org.federation.NoValueSelection.m_b:type_name -> org.federation.M
	0, // 7: org.federation.FederationService.Get:input_type -> org.federation.GetRequest
	2, // 8: org.federation.FederationService.GetNoValue:input_type -> org.federation.GetNoValueRequest
	1, // 9: org.federation.FederationService.Get:output_type -> org.federation.GetResponse
	3, // 10: org.federation.FederationService.GetNoValue:output_type -> org.federation.GetNoValueResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_federation_federation_proto_init() }
func file_federation_federation_proto_init() {
	if File_federation_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_federation_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_federation_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_federation_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoValueRequest); i {
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
		file_federation_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoValueResponse); i {
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
		file_federation_federation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSelection); i {
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
		file_federation_federation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSelection); i {
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
		file_federation_federation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoValueSelection); i {
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
		file_federation_federation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M); i {
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
		file_federation_federation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
	file_federation_federation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UserSelection_UserA)(nil),
		(*UserSelection_UserB)(nil),
		(*UserSelection_UserC)(nil),
	}
	file_federation_federation_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*MessageSelection_MsgA)(nil),
		(*MessageSelection_MsgB)(nil),
		(*MessageSelection_MsgC)(nil),
	}
	file_federation_federation_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*NoValueSelection_MA)(nil),
		(*NoValueSelection_MB)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_federation_federation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_federation_federation_proto_goTypes,
		DependencyIndexes: file_federation_federation_proto_depIdxs,
		MessageInfos:      file_federation_federation_proto_msgTypes,
	}.Build()
	File_federation_federation_proto = out.File
	file_federation_federation_proto_rawDesc = nil
	file_federation_federation_proto_goTypes = nil
	file_federation_federation_proto_depIdxs = nil
}
