// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: chatroom/v1/chatroom_service.proto

package chatroomv1

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error)
	LeaveRoom(ctx context.Context, in *LeaveRoomRequest, opts ...grpc.CallOption) (*LeaveRoomResponse, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/chatroom.v1.AuthService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error) {
	out := new(DeleteRoomResponse)
	err := c.cc.Invoke(ctx, "/chatroom.v1.AuthService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error) {
	out := new(JoinRoomResponse)
	err := c.cc.Invoke(ctx, "/chatroom.v1.AuthService/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LeaveRoom(ctx context.Context, in *LeaveRoomRequest, opts ...grpc.CallOption) (*LeaveRoomResponse, error) {
	out := new(LeaveRoomResponse)
	err := c.cc.Invoke(ctx, "/chatroom.v1.AuthService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, "/chatroom.v1.AuthService/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error)
	JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error)
	LeaveRoom(context.Context, *LeaveRoomRequest) (*LeaveRoomResponse, error)
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
}

// UnimplementedAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedAuthServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedAuthServiceServer) JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedAuthServiceServer) LeaveRoom(context.Context, *LeaveRoomRequest) (*LeaveRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (UnimplementedAuthServiceServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroom.v1.AuthService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroom.v1.AuthService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroom.v1.AuthService/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroom.v1.AuthService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LeaveRoom(ctx, req.(*LeaveRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroom.v1.AuthService/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatroom.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _AuthService_CreateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _AuthService_DeleteRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _AuthService_JoinRoom_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _AuthService_LeaveRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _AuthService_GetRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatroom/v1/chatroom_service.proto",
}
