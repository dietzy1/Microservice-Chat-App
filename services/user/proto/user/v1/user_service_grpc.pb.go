// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user/v1/user_service.proto

package userv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// internal RPC
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	AddChatServer(ctx context.Context, in *AddChatServerRequest, opts ...grpc.CallOption) (*AddChatServerResponse, error)
	RemoveChatServer(ctx context.Context, in *RemoveChatServerRequest, opts ...grpc.CallOption) (*RemoveChatServerResponse, error)
	// External RPC
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	EditDescription(ctx context.Context, in *EditDescriptionRequest, opts ...grpc.CallOption) (*EditDescriptionResponse, error)
	ChangeAvatar(ctx context.Context, in *ChangeAvatarRequest, opts ...grpc.CallOption) (*ChangeAvatarResponse, error)
	UploadAvatar(ctx context.Context, opts ...grpc.CallOption) (UserService_UploadAvatarClient, error)
	GetAvatars(ctx context.Context, in *GetAvatarsRequest, opts ...grpc.CallOption) (*GetAvatarsResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddChatServer(ctx context.Context, in *AddChatServerRequest, opts ...grpc.CallOption) (*AddChatServerResponse, error) {
	out := new(AddChatServerResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/AddChatServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveChatServer(ctx context.Context, in *RemoveChatServerRequest, opts ...grpc.CallOption) (*RemoveChatServerResponse, error) {
	out := new(RemoveChatServerResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/RemoveChatServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditDescription(ctx context.Context, in *EditDescriptionRequest, opts ...grpc.CallOption) (*EditDescriptionResponse, error) {
	out := new(EditDescriptionResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/EditDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeAvatar(ctx context.Context, in *ChangeAvatarRequest, opts ...grpc.CallOption) (*ChangeAvatarResponse, error) {
	out := new(ChangeAvatarResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/ChangeAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UploadAvatar(ctx context.Context, opts ...grpc.CallOption) (UserService_UploadAvatarClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/user.v1.UserService/UploadAvatar", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceUploadAvatarClient{stream}
	return x, nil
}

type UserService_UploadAvatarClient interface {
	Send(*UploadAvatarRequest) error
	CloseAndRecv() (*UploadAvatarResponse, error)
	grpc.ClientStream
}

type userServiceUploadAvatarClient struct {
	grpc.ClientStream
}

func (x *userServiceUploadAvatarClient) Send(m *UploadAvatarRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceUploadAvatarClient) CloseAndRecv() (*UploadAvatarResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadAvatarResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetAvatars(ctx context.Context, in *GetAvatarsRequest, opts ...grpc.CallOption) (*GetAvatarsResponse, error) {
	out := new(GetAvatarsResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/GetAvatars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// internal RPC
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	AddChatServer(context.Context, *AddChatServerRequest) (*AddChatServerResponse, error)
	RemoveChatServer(context.Context, *RemoveChatServerRequest) (*RemoveChatServerResponse, error)
	// External RPC
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	EditDescription(context.Context, *EditDescriptionRequest) (*EditDescriptionResponse, error)
	ChangeAvatar(context.Context, *ChangeAvatarRequest) (*ChangeAvatarResponse, error)
	UploadAvatar(UserService_UploadAvatarServer) error
	GetAvatars(context.Context, *GetAvatarsRequest) (*GetAvatarsResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) AddChatServer(context.Context, *AddChatServerRequest) (*AddChatServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChatServer not implemented")
}
func (UnimplementedUserServiceServer) RemoveChatServer(context.Context, *RemoveChatServerRequest) (*RemoveChatServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveChatServer not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) EditDescription(context.Context, *EditDescriptionRequest) (*EditDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDescription not implemented")
}
func (UnimplementedUserServiceServer) ChangeAvatar(context.Context, *ChangeAvatarRequest) (*ChangeAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAvatar not implemented")
}
func (UnimplementedUserServiceServer) UploadAvatar(UserService_UploadAvatarServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedUserServiceServer) GetAvatars(context.Context, *GetAvatarsRequest) (*GetAvatarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatars not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddChatServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChatServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddChatServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/AddChatServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddChatServer(ctx, req.(*AddChatServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveChatServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveChatServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveChatServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/RemoveChatServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveChatServer(ctx, req.(*RemoveChatServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/EditDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditDescription(ctx, req.(*EditDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/ChangeAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeAvatar(ctx, req.(*ChangeAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UploadAvatar_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).UploadAvatar(&userServiceUploadAvatarServer{stream})
}

type UserService_UploadAvatarServer interface {
	SendAndClose(*UploadAvatarResponse) error
	Recv() (*UploadAvatarRequest, error)
	grpc.ServerStream
}

type userServiceUploadAvatarServer struct {
	grpc.ServerStream
}

func (x *userServiceUploadAvatarServer) SendAndClose(m *UploadAvatarResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceUploadAvatarServer) Recv() (*UploadAvatarRequest, error) {
	m := new(UploadAvatarRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GetAvatars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvatarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAvatars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/GetAvatars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAvatars(ctx, req.(*GetAvatarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "AddChatServer",
			Handler:    _UserService_AddChatServer_Handler,
		},
		{
			MethodName: "RemoveChatServer",
			Handler:    _UserService_RemoveChatServer_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "EditDescription",
			Handler:    _UserService_EditDescription_Handler,
		},
		{
			MethodName: "ChangeAvatar",
			Handler:    _UserService_ChangeAvatar_Handler,
		},
		{
			MethodName: "GetAvatars",
			Handler:    _UserService_GetAvatars_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadAvatar",
			Handler:       _UserService_UploadAvatar_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "user/v1/user_service.proto",
}
