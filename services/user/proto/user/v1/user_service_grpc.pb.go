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
	ChangeAvatar(ctx context.Context, in *ChangeAvatarRequest, opts ...grpc.CallOption) (*ChangeAvatarResponse, error)
	ChangeName(ctx context.Context, in *ChangeNameRequest, opts ...grpc.CallOption) (*ChangeNameResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
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

func (c *userServiceClient) ChangeAvatar(ctx context.Context, in *ChangeAvatarRequest, opts ...grpc.CallOption) (*ChangeAvatarResponse, error) {
	out := new(ChangeAvatarResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/ChangeAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeName(ctx context.Context, in *ChangeNameRequest, opts ...grpc.CallOption) (*ChangeNameResponse, error) {
	out := new(ChangeNameResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/ChangeName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/ChangePassword", in, out, opts...)
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
	ChangeAvatar(context.Context, *ChangeAvatarRequest) (*ChangeAvatarResponse, error)
	ChangeName(context.Context, *ChangeNameRequest) (*ChangeNameResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
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
func (UnimplementedUserServiceServer) ChangeAvatar(context.Context, *ChangeAvatarRequest) (*ChangeAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAvatar not implemented")
}
func (UnimplementedUserServiceServer) ChangeName(context.Context, *ChangeNameRequest) (*ChangeNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeName not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
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

func _UserService_ChangeName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/ChangeName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeName(ctx, req.(*ChangeNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
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
			MethodName: "ChangeAvatar",
			Handler:    _UserService_ChangeAvatar_Handler,
		},
		{
			MethodName: "ChangeName",
			Handler:    _UserService_ChangeName_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user_service.proto",
}
