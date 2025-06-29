// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: another_service_name.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SomeEnum int32

const (
	SomeEnum_SOME_ENUM_UNSPECIFIED     SomeEnum = 0
	SomeEnum_SOME_ENUM_VARIANT         SomeEnum = 1
	SomeEnum_SOME_ENUM_ANOTHER_VARIANT SomeEnum = 2
)

// Enum value maps for SomeEnum.
var (
	SomeEnum_name = map[int32]string{
		0: "SOME_ENUM_UNSPECIFIED",
		1: "SOME_ENUM_VARIANT",
		2: "SOME_ENUM_ANOTHER_VARIANT",
	}
	SomeEnum_value = map[string]int32{
		"SOME_ENUM_UNSPECIFIED":     0,
		"SOME_ENUM_VARIANT":         1,
		"SOME_ENUM_ANOTHER_VARIANT": 2,
	}
)

func (x SomeEnum) Enum() *SomeEnum {
	p := new(SomeEnum)
	*p = x
	return p
}

func (x SomeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SomeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_another_service_name_proto_enumTypes[0].Descriptor()
}

func (SomeEnum) Type() protoreflect.EnumType {
	return &file_another_service_name_proto_enumTypes[0]
}

func (x SomeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SomeEnum.Descriptor instead.
func (SomeEnum) EnumDescriptor() ([]byte, []int) {
	return file_another_service_name_proto_rawDescGZIP(), []int{0}
}

type SomeMethodRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         SomeEnum               `protobuf:"varint,1,opt,name=value,proto3,enum=another_service_name.v1.SomeEnum" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomeMethodRequest) Reset() {
	*x = SomeMethodRequest{}
	mi := &file_another_service_name_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeMethodRequest) ProtoMessage() {}

func (x *SomeMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_another_service_name_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeMethodRequest.ProtoReflect.Descriptor instead.
func (*SomeMethodRequest) Descriptor() ([]byte, []int) {
	return file_another_service_name_proto_rawDescGZIP(), []int{0}
}

func (x *SomeMethodRequest) GetValue() SomeEnum {
	if x != nil {
		return x.Value
	}
	return SomeEnum_SOME_ENUM_UNSPECIFIED
}

type SomeMethodResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomeMethodResponse) Reset() {
	*x = SomeMethodResponse{}
	mi := &file_another_service_name_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeMethodResponse) ProtoMessage() {}

func (x *SomeMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_another_service_name_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeMethodResponse.ProtoReflect.Descriptor instead.
func (*SomeMethodResponse) Descriptor() ([]byte, []int) {
	return file_another_service_name_proto_rawDescGZIP(), []int{1}
}

func (x *SomeMethodResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_another_service_name_proto protoreflect.FileDescriptor

const file_another_service_name_proto_rawDesc = "" +
	"\n" +
	"\x1aanother_service_name.proto\x12\x17another_service_name.v1\"L\n" +
	"\x11SomeMethodRequest\x127\n" +
	"\x05value\x18\x01 \x01(\x0e2!.another_service_name.v1.SomeEnumR\x05value\".\n" +
	"\x12SomeMethodResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage*[\n" +
	"\bSomeEnum\x12\x19\n" +
	"\x15SOME_ENUM_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11SOME_ENUM_VARIANT\x10\x01\x12\x1d\n" +
	"\x19SOME_ENUM_ANOTHER_VARIANT\x10\x022t\n" +
	"\vSomeService\x12e\n" +
	"\n" +
	"SomeMethod\x12*.another_service_name.v1.SomeMethodRequest\x1a+.another_service_name.v1.SomeMethodResponseBSZQgithub.com/Sssilencee/protobuf-guidelines/protobuf/another_service_name/gen/go;pbb\x06proto3"

var (
	file_another_service_name_proto_rawDescOnce sync.Once
	file_another_service_name_proto_rawDescData []byte
)

func file_another_service_name_proto_rawDescGZIP() []byte {
	file_another_service_name_proto_rawDescOnce.Do(func() {
		file_another_service_name_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_another_service_name_proto_rawDesc), len(file_another_service_name_proto_rawDesc)))
	})
	return file_another_service_name_proto_rawDescData
}

var file_another_service_name_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_another_service_name_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_another_service_name_proto_goTypes = []any{
	(SomeEnum)(0),              // 0: another_service_name.v1.SomeEnum
	(*SomeMethodRequest)(nil),  // 1: another_service_name.v1.SomeMethodRequest
	(*SomeMethodResponse)(nil), // 2: another_service_name.v1.SomeMethodResponse
}
var file_another_service_name_proto_depIdxs = []int32{
	0, // 0: another_service_name.v1.SomeMethodRequest.value:type_name -> another_service_name.v1.SomeEnum
	1, // 1: another_service_name.v1.SomeService.SomeMethod:input_type -> another_service_name.v1.SomeMethodRequest
	2, // 2: another_service_name.v1.SomeService.SomeMethod:output_type -> another_service_name.v1.SomeMethodResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_another_service_name_proto_init() }
func file_another_service_name_proto_init() {
	if File_another_service_name_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_another_service_name_proto_rawDesc), len(file_another_service_name_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_another_service_name_proto_goTypes,
		DependencyIndexes: file_another_service_name_proto_depIdxs,
		EnumInfos:         file_another_service_name_proto_enumTypes,
		MessageInfos:      file_another_service_name_proto_msgTypes,
	}.Build()
	File_another_service_name_proto = out.File
	file_another_service_name_proto_goTypes = nil
	file_another_service_name_proto_depIdxs = nil
}
