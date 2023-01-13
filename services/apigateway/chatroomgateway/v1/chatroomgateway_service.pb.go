// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: chatroomgateway/v1/chatroomgateway_service.proto

package chatroomgatewayv1

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

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnerUuid string `protobuf:"bytes,2,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoomRequest) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error        string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ChatroomUuid string `protobuf:"bytes,3,opt,name=chatroom_uuid,json=chatroomUuid,proto3" json:"chatroom_uuid,omitempty"`
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateRoomResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateRoomResponse) GetChatroomUuid() string {
	if x != nil {
		return x.ChatroomUuid
	}
	return ""
}

type DeleteRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomUuid string `protobuf:"bytes,1,opt,name=chatroom_uuid,json=chatroomUuid,proto3" json:"chatroom_uuid,omitempty"`
	OwnerUuid    string `protobuf:"bytes,2,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
}

func (x *DeleteRoomRequest) Reset() {
	*x = DeleteRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomRequest) ProtoMessage() {}

func (x *DeleteRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRoomRequest) GetChatroomUuid() string {
	if x != nil {
		return x.ChatroomUuid
	}
	return ""
}

func (x *DeleteRoomRequest) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

type DeleteRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteRoomResponse) Reset() {
	*x = DeleteRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomResponse) ProtoMessage() {}

func (x *DeleteRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRoomResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeleteRoomResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomUuid string `protobuf:"bytes,1,opt,name=chatroom_uuid,json=chatroomUuid,proto3" json:"chatroom_uuid,omitempty"`
	UserUuid     string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRoomRequest) GetChatroomUuid() string {
	if x != nil {
		return x.ChatroomUuid
	}
	return ""
}

func (x *JoinRoomRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type JoinRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *JoinRoomResponse) Reset() {
	*x = JoinRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomResponse) ProtoMessage() {}

func (x *JoinRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomResponse.ProtoReflect.Descriptor instead.
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{5}
}

func (x *JoinRoomResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JoinRoomResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LeaveRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomUuid string `protobuf:"bytes,1,opt,name=chatroom_uuid,json=chatroomUuid,proto3" json:"chatroom_uuid,omitempty"`
	UserUuid     string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *LeaveRoomRequest) Reset() {
	*x = LeaveRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomRequest) ProtoMessage() {}

func (x *LeaveRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomRequest.ProtoReflect.Descriptor instead.
func (*LeaveRoomRequest) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{6}
}

func (x *LeaveRoomRequest) GetChatroomUuid() string {
	if x != nil {
		return x.ChatroomUuid
	}
	return ""
}

func (x *LeaveRoomRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type LeaveRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LeaveRoomResponse) Reset() {
	*x = LeaveRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomResponse) ProtoMessage() {}

func (x *LeaveRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomResponse.ProtoReflect.Descriptor instead.
func (*LeaveRoomResponse) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{7}
}

func (x *LeaveRoomResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LeaveRoomResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomUuid string `protobuf:"bytes,1,opt,name=chatroom_uuid,json=chatroomUuid,proto3" json:"chatroom_uuid,omitempty"`
}

func (x *GetRoomRequest) Reset() {
	*x = GetRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomRequest) ProtoMessage() {}

func (x *GetRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomRequest.ProtoReflect.Descriptor instead.
func (*GetRoomRequest) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetRoomRequest) GetChatroomUuid() string {
	if x != nil {
		return x.ChatroomUuid
	}
	return ""
}

type GetRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error     string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	OwnerUuid string   `protobuf:"bytes,4,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	Users     []string `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetRoomResponse) Reset() {
	*x = GetRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomResponse) ProtoMessage() {}

func (x *GetRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomResponse.ProtoReflect.Descriptor instead.
func (*GetRoomResponse) Descriptor() ([]byte, []int) {
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetRoomResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetRoomResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetRoomResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetRoomResponse) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

func (x *GetRoomResponse) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_chatroomgateway_v1_chatroomgateway_service_proto protoreflect.FileDescriptor

var file_chatroomgateway_v1_chatroomgateway_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x55, 0x75, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x42,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x53, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68,
	0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x10, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x11, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x35, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x32, 0xf8, 0x07, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xc8, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x25,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92,
	0x41, 0x46, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x1a,
	0x27, 0x47, 0x52, 0x50, 0x43, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61,
	0x6c, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0xc8, 0x01, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92, 0x41, 0x46, 0x0a, 0x08, 0x43,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x27, 0x47, 0x52, 0x50, 0x43,
	0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f,
	0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0xbc, 0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f,
	0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x92,
	0x41, 0x42, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x0f, 0x4a, 0x6f,
	0x69, 0x6e, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x25, 0x47,
	0x52, 0x50, 0x43, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61, 0x6c, 0x6c,
	0x20, 0x74, 0x6f, 0x20, 0x6a, 0x6f, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x69, 0x6e,
	0x72, 0x6f, 0x6f, 0x6d, 0x12, 0xc2, 0x01, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72,
	0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x68, 0x92, 0x41, 0x44, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x10,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x1a, 0x26, 0x47, 0x52, 0x50, 0x43, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63,
	0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x20, 0x61, 0x20, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01,
	0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x6c, 0x65, 0x61, 0x76, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0xc3, 0x01, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6f,
	0x92, 0x41, 0x40, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x47,
	0x65, 0x74, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x24, 0x47,
	0x52, 0x50, 0x43, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x63, 0x61, 0x6c, 0x6c,
	0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x7b, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x42,
	0xb9, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x65, 0x74, 0x7a, 0x79, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x76, 0x31, 0x92, 0x41, 0x61, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x2a, 0x01, 0x02, 0x72, 0x55, 0x0a, 0x21, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x20,
	0x47, 0x52, 0x50, 0x43, 0x20, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x20, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x65, 0x74, 0x7a, 0x79, 0x31, 0x2f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x43, 0x68, 0x61, 0x74, 0x2d, 0x41, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescOnce sync.Once
	file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescData = file_chatroomgateway_v1_chatroomgateway_service_proto_rawDesc
)

