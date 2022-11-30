// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: command.proto

package models

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

type CommandType int32

const (
	CommandType_AddItem     CommandType = 0
	CommandType_GetItem     CommandType = 1
	CommandType_GetAllItems CommandType = 2
	CommandType_RemoveItem  CommandType = 3
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "AddItem",
		1: "GetItem",
		2: "GetAllItems",
		3: "RemoveItem",
	}
	CommandType_value = map[string]int32{
		"AddItem":     0,
		"GetItem":     1,
		"GetAllItems": 2,
		"RemoveItem":  3,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_command_proto_enumTypes[0].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_command_proto_enumTypes[0]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        CommandType `protobuf:"varint,1,opt,name=type,proto3,enum=CommandType" json:"type,omitempty"`
	ItemID      int64       `protobuf:"varint,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	ItemPayload string      `protobuf:"bytes,3,opt,name=ItemPayload,proto3" json:"ItemPayload,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_AddItem
}

func (x *Command) GetItemID() int64 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *Command) GetItemPayload() string {
	if x != nil {
		return x.ItemPayload
	}
	return ""
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x48, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x03,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x61, 0x6e, 0x69, 0x65, 0x6c, 0x2d, 0x38, 0x38, 0x38, 0x2f, 0x42, 0x6c, 0x6f, 0x58, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2d, 0x4d, 0x51, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_command_proto_goTypes = []interface{}{
	(CommandType)(0), // 0: CommandType
	(*Command)(nil),  // 1: Command
}
var file_command_proto_depIdxs = []int32{
	0, // 0: Command.type:type_name -> CommandType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		EnumInfos:         file_command_proto_enumTypes,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
