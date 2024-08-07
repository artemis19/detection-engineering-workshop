// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: enrichment.proto

package enrichmentpb

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

type Host_Platform int32

const (
	Host_UNSPECIFIED Host_Platform = 0
	Host_LINUX       Host_Platform = 1
	Host_WINDOWS     Host_Platform = 2
	Host_MAC         Host_Platform = 3
)

// Enum value maps for Host_Platform.
var (
	Host_Platform_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "LINUX",
		2: "WINDOWS",
		3: "MAC",
	}
	Host_Platform_value = map[string]int32{
		"UNSPECIFIED": 0,
		"LINUX":       1,
		"WINDOWS":     2,
		"MAC":         3,
	}
)

func (x Host_Platform) Enum() *Host_Platform {
	p := new(Host_Platform)
	*p = x
	return p
}

func (x Host_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Host_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_enrichment_proto_enumTypes[0].Descriptor()
}

func (Host_Platform) Type() protoreflect.EnumType {
	return &file_enrichment_proto_enumTypes[0]
}

func (x Host_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Host_Platform.Descriptor instead.
func (Host_Platform) EnumDescriptor() ([]byte, []int) {
	return file_enrichment_proto_rawDescGZIP(), []int{1, 0}
}

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enrichment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_enrichment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_enrichment_proto_rawDescGZIP(), []int{0}
}

func (x *IP) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Platform Host_Platform `protobuf:"varint,2,opt,name=platform,proto3,enum=enrichment.Host_Platform" json:"platform,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enrichment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_enrichment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_enrichment_proto_rawDescGZIP(), []int{1}
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetPlatform() Host_Platform {
	if x != nil {
		return x.Platform
	}
	return Host_UNSPECIFIED
}

var File_enrichment_proto protoreflect.FileDescriptor

var file_enrichment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x14,
	0x0a, 0x02, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x22, 0x8f, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x3c, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x4d, 0x41, 0x43, 0x10, 0x03, 0x32, 0x6c, 0x0a, 0x0a, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x49, 0x50, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x0e, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x50,
	0x1a, 0x10, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x49, 0x50,
	0x12, 0x10, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x49, 0x50, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2f, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x65, 0x4e, 0x65, 0x65, 0x64, 0x6c,
	0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_enrichment_proto_rawDescOnce sync.Once
	file_enrichment_proto_rawDescData = file_enrichment_proto_rawDesc
)

func file_enrichment_proto_rawDescGZIP() []byte {
	file_enrichment_proto_rawDescOnce.Do(func() {
		file_enrichment_proto_rawDescData = protoimpl.X.CompressGZIP(file_enrichment_proto_rawDescData)
	})
	return file_enrichment_proto_rawDescData
}

var file_enrichment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enrichment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_enrichment_proto_goTypes = []interface{}{
	(Host_Platform)(0), // 0: enrichment.Host.Platform
	(*IP)(nil),         // 1: enrichment.IP
	(*Host)(nil),       // 2: enrichment.Host
}
var file_enrichment_proto_depIdxs = []int32{
	0, // 0: enrichment.Host.platform:type_name -> enrichment.Host.Platform
	1, // 1: enrichment.Enrichment.IPToHost:input_type -> enrichment.IP
	2, // 2: enrichment.Enrichment.HostToIP:input_type -> enrichment.Host
	2, // 3: enrichment.Enrichment.IPToHost:output_type -> enrichment.Host
	1, // 4: enrichment.Enrichment.HostToIP:output_type -> enrichment.IP
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enrichment_proto_init() }
func file_enrichment_proto_init() {
	if File_enrichment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enrichment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
		file_enrichment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
			RawDescriptor: file_enrichment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_enrichment_proto_goTypes,
		DependencyIndexes: file_enrichment_proto_depIdxs,
		EnumInfos:         file_enrichment_proto_enumTypes,
		MessageInfos:      file_enrichment_proto_msgTypes,
	}.Build()
	File_enrichment_proto = out.File
	file_enrichment_proto_rawDesc = nil
	file_enrichment_proto_goTypes = nil
	file_enrichment_proto_depIdxs = nil
}