func file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescGZIP() []byte {
	file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescOnce.Do(func() {
		file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescData)
	})
	return file_chatroomgateway_v1_chatroomgateway_service_proto_rawDescData
}

var file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chatroomgateway_v1_chatroomgateway_service_proto_goTypes = []interface{}{
	(*CreateRoomRequest)(nil),  // 0: chatroomgateway.v1.CreateRoomRequest
	(*CreateRoomResponse)(nil), // 1: chatroomgateway.v1.CreateRoomResponse
	(*DeleteRoomRequest)(nil),  // 2: chatroomgateway.v1.DeleteRoomRequest
	(*DeleteRoomResponse)(nil), // 3: chatroomgateway.v1.DeleteRoomResponse
	(*JoinRoomRequest)(nil),    // 4: chatroomgateway.v1.JoinRoomRequest
	(*JoinRoomResponse)(nil),   // 5: chatroomgateway.v1.JoinRoomResponse
	(*LeaveRoomRequest)(nil),   // 6: chatroomgateway.v1.LeaveRoomRequest
	(*LeaveRoomResponse)(nil),  // 7: chatroomgateway.v1.LeaveRoomResponse
	(*GetRoomRequest)(nil),     // 8: chatroomgateway.v1.GetRoomRequest
	(*GetRoomResponse)(nil),    // 9: chatroomgateway.v1.GetRoomResponse
}
var file_chatroomgateway_v1_chatroomgateway_service_proto_depIdxs = []int32{
	0, // 0: chatroomgateway.v1.ChatroomGatewayService.CreateRoom:input_type -> chatroomgateway.v1.CreateRoomRequest
	2, // 1: chatroomgateway.v1.ChatroomGatewayService.DeleteRoom:input_type -> chatroomgateway.v1.DeleteRoomRequest
	4, // 2: chatroomgateway.v1.ChatroomGatewayService.JoinRoom:input_type -> chatroomgateway.v1.JoinRoomRequest
	6, // 3: chatroomgateway.v1.ChatroomGatewayService.LeaveRoom:input_type -> chatroomgateway.v1.LeaveRoomRequest
	8, // 4: chatroomgateway.v1.ChatroomGatewayService.GetRoom:input_type -> chatroomgateway.v1.GetRoomRequest
	1, // 5: chatroomgateway.v1.ChatroomGatewayService.CreateRoom:output_type -> chatroomgateway.v1.CreateRoomResponse
	3, // 6: chatroomgateway.v1.ChatroomGatewayService.DeleteRoom:output_type -> chatroomgateway.v1.DeleteRoomResponse
	5, // 7: chatroomgateway.v1.ChatroomGatewayService.JoinRoom:output_type -> chatroomgateway.v1.JoinRoomResponse
	7, // 8: chatroomgateway.v1.ChatroomGatewayService.LeaveRoom:output_type -> chatroomgateway.v1.LeaveRoomResponse
	9, // 9: chatroomgateway.v1.ChatroomGatewayService.GetRoom:output_type -> chatroomgateway.v1.GetRoomResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatroomgateway_v1_chatroomgateway_service_proto_init() }
func file_chatroomgateway_v1_chatroomgateway_service_proto_init() {
	if File_chatroomgateway_v1_chatroomgateway_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomResponse); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomRequest); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomResponse); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomRequest); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomResponse); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRoomRequest); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRoomResponse); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomRequest); i {
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
		file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomResponse); i {
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
			RawDescriptor: file_chatroomgateway_v1_chatroomgateway_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatroomgateway_v1_chatroomgateway_service_proto_goTypes,
		DependencyIndexes: file_chatroomgateway_v1_chatroomgateway_service_proto_depIdxs,
		MessageInfos:      file_chatroomgateway_v1_chatroomgateway_service_proto_msgTypes,
	}.Build()
	File_chatroomgateway_v1_chatroomgateway_service_proto = out.File
	file_chatroomgateway_v1_chatroomgateway_service_proto_rawDesc = nil
	file_chatroomgateway_v1_chatroomgateway_service_proto_goTypes = nil
	file_chatroomgateway_v1_chatroomgateway_service_proto_depIdxs = nil
}
