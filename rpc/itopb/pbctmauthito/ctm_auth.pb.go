// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: ctm_auth.proto

package pbctmauthito

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

type GenCtmTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenId     string `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	UserId     int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	UserMobile string `protobuf:"bytes,3,opt,name=userMobile,proto3" json:"userMobile,omitempty"`
}

func (x *GenCtmTokenReq) Reset() {
	*x = GenCtmTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCtmTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCtmTokenReq) ProtoMessage() {}

func (x *GenCtmTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCtmTokenReq.ProtoReflect.Descriptor instead.
func (*GenCtmTokenReq) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GenCtmTokenReq) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *GenCtmTokenReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GenCtmTokenReq) GetUserMobile() string {
	if x != nil {
		return x.UserMobile
	}
	return ""
}

type GenCtmTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CtmToken string `protobuf:"bytes,1,opt,name=ctmToken,proto3" json:"ctmToken,omitempty"`
}

func (x *GenCtmTokenResp) Reset() {
	*x = GenCtmTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCtmTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCtmTokenResp) ProtoMessage() {}

func (x *GenCtmTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCtmTokenResp.ProtoReflect.Descriptor instead.
func (*GenCtmTokenResp) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GenCtmTokenResp) GetCtmToken() string {
	if x != nil {
		return x.CtmToken
	}
	return ""
}

type ParseCtmTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CtmToken string `protobuf:"bytes,1,opt,name=ctmToken,proto3" json:"ctmToken,omitempty"`
}

func (x *ParseCtmTokenReq) Reset() {
	*x = ParseCtmTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseCtmTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseCtmTokenReq) ProtoMessage() {}

func (x *ParseCtmTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseCtmTokenReq.ProtoReflect.Descriptor instead.
func (*ParseCtmTokenReq) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ParseCtmTokenReq) GetCtmToken() string {
	if x != nil {
		return x.CtmToken
	}
	return ""
}

type ParseCtmTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenId     string `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	UserId     int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	UserMobile string `protobuf:"bytes,3,opt,name=userMobile,proto3" json:"userMobile,omitempty"`
}

func (x *ParseCtmTokenResp) Reset() {
	*x = ParseCtmTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseCtmTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseCtmTokenResp) ProtoMessage() {}

func (x *ParseCtmTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseCtmTokenResp.ProtoReflect.Descriptor instead.
func (*ParseCtmTokenResp) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ParseCtmTokenResp) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *ParseCtmTokenResp) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ParseCtmTokenResp) GetUserMobile() string {
	if x != nil {
		return x.UserMobile
	}
	return ""
}

type CtmJsCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsCode string `protobuf:"bytes,3,opt,name=js_code,json=jsCode,proto3" json:"js_code,omitempty"`
}

func (x *CtmJsCodeReq) Reset() {
	*x = CtmJsCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtmJsCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtmJsCodeReq) ProtoMessage() {}

func (x *CtmJsCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtmJsCodeReq.ProtoReflect.Descriptor instead.
func (*CtmJsCodeReq) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{4}
}

func (x *CtmJsCodeReq) GetJsCode() string {
	if x != nil {
		return x.JsCode
	}
	return ""
}

type CtmJsCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionKey string `protobuf:"bytes,1,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
	UnionId    string `protobuf:"bytes,2,opt,name=unionId,proto3" json:"unionId,omitempty"`
	OpenId     string `protobuf:"bytes,3,opt,name=openId,proto3" json:"openId,omitempty"`
}

func (x *CtmJsCodeResp) Reset() {
	*x = CtmJsCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtmJsCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtmJsCodeResp) ProtoMessage() {}

func (x *CtmJsCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtmJsCodeResp.ProtoReflect.Descriptor instead.
func (*CtmJsCodeResp) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{5}
}

func (x *CtmJsCodeResp) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *CtmJsCodeResp) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

func (x *CtmJsCodeResp) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

type CtmPhoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CtmPhoneReq) Reset() {
	*x = CtmPhoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtmPhoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtmPhoneReq) ProtoMessage() {}

func (x *CtmPhoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtmPhoneReq.ProtoReflect.Descriptor instead.
func (*CtmPhoneReq) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{6}
}

func (x *CtmPhoneReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CtmPhoneResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber     string `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	PurePhoneNumber string `protobuf:"bytes,2,opt,name=purePhoneNumber,proto3" json:"purePhoneNumber,omitempty"`
	CountryCode     string `protobuf:"bytes,3,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
}

func (x *CtmPhoneResp) Reset() {
	*x = CtmPhoneResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctm_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtmPhoneResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtmPhoneResp) ProtoMessage() {}

func (x *CtmPhoneResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctm_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtmPhoneResp.ProtoReflect.Descriptor instead.
func (*CtmPhoneResp) Descriptor() ([]byte, []int) {
	return file_ctm_auth_proto_rawDescGZIP(), []int{7}
}

func (x *CtmPhoneResp) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CtmPhoneResp) GetPurePhoneNumber() string {
	if x != nil {
		return x.PurePhoneNumber
	}
	return ""
}

func (x *CtmPhoneResp) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

var File_ctm_auth_proto protoreflect.FileDescriptor

var file_ctm_auth_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x74, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x60, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x43, 0x74,
	0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x6e,
	0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x63, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x27, 0x0a,
	0x0c, 0x43, 0x74, 0x6d, 0x4a, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x6a, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6a, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x62, 0x0a, 0x0d, 0x43, 0x74, 0x6d, 0x4a, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x74,
	0x6d, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x7c, 0x0a,
	0x0c, 0x43, 0x74, 0x6d, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x0f, 0x70, 0x75, 0x72, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x75, 0x72, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xfb, 0x01, 0x0a, 0x0a,
	0x4d, 0x70, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x76, 0x63, 0x12, 0x37, 0x0a, 0x0a, 0x4d, 0x70,
	0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x43, 0x74, 0x6d, 0x4a, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x74, 0x6d, 0x4a, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x07, 0x4d, 0x70, 0x43, 0x74, 0x6d, 0x4d, 0x70, 0x12, 0x12,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x74, 0x6d, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x74, 0x6d, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x43, 0x74,
	0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x73, 0x65, 0x43, 0x74,
	0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x43, 0x74, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x43, 0x74, 0x6d,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70,
	0x62, 0x63, 0x74, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x69, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ctm_auth_proto_rawDescOnce sync.Once
	file_ctm_auth_proto_rawDescData = file_ctm_auth_proto_rawDesc
)

func file_ctm_auth_proto_rawDescGZIP() []byte {
	file_ctm_auth_proto_rawDescOnce.Do(func() {
		file_ctm_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_ctm_auth_proto_rawDescData)
	})
	return file_ctm_auth_proto_rawDescData
}

var file_ctm_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ctm_auth_proto_goTypes = []interface{}{
	(*GenCtmTokenReq)(nil),    // 0: login.GenCtmTokenReq
	(*GenCtmTokenResp)(nil),   // 1: login.GenCtmTokenResp
	(*ParseCtmTokenReq)(nil),  // 2: login.ParseCtmTokenReq
	(*ParseCtmTokenResp)(nil), // 3: login.ParseCtmTokenResp
	(*CtmJsCodeReq)(nil),      // 4: login.CtmJsCodeReq
	(*CtmJsCodeResp)(nil),     // 5: login.CtmJsCodeResp
	(*CtmPhoneReq)(nil),       // 6: login.CtmPhoneReq
	(*CtmPhoneResp)(nil),      // 7: login.CtmPhoneResp
}
var file_ctm_auth_proto_depIdxs = []int32{
	4, // 0: login.MpLoginSvc.MpCtmToken:input_type -> login.CtmJsCodeReq
	6, // 1: login.MpLoginSvc.MpCtmMp:input_type -> login.CtmPhoneReq
	0, // 2: login.MpLoginSvc.GenCtmToken:input_type -> login.GenCtmTokenReq
	2, // 3: login.MpLoginSvc.ParseCtmToken:input_type -> login.ParseCtmTokenReq
	5, // 4: login.MpLoginSvc.MpCtmToken:output_type -> login.CtmJsCodeResp
	7, // 5: login.MpLoginSvc.MpCtmMp:output_type -> login.CtmPhoneResp
	1, // 6: login.MpLoginSvc.GenCtmToken:output_type -> login.GenCtmTokenResp
	3, // 7: login.MpLoginSvc.ParseCtmToken:output_type -> login.ParseCtmTokenResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ctm_auth_proto_init() }
func file_ctm_auth_proto_init() {
	if File_ctm_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ctm_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCtmTokenReq); i {
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
		file_ctm_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCtmTokenResp); i {
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
		file_ctm_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseCtmTokenReq); i {
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
		file_ctm_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseCtmTokenResp); i {
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
		file_ctm_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtmJsCodeReq); i {
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
		file_ctm_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtmJsCodeResp); i {
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
		file_ctm_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtmPhoneReq); i {
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
		file_ctm_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtmPhoneResp); i {
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
			RawDescriptor: file_ctm_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ctm_auth_proto_goTypes,
		DependencyIndexes: file_ctm_auth_proto_depIdxs,
		MessageInfos:      file_ctm_auth_proto_msgTypes,
	}.Build()
	File_ctm_auth_proto = out.File
	file_ctm_auth_proto_rawDesc = nil
	file_ctm_auth_proto_goTypes = nil
	file_ctm_auth_proto_depIdxs = nil
}