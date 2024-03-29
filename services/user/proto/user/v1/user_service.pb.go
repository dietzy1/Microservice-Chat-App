// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: user/v1/user_service.proto

package userv1

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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{1}
}

type AddChatServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid       string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	ChatserverUuid string `protobuf:"bytes,2,opt,name=chatserver_uuid,json=chatserverUuid,proto3" json:"chatserver_uuid,omitempty"`
}

func (x *AddChatServerRequest) Reset() {
	*x = AddChatServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChatServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChatServerRequest) ProtoMessage() {}

func (x *AddChatServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChatServerRequest.ProtoReflect.Descriptor instead.
func (*AddChatServerRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddChatServerRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *AddChatServerRequest) GetChatserverUuid() string {
	if x != nil {
		return x.ChatserverUuid
	}
	return ""
}

type AddChatServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddChatServerResponse) Reset() {
	*x = AddChatServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChatServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChatServerResponse) ProtoMessage() {}

func (x *AddChatServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChatServerResponse.ProtoReflect.Descriptor instead.
func (*AddChatServerResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{3}
}

type RemoveChatServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid       string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	ChatserverUuid string `protobuf:"bytes,2,opt,name=chatserver_uuid,json=chatserverUuid,proto3" json:"chatserver_uuid,omitempty"`
}

func (x *RemoveChatServerRequest) Reset() {
	*x = RemoveChatServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveChatServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveChatServerRequest) ProtoMessage() {}

func (x *RemoveChatServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveChatServerRequest.ProtoReflect.Descriptor instead.
func (*RemoveChatServerRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveChatServerRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *RemoveChatServerRequest) GetChatserverUuid() string {
	if x != nil {
		return x.ChatserverUuid
	}
	return ""
}

type RemoveChatServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveChatServerResponse) Reset() {
	*x = RemoveChatServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveChatServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveChatServerResponse) ProtoMessage() {}

