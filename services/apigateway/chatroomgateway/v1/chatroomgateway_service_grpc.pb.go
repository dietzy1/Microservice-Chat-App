// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chatroomgateway/v1/chatroomgateway_service.proto

package chatroomgatewayv1

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

const (
	ChatroomGatewayService_CreateRoom_FullMethodName    = "/chatroomgateway.v1.ChatroomGatewayService/CreateRoom"
	ChatroomGatewayService_DeleteRoom_FullMethodName    = "/chatroomgateway.v1.ChatroomGatewayService/DeleteRoom"
	ChatroomGatewayService_GetRoom_FullMethodName       = "/chatroomgateway.v1.ChatroomGatewayService/GetRoom"
	ChatroomGatewayService_GetRooms_FullMethodName      = "/chatroomgateway.v1.ChatroomGatewayService/GetRooms"
	ChatroomGatewayService_CreateChannel_FullMethodName = "/chatroomgateway.v1.ChatroomGatewayService/CreateChannel"
	ChatroomGatewayService_DeleteChannel_FullMethodName = "/chatroomgateway.v1.ChatroomGatewayService/DeleteChannel"
	ChatroomGatewayService_GetChannel_FullMethodName    = "/chatroomgateway.v1.ChatroomGatewayService/GetChannel"
	ChatroomGatewayService_InviteUser_FullMethodName    = "/chatroomgateway.v1.ChatroomGatewayService/InviteUser"
	ChatroomGatewayService_RemoveUser_FullMethodName    = "/chatroomgateway.v1.ChatroomGatewayService/RemoveUser"
	ChatroomGatewayService_AddUser_FullMethodName       = "/chatroomgateway.v1.ChatroomGatewayService/AddUser"
)

// ChatroomGatewayServiceClient is the client API for ChatroomGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatroomGatewayServiceClient interface {
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error)
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error)
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
}

type chatroomGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatroomGatewayServiceClient(cc grpc.ClientConnInterface) ChatroomGatewayServiceClient {
	return &chatroomGatewayServiceClient{cc}
}

func (c *chatroomGatewayServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error) {
	out := new(DeleteRoomResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_DeleteRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_GetRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error) {
	out := new(GetRoomsResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_GetRooms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_CreateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_DeleteChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error) {
	out := new(GetChannelResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_GetChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error) {
	out := new(InviteUserResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_InviteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomGatewayServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, ChatroomGatewayService_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatroomGatewayServiceServer is the server API for ChatroomGatewayService service.
// All implementations should embed UnimplementedChatroomGatewayServiceServer
// for forward compatibility
type ChatroomGatewayServiceServer interface {
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error)
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error)
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error)
	InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
}

// UnimplementedChatroomGatewayServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChatroomGatewayServiceServer struct {
}

func (UnimplementedChatroomGatewayServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUser not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedChatroomGatewayServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}

// UnsafeChatroomGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatroomGatewayServiceServer will
// result in compilation errors.
type UnsafeChatroomGatewayServiceServer interface {
	mustEmbedUnimplementedChatroomGatewayServiceServer()
}

func RegisterChatroomGatewayServiceServer(s grpc.ServiceRegistrar, srv ChatroomGatewayServiceServer) {
	s.RegisterService(&ChatroomGatewayService_ServiceDesc, srv)
}

func _ChatroomGatewayService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_DeleteRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_GetRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_GetRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_CreateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_DeleteChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_GetChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).GetChannel(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_InviteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatroomGatewayService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomGatewayServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatroomGatewayService_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomGatewayServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatroomGatewayService_ServiceDesc is the grpc.ServiceDesc for ChatroomGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatroomGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatroomgateway.v1.ChatroomGatewayService",
	HandlerType: (*ChatroomGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _ChatroomGatewayService_CreateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _ChatroomGatewayService_DeleteRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _ChatroomGatewayService_GetRoom_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _ChatroomGatewayService_GetRooms_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _ChatroomGatewayService_CreateChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _ChatroomGatewayService_DeleteChannel_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _ChatroomGatewayService_GetChannel_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _ChatroomGatewayService_InviteUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _ChatroomGatewayService_RemoveUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ChatroomGatewayService_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatroomgateway/v1/chatroomgateway_service.proto",
}
