// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: token.proto

package tokenpb

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

type GenTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	AppKey  int32  `protobuf:"varint,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
}

func (x *GenTokenReq) Reset() {
	*x = GenTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenTokenReq) ProtoMessage() {}

func (x *GenTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenTokenReq.ProtoReflect.Descriptor instead.
func (*GenTokenReq) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{0}
}

func (x *GenTokenReq) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *GenTokenReq) GetAppKey() int32 {
	if x != nil {
		return x.AppKey
	}
	return 0
}

type GenTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GenTokenResp) Reset() {
	*x = GenTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenTokenResp) ProtoMessage() {}

func (x *GenTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenTokenResp.ProtoReflect.Descriptor instead.
func (*GenTokenResp) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{1}
}

func (x *GenTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TokenDecodedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AppKey int32  `protobuf:"varint,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
}

func (x *TokenDecodedReq) Reset() {
	*x = TokenDecodedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenDecodedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenDecodedReq) ProtoMessage() {}

func (x *TokenDecodedReq) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenDecodedReq.ProtoReflect.Descriptor instead.
func (*TokenDecodedReq) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{2}
}

func (x *TokenDecodedReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenDecodedReq) GetAppKey() int32 {
	if x != nil {
		return x.AppKey
	}
	return 0
}

type TokenDecodedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	IsValid bool   `protobuf:"varint,2,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *TokenDecodedResp) Reset() {
	*x = TokenDecodedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenDecodedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenDecodedResp) ProtoMessage() {}

func (x *TokenDecodedResp) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenDecodedResp.ProtoReflect.Descriptor instead.
func (*TokenDecodedResp) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{3}
}

func (x *TokenDecodedResp) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *TokenDecodedResp) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_token_proto protoreflect.FileDescriptor

var file_token_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x70, 0x70, 0x4b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x0f, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x22, 0x46, 0x0a, 0x10,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x32, 0x82, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a,
	0x07, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_proto_rawDescOnce sync.Once
	file_token_proto_rawDescData = file_token_proto_rawDesc
)

func file_token_proto_rawDescGZIP() []byte {
	file_token_proto_rawDescOnce.Do(func() {
		file_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_proto_rawDescData)
	})
	return file_token_proto_rawDescData
}

var file_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_token_proto_goTypes = []interface{}{
	(*GenTokenReq)(nil),      // 0: token.GenTokenReq
	(*GenTokenResp)(nil),     // 1: token.GenTokenResp
	(*TokenDecodedReq)(nil),  // 2: token.TokenDecodedReq
	(*TokenDecodedResp)(nil), // 3: token.TokenDecodedResp
}
var file_token_proto_depIdxs = []int32{
	0, // 0: token.TokenService.getJwtToken:input_type -> token.GenTokenReq
	2, // 1: token.TokenService.decoded:input_type -> token.TokenDecodedReq
	1, // 2: token.TokenService.getJwtToken:output_type -> token.GenTokenResp
	3, // 3: token.TokenService.decoded:output_type -> token.TokenDecodedResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_token_proto_init() }
func file_token_proto_init() {
	if File_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenTokenReq); i {
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
		file_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenTokenResp); i {
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
		file_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenDecodedReq); i {
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
		file_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenDecodedResp); i {
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
			RawDescriptor: file_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_proto_goTypes,
		DependencyIndexes: file_token_proto_depIdxs,
		MessageInfos:      file_token_proto_msgTypes,
	}.Build()
	File_token_proto = out.File
	file_token_proto_rawDesc = nil
	file_token_proto_goTypes = nil
	file_token_proto_depIdxs = nil
}