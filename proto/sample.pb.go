// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/sample.proto

package proto

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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_proto_sample_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *HelloRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res1 int32 `protobuf:"varint,1,opt,name=res1,proto3" json:"res1,omitempty"`
	Res2 int32 `protobuf:"varint,2,opt,name=res2,proto3" json:"res2,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	mi := &file_proto_sample_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetRes1() int32 {
	if x != nil {
		return x.Res1
	}
	return 0
}

func (x *HelloReply) GetRes2() int32 {
	if x != nil {
		return x.Res2
	}
	return 0
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	mi := &file_proto_sample_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{2}
}

func (x *VoteRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type VoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granted bool `protobuf:"varint,1,opt,name=granted,proto3" json:"granted,omitempty"`
}

func (x *VoteReply) Reset() {
	*x = VoteReply{}
	mi := &file_proto_sample_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteReply) ProtoMessage() {}

func (x *VoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteReply.ProtoReflect.Descriptor instead.
func (*VoteReply) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{3}
}

func (x *VoteReply) GetGranted() bool {
	if x != nil {
		return x.Granted
	}
	return false
}

type HeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartBeatRequest) Reset() {
	*x = HeartBeatRequest{}
	mi := &file_proto_sample_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRequest) ProtoMessage() {}

func (x *HeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{4}
}

type HeartBeatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *HeartBeatReply) Reset() {
	*x = HeartBeatReply{}
	mi := &file_proto_sample_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartBeatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatReply) ProtoMessage() {}

func (x *HeartBeatReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatReply.ProtoReflect.Descriptor instead.
func (*HeartBeatReply) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{5}
}

func (x *HeartBeatReply) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

var File_proto_sample_proto protoreflect.FileDescriptor

var file_proto_sample_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x36, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x22, 0x34, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x65, 0x73, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x65, 0x73, 0x32, 0x22, 0x21, 0x0a, 0x0b, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x25, 0x0a,
	0x09, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0xbb, 0x01, 0x0a,
	0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sample_proto_rawDescOnce sync.Once
	file_proto_sample_proto_rawDescData = file_proto_sample_proto_rawDesc
)

func file_proto_sample_proto_rawDescGZIP() []byte {
	file_proto_sample_proto_rawDescOnce.Do(func() {
		file_proto_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sample_proto_rawDescData)
	})
	return file_proto_sample_proto_rawDescData
}

var file_proto_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_sample_proto_goTypes = []any{
	(*HelloRequest)(nil),     // 0: sample.HelloRequest
	(*HelloReply)(nil),       // 1: sample.HelloReply
	(*VoteRequest)(nil),      // 2: sample.VoteRequest
	(*VoteReply)(nil),        // 3: sample.VoteReply
	(*HeartBeatRequest)(nil), // 4: sample.HeartBeatRequest
	(*HeartBeatReply)(nil),   // 5: sample.HeartBeatReply
}
var file_proto_sample_proto_depIdxs = []int32{
	0, // 0: sample.Greeter.SayHello:input_type -> sample.HelloRequest
	2, // 1: sample.Greeter.RequestVote:input_type -> sample.VoteRequest
	4, // 2: sample.Greeter.HeartBeat:input_type -> sample.HeartBeatRequest
	1, // 3: sample.Greeter.SayHello:output_type -> sample.HelloReply
	3, // 4: sample.Greeter.RequestVote:output_type -> sample.VoteReply
	5, // 5: sample.Greeter.HeartBeat:output_type -> sample.HeartBeatReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sample_proto_init() }
func file_proto_sample_proto_init() {
	if File_proto_sample_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sample_proto_goTypes,
		DependencyIndexes: file_proto_sample_proto_depIdxs,
		MessageInfos:      file_proto_sample_proto_msgTypes,
	}.Build()
	File_proto_sample_proto = out.File
	file_proto_sample_proto_rawDesc = nil
	file_proto_sample_proto_goTypes = nil
	file_proto_sample_proto_depIdxs = nil
}