func (x *RemoveChatServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveChatServerResponse.ProtoReflect.Descriptor instead.
func (*RemoveChatServerResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{5}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid        string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Icon        *Icon    `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	JoinDate    string   `protobuf:"bytes,5,opt,name=join_date,json=joinDate,proto3" json:"join_date,omitempty"`
	ChatServers []string `protobuf:"bytes,6,rep,name=chat_servers,json=chatServers,proto3" json:"chat_servers,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetUserResponse) GetIcon() *Icon {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *GetUserResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetUserResponse) GetJoinDate() string {
	if x != nil {
		return x.JoinDate
	}
	return ""
}

func (x *GetUserResponse) GetChatServers() []string {
	if x != nil {
		return x.ChatServers
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuids []string `protobuf:"bytes,1,rep,name=user_uuids,json=userUuids,proto3" json:"user_uuids,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetUsersRequest) GetUserUuids() []string {
	if x != nil {
		return x.UserUuids
	}
	return nil
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*GetUserResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetUsersResponse) GetUsers() []*GetUserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type EditDescriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid    string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EditDescriptionRequest) Reset() {
	*x = EditDescriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditDescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDescriptionRequest) ProtoMessage() {}

func (x *EditDescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDescriptionRequest.ProtoReflect.Descriptor instead.
func (*EditDescriptionRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *EditDescriptionRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *EditDescriptionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EditDescriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditDescriptionResponse) Reset() {
	*x = EditDescriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditDescriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDescriptionResponse) ProtoMessage() {}

func (x *EditDescriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDescriptionResponse.ProtoReflect.Descriptor instead.
func (*EditDescriptionResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{11}
}

type ChangeAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	IconUuid string `protobuf:"bytes,2,opt,name=icon_uuid,json=iconUuid,proto3" json:"icon_uuid,omitempty"`
}

func (x *ChangeAvatarRequest) Reset() {
	*x = ChangeAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAvatarRequest) ProtoMessage() {}

func (x *ChangeAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[12]
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
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{12}
}

func (x *ChangeAvatarRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *ChangeAvatarRequest) GetIconUuid() string {
	if x != nil {
		return x.IconUuid
	}
	return ""
}

type ChangeAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeAvatarResponse) Reset() {
	*x = ChangeAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAvatarResponse) ProtoMessage() {}

func (x *ChangeAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[13]
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
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{13}
}

type Icon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Icon) Reset() {
	*x = Icon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_service_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Icon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Icon) ProtoMessage() {}

func (x *Icon) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_service_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Icon.ProtoReflect.Descriptor instead.
func (*Icon) Descriptor() ([]byte, []int) {
	return file_user_v1_user_service_proto_rawDescGZIP(), []int{14}
}

func (x *Icon) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Icon) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_user_v1_user_service_proto protoreflect.FileDescriptor

var file_user_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x4c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x41, 0x64, 0x64,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5f, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x75, 0x69,
	0x64, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0xbe, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x30, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x73, 0x22,
	0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x22, 0x57, 0x0a, 0x16, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x17,
	0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x63, 0x6f, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2e, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x32, 0xad, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x41, 0x64, 0x64,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x45, 0x64, 0x69,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x65, 0x74, 0x7a, 0x79, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_service_proto_rawDescOnce sync.Once
	file_user_v1_user_service_proto_rawDescData = file_user_v1_user_service_proto_rawDesc
)

func file_user_v1_user_service_proto_rawDescGZIP() []byte {
	file_user_v1_user_service_proto_rawDescOnce.Do(func() {
		file_user_v1_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_service_proto_rawDescData)
	})
	return file_user_v1_user_service_proto_rawDescData
}

var file_user_v1_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_user_v1_user_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),        // 0: user.v1.CreateUserRequest
	(*CreateUserResponse)(nil),       // 1: user.v1.CreateUserResponse
	(*AddChatServerRequest)(nil),     // 2: user.v1.AddChatServerRequest
	(*AddChatServerResponse)(nil),    // 3: user.v1.AddChatServerResponse
	(*RemoveChatServerRequest)(nil),  // 4: user.v1.RemoveChatServerRequest
	(*RemoveChatServerResponse)(nil), // 5: user.v1.RemoveChatServerResponse
	(*GetUserRequest)(nil),           // 6: user.v1.GetUserRequest
	(*GetUserResponse)(nil),          // 7: user.v1.GetUserResponse
	(*GetUsersRequest)(nil),          // 8: user.v1.GetUsersRequest
	(*GetUsersResponse)(nil),         // 9: user.v1.GetUsersResponse
	(*EditDescriptionRequest)(nil),   // 10: user.v1.EditDescriptionRequest
	(*EditDescriptionResponse)(nil),  // 11: user.v1.EditDescriptionResponse
	(*ChangeAvatarRequest)(nil),      // 12: user.v1.ChangeAvatarRequest
	(*ChangeAvatarResponse)(nil),     // 13: user.v1.ChangeAvatarResponse
	(*Icon)(nil),                     // 14: user.v1.Icon
}
var file_user_v1_user_service_proto_depIdxs = []int32{
	14, // 0: user.v1.GetUserResponse.icon:type_name -> user.v1.Icon
	7,  // 1: user.v1.GetUsersResponse.users:type_name -> user.v1.GetUserResponse
	0,  // 2: user.v1.UserService.CreateUser:input_type -> user.v1.CreateUserRequest
	2,  // 3: user.v1.UserService.AddChatServer:input_type -> user.v1.AddChatServerRequest
	4,  // 4: user.v1.UserService.RemoveChatServer:input_type -> user.v1.RemoveChatServerRequest
	6,  // 5: user.v1.UserService.GetUser:input_type -> user.v1.GetUserRequest
	8,  // 6: user.v1.UserService.GetUsers:input_type -> user.v1.GetUsersRequest
	10, // 7: user.v1.UserService.EditDescription:input_type -> user.v1.EditDescriptionRequest
	12, // 8: user.v1.UserService.ChangeAvatar:input_type -> user.v1.ChangeAvatarRequest
	1,  // 9: user.v1.UserService.CreateUser:output_type -> user.v1.CreateUserResponse
	3,  // 10: user.v1.UserService.AddChatServer:output_type -> user.v1.AddChatServerResponse
	5,  // 11: user.v1.UserService.RemoveChatServer:output_type -> user.v1.RemoveChatServerResponse
	7,  // 12: user.v1.UserService.GetUser:output_type -> user.v1.GetUserResponse
	9,  // 13: user.v1.UserService.GetUsers:output_type -> user.v1.GetUsersResponse
	11, // 14: user.v1.UserService.EditDescription:output_type -> user.v1.EditDescriptionResponse
	13, // 15: user.v1.UserService.ChangeAvatar:output_type -> user.v1.ChangeAvatarResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_v1_user_service_proto_init() }
func file_user_v1_user_service_proto_init() {
	if File_user_v1_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_v1_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_user_v1_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChatServerRequest); i {
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
		file_user_v1_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChatServerResponse); i {
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
		file_user_v1_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveChatServerRequest); i {
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
		file_user_v1_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveChatServerResponse); i {
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
		file_user_v1_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_user_v1_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_user_v1_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_user_v1_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_user_v1_user_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditDescriptionRequest); i {
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
		file_user_v1_user_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditDescriptionResponse); i {
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
		file_user_v1_user_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_v1_user_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_v1_user_service_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Icon); i {
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
			RawDescriptor: file_user_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_service_proto_goTypes,
		DependencyIndexes: file_user_v1_user_service_proto_depIdxs,
		MessageInfos:      file_user_v1_user_service_proto_msgTypes,
	}.Build()
	File_user_v1_user_service_proto = out.File
	file_user_v1_user_service_proto_rawDesc = nil
	file_user_v1_user_service_proto_goTypes = nil
	file_user_v1_user_service_proto_depIdxs = nil
}
