// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: accountgateway/v1/accountgateway_service.proto

package accountgatewayv1

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

// AccountGatewayServiceClient is the client API for AccountGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountGatewayServiceClient interface {
	// External RPC
	ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type accountGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountGatewayServiceClient(cc grpc.ClientConnInterface) AccountGatewayServiceClient {
	return &accountGatewayServiceClient{cc}
}

func (c *accountGatewayServiceClient) ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error) {
	out := new(ChangeUsernameResponse)
	err := c.cc.Invoke(ctx, "/accountgateway.v1.AccountGatewayService/ChangeUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGatewayServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/accountgateway.v1.AccountGatewayService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGatewayServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/accountgateway.v1.AccountGatewayService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountGatewayServiceServer is the server API for AccountGatewayService service.
// All implementations should embed UnimplementedAccountGatewayServiceServer
// for forward compatibility
type AccountGatewayServiceServer interface {
	// External RPC
	ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
}

// UnimplementedAccountGatewayServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountGatewayServiceServer struct {
}

func (UnimplementedAccountGatewayServiceServer) ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUsername not implemented")
}
func (UnimplementedAccountGatewayServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAccountGatewayServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

// UnsafeAccountGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountGatewayServiceServer will
// result in compilation errors.
type UnsafeAccountGatewayServiceServer interface {
	mustEmbedUnimplementedAccountGatewayServiceServer()
}

func RegisterAccountGatewayServiceServer(s grpc.ServiceRegistrar, srv AccountGatewayServiceServer) {
	s.RegisterService(&AccountGatewayService_ServiceDesc, srv)
}

func _AccountGatewayService_ChangeUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGatewayServiceServer).ChangeUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountgateway.v1.AccountGatewayService/ChangeUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGatewayServiceServer).ChangeUsername(ctx, req.(*ChangeUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGatewayService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGatewayServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountgateway.v1.AccountGatewayService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGatewayServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGatewayService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGatewayServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountgateway.v1.AccountGatewayService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGatewayServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountGatewayService_ServiceDesc is the grpc.ServiceDesc for AccountGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accountgateway.v1.AccountGatewayService",
	HandlerType: (*AccountGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeUsername",
			Handler:    _AccountGatewayService_ChangeUsername_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AccountGatewayService_ChangePassword_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountGatewayService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accountgateway/v1/accountgateway_service.proto",
}
