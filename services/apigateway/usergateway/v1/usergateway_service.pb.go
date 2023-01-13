// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: usergateway/v1/usergateway_service.proto

package usergatewayv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ChangeAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Avatar string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *ChangeAvatarRequest) Reset() {
	*x = ChangeAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAvatarRequest) ProtoMessage() {}

func (x *ChangeAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAvatarRequest.ProtoReflect.Descriptor instead.
func (*ChangeAvatarRequest) Descriptor() ([]byte, []int) {
	return file_usergateway_v1_usergateway_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeAvatarRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChangeAvatarRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type ChangeAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error      string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Avatarlink string `protobuf:"bytes,3,opt,name=avatarlink,proto3" json:"avatarlink,omitempty"`
}

func (x *ChangeAvatarResponse) Reset() {
	*x = ChangeAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAvatarResponse) ProtoMessage() {}

func (x *ChangeAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAvatarResponse.ProtoReflect.Descriptor instead.
func (*ChangeAvatarResponse) Descriptor() ([]byte, []int) {
	return file_usergateway_v1_usergateway_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeAvatarResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ChangeAvatarResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ChangeAvatarResponse) GetAvatarlink() string {
	if x != nil {
		return x.Avatarlink
	}
	return ""
}

type ChangeUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ChangeUsernameRequest) Reset() {
	*x = ChangeUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUsernameRequest) ProtoMessage() {}

func (x *ChangeUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUsernameRequest.ProtoReflect.Descriptor instead.
func (*ChangeUsernameRequest) Descriptor() ([]byte, []int) {
	return file_usergateway_v1_usergateway_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeUsernameRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChangeUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ChangeUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ChangeUsernameResponse) Reset() {
	*x = ChangeUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUsernameResponse) ProtoMessage() {}

func (x *ChangeUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUsernameResponse.ProtoReflect.Descriptor instead.
func (*ChangeUsernameResponse) Descriptor() ([]byte, []int) {
	return file_usergateway_v1_usergateway_service_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeUsernameResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ChangeUsernameResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ChangeUsernameResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_usergateway_v1_usergateway_service_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePasswordRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChangePasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ChangePasswordResponse) Reset() {
	*x = ChangePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordResponse) ProtoMessage() {}

func (x *ChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usergateway_v1_usergateway_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_usergateway_v1_usergateway_service_proto_rawDescGZIP(), []int{5}
}

func (x *ChangePasswordResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ChangePasswordResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_usergateway_v1_usergateway_service_proto protoreflect.FileDescriptor

var file_usergateway_v1_usergateway_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x75, 0x73, 0x65, 0x72,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x22, 0x64, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x4c, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb1,
	0x05, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd4, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x79, 0x92, 0x41, 0x56, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x20,
	0x6f, 0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x31, 0x47, 0x52, 0x50, 0x43, 0x20,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0xe0, 0x01, 0x0a,
	0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f,
	0x92, 0x41, 0x5a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f,
	0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x33, 0x47, 0x52, 0x50, 0x43, 0x20, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0xe0, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x7f, 0x92, 0x41, 0x5a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x33, 0x47, 0x52, 0x50,
	0x43, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x20, 0x74,
	0x6f, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0xb1, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x65, 0x74, 0x7a, 0x79, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70,
	0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x76, 0x31, 0x92, 0x41, 0x61, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x72,
	0x55, 0x0a, 0x21, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x20, 0x47, 0x52, 0x50, 0x43, 0x20,
	0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x20, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x65, 0x74, 0x7a, 0x79, 0x31,
	0x2f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x43, 0x68,
	0x61, 0x74, 0x2d, 0x41, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usergateway_v1_usergateway_service_proto_rawDescOnce sync.Once
	file_usergateway_v1_usergateway_service_proto_rawDescData = file_usergateway_v1_usergateway_service_proto_rawDesc
)

func file_usergateway_v1_usergateway_service_proto_rawDescGZIP() []byte {
	file_usergateway_v1_usergateway_service_proto_rawDescOnce.Do(func() {
		file_usergateway_v1_usergateway_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_usergateway_v1_usergateway_service_proto_rawDescData)
	})
	return file_usergateway_v1_usergateway_service_proto_rawDescData
}

var file_usergateway_v1_usergateway_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_usergateway_v1_usergateway_service_proto_goTypes = []interface{}{
	(*ChangeAvatarRequest)(nil),    // 0: usergateway.v1.ChangeAvatarRequest
	(*ChangeAvatarResponse)(nil),   // 1: usergateway.v1.ChangeAvatarResponse
	(*ChangeUsernameRequest)(nil),  // 2: usergateway.v1.ChangeUsernameRequest
	(*ChangeUsernameResponse)(nil), // 3: usergateway.v1.ChangeUsernameResponse
	(*ChangePasswordRequest)(nil),  // 4: usergateway.v1.ChangePasswordRequest
	(*ChangePasswordResponse)(nil), // 5: usergateway.v1.ChangePasswordResponse
}
var file_usergateway_v1_usergateway_service_proto_depIdxs = []int32{
	0, // 0: usergateway.v1.UserGatewayService.ChangeAvatar:input_type -> usergateway.v1.ChangeAvatarRequest
	2, // 1: usergateway.v1.UserGatewayService.ChangeUsername:input_type -> usergateway.v1.ChangeUsernameRequest
	4, // 2: usergateway.v1.UserGatewayService.ChangePassword:input_type -> usergateway.v1.ChangePasswordRequest
	1, // 3: usergateway.v1.UserGatewayService.ChangeAvatar:output_type -> usergateway.v1.ChangeAvatarResponse
	3, // 4: usergateway.v1.UserGatewayService.ChangeUsername:output_type -> usergateway.v1.ChangeUsernameResponse
	5, // 5: usergateway.v1.UserGatewayService.ChangePassword:output_type -> usergateway.v1.ChangePasswordResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_usergateway_v1_usergateway_service_proto_init() }
func file_usergateway_v1_usergateway_service_proto_init() {
	if File_usergateway_v1_usergateway_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usergateway_v1_usergateway_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAvatarRequest); i {
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
		file_usergateway_v1_usergateway_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAvatarResponse); i {
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
		file_usergateway_v1_usergateway_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUsernameRequest); i {
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
		file_usergateway_v1_usergateway_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUsernameResponse); i {
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
		file_usergateway_v1_usergateway_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_usergateway_v1_usergateway_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordResponse); i {
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
			RawDescriptor: file_usergateway_v1_usergateway_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usergateway_v1_usergateway_service_proto_goTypes,
		DependencyIndexes: file_usergateway_v1_usergateway_service_proto_depIdxs,
		MessageInfos:      file_usergateway_v1_usergateway_service_proto_msgTypes,
	}.Build()
	File_usergateway_v1_usergateway_service_proto = out.File
	file_usergateway_v1_usergateway_service_proto_rawDesc = nil
	file_usergateway_v1_usergateway_service_proto_goTypes = nil
	file_usergateway_v1_usergateway_service_proto_depIdxs = nil
}
