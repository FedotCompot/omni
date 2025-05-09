// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: omni/specs/ephemeral.proto

package specs

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

// MachineStatusLinkSpec describes the combination of MessageStatusSpec and SideroLinkSpec and SiderolinkCounterSpec
type MachineStatusLinkSpec struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	MessageStatus     *MachineStatusSpec     `protobuf:"bytes,1,opt,name=message_status,json=messageStatus,proto3" json:"message_status,omitempty"`
	SiderolinkCounter *SiderolinkCounterSpec `protobuf:"bytes,2,opt,name=siderolink_counter,json=siderolinkCounter,proto3" json:"siderolink_counter,omitempty"`
	MachineCreatedAt  int64                  `protobuf:"varint,3,opt,name=machine_created_at,json=machineCreatedAt,proto3" json:"machine_created_at,omitempty"`
	TearingDown       bool                   `protobuf:"varint,4,opt,name=tearing_down,json=tearingDown,proto3" json:"tearing_down,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *MachineStatusLinkSpec) Reset() {
	*x = MachineStatusLinkSpec{}
	mi := &file_omni_specs_ephemeral_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineStatusLinkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineStatusLinkSpec) ProtoMessage() {}

func (x *MachineStatusLinkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_ephemeral_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineStatusLinkSpec.ProtoReflect.Descriptor instead.
func (*MachineStatusLinkSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_ephemeral_proto_rawDescGZIP(), []int{0}
}

func (x *MachineStatusLinkSpec) GetMessageStatus() *MachineStatusSpec {
	if x != nil {
		return x.MessageStatus
	}
	return nil
}

func (x *MachineStatusLinkSpec) GetSiderolinkCounter() *SiderolinkCounterSpec {
	if x != nil {
		return x.SiderolinkCounter
	}
	return nil
}

func (x *MachineStatusLinkSpec) GetMachineCreatedAt() int64 {
	if x != nil {
		return x.MachineCreatedAt
	}
	return 0
}

func (x *MachineStatusLinkSpec) GetTearingDown() bool {
	if x != nil {
		return x.TearingDown
	}
	return false
}

var File_omni_specs_ephemeral_proto protoreflect.FileDescriptor

var file_omni_specs_ephemeral_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x65, 0x70, 0x68,
	0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x1a, 0x15, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f,
	0x6f, 0x6d, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x6d, 0x6e, 0x69,
	0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x15, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x3f, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x4b, 0x0a, 0x12, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x11, 0x73, 0x69,
	0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x77, 0x6e,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_omni_specs_ephemeral_proto_rawDescOnce sync.Once
	file_omni_specs_ephemeral_proto_rawDescData []byte
)

func file_omni_specs_ephemeral_proto_rawDescGZIP() []byte {
	file_omni_specs_ephemeral_proto_rawDescOnce.Do(func() {
		file_omni_specs_ephemeral_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_omni_specs_ephemeral_proto_rawDesc), len(file_omni_specs_ephemeral_proto_rawDesc)))
	})
	return file_omni_specs_ephemeral_proto_rawDescData
}

var file_omni_specs_ephemeral_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_omni_specs_ephemeral_proto_goTypes = []any{
	(*MachineStatusLinkSpec)(nil), // 0: specs.MachineStatusLinkSpec
	(*MachineStatusSpec)(nil),     // 1: specs.MachineStatusSpec
	(*SiderolinkCounterSpec)(nil), // 2: specs.SiderolinkCounterSpec
}
var file_omni_specs_ephemeral_proto_depIdxs = []int32{
	1, // 0: specs.MachineStatusLinkSpec.message_status:type_name -> specs.MachineStatusSpec
	2, // 1: specs.MachineStatusLinkSpec.siderolink_counter:type_name -> specs.SiderolinkCounterSpec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_omni_specs_ephemeral_proto_init() }
func file_omni_specs_ephemeral_proto_init() {
	if File_omni_specs_ephemeral_proto != nil {
		return
	}
	file_omni_specs_omni_proto_init()
	file_omni_specs_siderolink_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_omni_specs_ephemeral_proto_rawDesc), len(file_omni_specs_ephemeral_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_omni_specs_ephemeral_proto_goTypes,
		DependencyIndexes: file_omni_specs_ephemeral_proto_depIdxs,
		MessageInfos:      file_omni_specs_ephemeral_proto_msgTypes,
	}.Build()
	File_omni_specs_ephemeral_proto = out.File
	file_omni_specs_ephemeral_proto_goTypes = nil
	file_omni_specs_ephemeral_proto_depIdxs = nil
}
