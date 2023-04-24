// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: link-shortener.proto

package protos

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

type GetRawURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GetRawURLRequest) Reset() {
	*x = GetRawURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawURLRequest) ProtoMessage() {}

func (x *GetRawURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawURLRequest.ProtoReflect.Descriptor instead.
func (*GetRawURLRequest) Descriptor() ([]byte, []int) {
	return file_link_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *GetRawURLRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetRawURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawURL string `protobuf:"bytes,1,opt,name=RawURL,proto3" json:"RawURL,omitempty"`
}

func (x *GetRawURLResponse) Reset() {
	*x = GetRawURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawURLResponse) ProtoMessage() {}

func (x *GetRawURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawURLResponse.ProtoReflect.Descriptor instead.
func (*GetRawURLResponse) Descriptor() ([]byte, []int) {
	return file_link_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetRawURLResponse) GetRawURL() string {
	if x != nil {
		return x.RawURL
	}
	return ""
}

type ShortenURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawURL string `protobuf:"bytes,1,opt,name=RawURL,proto3" json:"RawURL,omitempty"`
}

func (x *ShortenURLRequest) Reset() {
	*x = ShortenURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLRequest) ProtoMessage() {}

func (x *ShortenURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenURLRequest) Descriptor() ([]byte, []int) {
	return file_link_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenURLRequest) GetRawURL() string {
	if x != nil {
		return x.RawURL
	}
	return ""
}

type ShortenURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *ShortenURLResponse) Reset() {
	*x = ShortenURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLResponse) ProtoMessage() {}

func (x *ShortenURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLResponse.ProtoReflect.Descriptor instead.
func (*ShortenURLResponse) Descriptor() ([]byte, []int) {
	return file_link_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenURLResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_link_shortener_proto protoreflect.FileDescriptor

var file_link_shortener_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x28,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x77, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x61, 0x77, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x61, 0x77, 0x55, 0x52, 0x4c, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61,
	0x77, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x61, 0x77, 0x55,
	0x52, 0x4c, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x96,
	0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x6c, 0x69, 0x6e, 0x6b, 0x2d,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_shortener_proto_rawDescOnce sync.Once
	file_link_shortener_proto_rawDescData = file_link_shortener_proto_rawDesc
)

func file_link_shortener_proto_rawDescGZIP() []byte {
	file_link_shortener_proto_rawDescOnce.Do(func() {
		file_link_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_shortener_proto_rawDescData)
	})
	return file_link_shortener_proto_rawDescData
}

var file_link_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_link_shortener_proto_goTypes = []interface{}{
	(*GetRawURLRequest)(nil),   // 0: protos.GetRawURLRequest
	(*GetRawURLResponse)(nil),  // 1: protos.GetRawURLResponse
	(*ShortenURLRequest)(nil),  // 2: protos.ShortenURLRequest
	(*ShortenURLResponse)(nil), // 3: protos.ShortenURLResponse
}
var file_link_shortener_proto_depIdxs = []int32{
	0, // 0: protos.LinkShortener.GetRawURL:input_type -> protos.GetRawURLRequest
	2, // 1: protos.LinkShortener.ShortenURL:input_type -> protos.ShortenURLRequest
	1, // 2: protos.LinkShortener.GetRawURL:output_type -> protos.GetRawURLResponse
	3, // 3: protos.LinkShortener.ShortenURL:output_type -> protos.ShortenURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_link_shortener_proto_init() }
func file_link_shortener_proto_init() {
	if File_link_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_link_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawURLRequest); i {
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
		file_link_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawURLResponse); i {
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
		file_link_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLRequest); i {
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
		file_link_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLResponse); i {
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
			RawDescriptor: file_link_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_shortener_proto_goTypes,
		DependencyIndexes: file_link_shortener_proto_depIdxs,
		MessageInfos:      file_link_shortener_proto_msgTypes,
	}.Build()
	File_link_shortener_proto = out.File
	file_link_shortener_proto_rawDesc = nil
	file_link_shortener_proto_goTypes = nil
	file_link_shortener_proto_depIdxs = nil
}