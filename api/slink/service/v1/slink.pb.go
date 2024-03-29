// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: slink/service/v1/slink.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateShortLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link     string                 `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	ExpireAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
}

func (x *CreateShortLinkRequest) Reset() {
	*x = CreateShortLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slink_service_v1_slink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortLinkRequest) ProtoMessage() {}

func (x *CreateShortLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slink_service_v1_slink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateShortLinkRequest) Descriptor() ([]byte, []int) {
	return file_slink_service_v1_slink_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShortLinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CreateShortLinkRequest) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

type CreateShortLinkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ExpireAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
}

func (x *CreateShortLinkReply) Reset() {
	*x = CreateShortLinkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slink_service_v1_slink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortLinkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortLinkReply) ProtoMessage() {}

func (x *CreateShortLinkReply) ProtoReflect() protoreflect.Message {
	mi := &file_slink_service_v1_slink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortLinkReply.ProtoReflect.Descriptor instead.
func (*CreateShortLinkReply) Descriptor() ([]byte, []int) {
	return file_slink_service_v1_slink_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortLinkReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CreateShortLinkReply) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

var File_slink_service_v1_slink_proto protoreflect.FileDescriptor

var file_slink_service_v1_slink_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x4a, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x11, 0xfa, 0x42, 0x0e, 0xb2, 0x01, 0x0b, 0x08, 0x01, 0x40, 0x01, 0x4a, 0x05, 0x08, 0x80,
	0xbd, 0xa3, 0x01, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0x61, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74,
	0x32, 0x8e, 0x01, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x28, 0x2e, 0x73, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x42, 0x34, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x6f, 0x6d, 0x69, 0x6e, 0x73, 0x75, 0x2f, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slink_service_v1_slink_proto_rawDescOnce sync.Once
	file_slink_service_v1_slink_proto_rawDescData = file_slink_service_v1_slink_proto_rawDesc
)

func file_slink_service_v1_slink_proto_rawDescGZIP() []byte {
	file_slink_service_v1_slink_proto_rawDescOnce.Do(func() {
		file_slink_service_v1_slink_proto_rawDescData = protoimpl.X.CompressGZIP(file_slink_service_v1_slink_proto_rawDescData)
	})
	return file_slink_service_v1_slink_proto_rawDescData
}

var file_slink_service_v1_slink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_slink_service_v1_slink_proto_goTypes = []interface{}{
	(*CreateShortLinkRequest)(nil), // 0: slink.service.v1.CreateShortLinkRequest
	(*CreateShortLinkReply)(nil),   // 1: slink.service.v1.CreateShortLinkReply
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_slink_service_v1_slink_proto_depIdxs = []int32{
	2, // 0: slink.service.v1.CreateShortLinkRequest.expire_at:type_name -> google.protobuf.Timestamp
	2, // 1: slink.service.v1.CreateShortLinkReply.expire_at:type_name -> google.protobuf.Timestamp
	0, // 2: slink.service.v1.ShortLinkService.CreateShortLink:input_type -> slink.service.v1.CreateShortLinkRequest
	1, // 3: slink.service.v1.ShortLinkService.CreateShortLink:output_type -> slink.service.v1.CreateShortLinkReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_slink_service_v1_slink_proto_init() }
func file_slink_service_v1_slink_proto_init() {
	if File_slink_service_v1_slink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slink_service_v1_slink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortLinkRequest); i {
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
		file_slink_service_v1_slink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortLinkReply); i {
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
			RawDescriptor: file_slink_service_v1_slink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slink_service_v1_slink_proto_goTypes,
		DependencyIndexes: file_slink_service_v1_slink_proto_depIdxs,
		MessageInfos:      file_slink_service_v1_slink_proto_msgTypes,
	}.Build()
	File_slink_service_v1_slink_proto = out.File
	file_slink_service_v1_slink_proto_rawDesc = nil
	file_slink_service_v1_slink_proto_goTypes = nil
	file_slink_service_v1_slink_proto_depIdxs = nil
}
