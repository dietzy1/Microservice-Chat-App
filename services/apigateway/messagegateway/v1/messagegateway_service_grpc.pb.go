// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: messagegateway/v1/messagegateway_service.proto

package messagegatewayv1

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
	MessageGatewayService_GetMessages_FullMethodName   = "/messagegateway.v1.MessageGatewayService/GetMessages"
	MessageGatewayService_EditMessage_FullMethodName   = "/messagegateway.v1.MessageGatewayService/EditMessage"
	MessageGatewayService_DeleteMessage_FullMethodName = "/messagegateway.v1.MessageGatewayService/DeleteMessage"
)

// MessageGatewayServiceClient is the client API for MessageGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageGatewayServiceClient interface {
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	EditMessage(ctx context.Context, in *EditMessageRequest, opts ...grpc.CallOption) (*EditMessageResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
}

type messageGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageGatewayServiceClient(cc grpc.ClientConnInterface) MessageGatewayServiceClient {
	return &messageGatewayServiceClient{cc}
}

func (c *messageGatewayServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, MessageGatewayService_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageGatewayServiceClient) EditMessage(ctx context.Context, in *EditMessageRequest, opts ...grpc.CallOption) (*EditMessageResponse, error) {
	out := new(EditMessageResponse)
	err := c.cc.Invoke(ctx, MessageGatewayService_EditMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageGatewayServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, MessageGatewayService_DeleteMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageGatewayServiceServer is the server API for MessageGatewayService service.
// All implementations should embed UnimplementedMessageGatewayServiceServer
// for forward compatibility
type MessageGatewayServiceServer interface {
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	EditMessage(context.Context, *EditMessageRequest) (*EditMessageResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
}

// UnimplementedMessageGatewayServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMessageGatewayServiceServer struct {
}

func (UnimplementedMessageGatewayServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedMessageGatewayServiceServer) EditMessage(context.Context, *EditMessageRequest) (*EditMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessage not implemented")
}
func (UnimplementedMessageGatewayServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}

// UnsafeMessageGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageGatewayServiceServer will
// result in compilation errors.
type UnsafeMessageGatewayServiceServer interface {
	mustEmbedUnimplementedMessageGatewayServiceServer()
}

func RegisterMessageGatewayServiceServer(s grpc.ServiceRegistrar, srv MessageGatewayServiceServer) {
	s.RegisterService(&MessageGatewayService_ServiceDesc, srv)
}

func _MessageGatewayService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGatewayServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageGatewayService_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGatewayServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageGatewayService_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGatewayServiceServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageGatewayService_EditMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGatewayServiceServer).EditMessage(ctx, req.(*EditMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageGatewayService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGatewayServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageGatewayService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGatewayServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageGatewayService_ServiceDesc is the grpc.ServiceDesc for MessageGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messagegateway.v1.MessageGatewayService",
	HandlerType: (*MessageGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessages",
			Handler:    _MessageGatewayService_GetMessages_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _MessageGatewayService_EditMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _MessageGatewayService_DeleteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messagegateway/v1/messagegateway_service.proto",
}
