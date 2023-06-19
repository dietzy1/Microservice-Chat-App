package server

import (
	"context"
	"reflect"
	"testing"

	accountclientv1 "github.com/dietzy1/chatapp/pkg/clients"
	accountv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockAccountClient struct {
	ChangeUsernameFunc   func(ctx context.Context, in *accountv1.ChangeUsernameRequest, opts ...grpc.CallOption) (*accountv1.ChangeUsernameResponse, error)
	ChangePasswordFunc   func(ctx context.Context, in *accountv1.ChangePasswordRequest, opts ...grpc.CallOption) (*accountv1.ChangePasswordResponse, error)
	DeleteAccountFunc    func(ctx context.Context, in *accountv1.DeleteAccountRequest, opts ...grpc.CallOption) (*accountv1.DeleteAccountResponse, error)
	RegisterAccountFunc  func(ctx context.Context, in *accountv1.RegisterAccountRequest, opts ...grpc.CallOption) (*accountv1.RegisterAccountResponse, error)
	DemoUserRegisterFunc func(ctx context.Context, in *accountv1.DemoUserRegisterRequest, opts ...grpc.CallOption) (*accountv1.DemoUserRegisterResponse, error)
	UpgradeDemoUserFunc  func(ctx context.Context, in *accountv1.UpgradeDemoUserRequest, opts ...grpc.CallOption) (*accountv1.UpgradeDemoUserResponse, error)
}

func (m *mockAccountClient) ChangeUsername(ctx context.Context, in *accountv1.ChangeUsernameRequest, opts ...grpc.CallOption) (*accountv1.ChangeUsernameResponse, error) {
	return m.ChangeUsernameFunc(ctx, in, opts...)

}

func (m *mockAccountClient) ChangePassword(ctx context.Context, in *accountv1.ChangePasswordRequest, opts ...grpc.CallOption) (*accountv1.ChangePasswordResponse, error) {
	return m.ChangePasswordFunc(ctx, in, opts...)
}

func (m *mockAccountClient) DeleteAccount(ctx context.Context, in *accountv1.DeleteAccountRequest, opts ...grpc.CallOption) (*accountv1.DeleteAccountResponse, error) {
	return m.DeleteAccountFunc(ctx, in, opts...)
}

func (m *mockAccountClient) RegisterAccount(ctx context.Context, in *accountv1.RegisterAccountRequest, opts ...grpc.CallOption) (*accountv1.RegisterAccountResponse, error) {
	return m.RegisterAccountFunc(ctx, in, opts...)
}

func (m *mockAccountClient) DemoUserRegister(ctx context.Context, in *accountv1.DemoUserRegisterRequest, opts ...grpc.CallOption) (*accountv1.DemoUserRegisterResponse, error) {
	return m.DemoUserRegisterFunc(ctx, in, opts...)
}

func (m *mockAccountClient) UpgradeDemoUser(ctx context.Context, in *accountv1.UpgradeDemoUserRequest, opts ...grpc.CallOption) (*accountv1.UpgradeDemoUserResponse, error) {
	return m.UpgradeDemoUserFunc(ctx, in, opts...)
}

func (m *mockAccountClient) ChangeUsernameCalls() int {
	return 1
}

func (m *mockAccountClient) ChangePasswordCalls() int {
	return 1
}

func (m *mockAccountClient) DeleteAccountCalls() int {
	return 1
}

func (m *mockAccountClient) RegisterAccountCalls() int {
	return 1
}

func (m *mockAccountClient) DemoUserRegisterCalls() int {
	return 1
}

func TestChangeUsername(t *testing.T) {
	// Create a new server instance
	s := &server{}

	// Create a new context
	ctx := context.Background()

	// Create a new ChangeUsernameRequest
	req := &accountv1.ChangeUsernameRequest{
		UserUuid: "123",
		Username: "new_username",
	}

	// Create a new ChangeUsernameResponse
	want := &accountv1.ChangeUsernameResponse{}

	// Create a new mock account client
	mockAccountClient := &mockAccountClient{}

	// Set the expected request and response for the ChangeUsername method
	mockAccountClient.ChangeUsernameFunc = func(ctx context.Context, in *accountv1.ChangeUsernameRequest, opts ...grpc.CallOption) (*accountv1.ChangeUsernameResponse, error) {
		if in.UserUuid != req.UserUuid || in.Username != req.Username {
			return nil, status.Error(codes.InvalidArgument, "Invalid request")
		}
		return &accountv1.ChangeUsernameResponse{}, nil
	}

	// Set the account client for the server
	s.accountClient = mockAccountClient

	// Call the ChangeUsername method
	got, err := s.ChangeUsername(ctx, req)

	// Check for errors
	if err != nil {
		t.Fatalf("ChangeUsername returned an error: %v", err)
	}

	// Check the response
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChangeUsername returned %v, want %v", got, want)
	}

	// Check the mock account client
	if mockAccountClient.ChangeUsernameCalls() != 1 {
		t.Errorf("ChangeUsername was not called")
	}
}

func TestChangePassword(t *testing.T) {
	// Create a new server instance
	s := &server{}

	// Create a new context
	ctx := context.Background()

	// Create a new ChangePasswordRequest
	req := &accountv1.ChangePasswordRequest{
		UserUuid: "123",
		Password: "new_password",
	}

	// Create a new ChangePasswordResponse
	want := &accountv1.ChangePasswordResponse{}

	// Create a new mock account client
	mockAccountClient := &mockAccountClient{}

	// Set the expected request and response for the ChangePassword method
	mockAccountClient.ChangePasswordFunc = func(ctx context.Context, in *accountclientv1.ChangePasswordRequest, opts ...grpc.CallOption) (*accountclientv1.ChangePasswordResponse, error) {
		if in.UserUuid != req.UserUuid || in.Password != req.Password {
			return nil, status.Error(codes.InvalidArgument, "Invalid request")
		}
		return &accountclientv1.ChangePasswordResponse{}, nil
	}

	// Set the account client for the server
	s.accountClient = mockAccountClient

	// Call the ChangePassword method
	got, err := s.ChangePassword(ctx, req)

	// Check for errors
	if err != nil {
		t.Fatalf("ChangePassword returned an error: %v", err)
	}

	// Check the response
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChangePassword returned %v, want %v", got, want)
	}

	// Check the mock account client
	if mockAccountClient.ChangePasswordCalls() != 1 {
		t.Errorf("ChangePassword was not called")
	}
}

func TestRegisterAccount(t *testing.T) {
	// Create a new server instance
	s := &server{}

	// Create a new context
	ctx := context.Background()

	// Create a new RegisterAccountRequest
	req := &accountv1.RegisterAccountRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	// Create a new RegisterAccountResponse
	want := &accountv1.RegisterAccountResponse{}

	// Create a new mock account client
	mockAccountClient := &mockAccountClient{}

	// Set the expected request and response for the RegisterAccount method
	mockAccountClient.RegisterAccountFunc = func(ctx context.Context, in *accountclientv1.RegisterAccountRequest, opts ...grpc.CallOption) (*accountclientv1.RegisterAccountResponse, error) {
		if in.Username != req.Username || in.Password != req.Password {
			return nil, status.Error(codes.InvalidArgument, "Invalid request")
		}
		return &accountclientv1.RegisterAccountResponse{Session: "testsession"}, nil
	}

	// Set the account client for the server
	s.accountClient = mockAccountClient

	// Call the RegisterAccount method
	got, err := s.RegisterAccount(ctx, req)

	// Check for errors
	if err != nil {
		t.Fatalf("RegisterAccount returned an error: %v", err)
	}

	// Check the response
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RegisterAccount returned %v, want %v", got, want)
	}

	// Check the mock account client
	if mockAccountClient.RegisterAccountCalls() != 1 {
		t.Errorf("RegisterAccount was not called")
	}
}

func TestDemoUserRegister(t *testing.T) {
	// Create a new server instance
	s := &server{}

	// Create a new context
	ctx := context.Background()

	// Create a new DemoUserRegisterRequest
	req := &accountv1.DemoUserRegisterRequest{}

	// Create a new DemoUserRegisterResponse
	want := &accountv1.DemoUserRegisterResponse{}

	// Create a new mock account client
	mockAccountClient := &mockAccountClient{}

	// Set the expected request and response for the DemoUserRegister method
	mockAccountClient.DemoUserRegisterFunc = func(ctx context.Context, in *accountv1.DemoUserRegisterRequest, opts ...grpc.CallOption) (*accountv1.DemoUserRegisterResponse, error) {
		return &accountv1.DemoUserRegisterResponse{Session: "testsession"}, nil
	}

	// Set the account client for the server
	s.accountClient = mockAccountClient

	// Call the DemoUserRegister method
	got, err := s.DemoUserRegister(ctx, req)

	// Check for errors
	if err != nil {
		t.Fatalf("DemoUserRegister returned an error: %v", err)
	}

	// Check the response
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DemoUserRegister returned %v, want %v", got, want)
	}

	// Check the mock account client
	if mockAccountClient.DemoUserRegisterCalls() != 1 {
		t.Errorf("DemoUserRegister was not called")
	}
}

func TestUpgradeDemoUser(t *testing.T) {
	// Create a new server instance
	s := &server{}

	// Create a new context
	ctx := context.Background()

	// Create a new UpgradeDemoUserRequest
	req := &accountv1.UpgradeDemoUserRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	// Create a new UpgradeDemoUserResponse
	want := &accountv1.UpgradeDemoUserResponse{
		Session: "testsession",
	}

	// Create a new mock account client
	mockAccountClient := &mockAccountClient{}

	// Set the expected request and response for the UpgradeDemoUser method
	mockAccountClient.UpgradeDemoUserFunc = func(ctx context.Context, in *accountv1.UpgradeDemoUserRequest, opts ...grpc.CallOption) (*accountv1.UpgradeDemoUserResponse, error) {
		if in.Username != req.Username || in.Password != req.Password {
			return nil, status.Error(codes.InvalidArgument, "Invalid request")
		}
		return &accountv1.UpgradeDemoUserResponse{Session: "testsession"}, nil
	}

	// Set the account client for the server
	s.accountClient = mockAccountClient

	// Call the UpgradeDemoUser method
	got, err := s.UpgradeDemoUser(ctx, req)

	// Check for errors
	if err != nil {
		t.Fatalf("UpgradeDemoUser returned an error: %v", err)
	}

	// Check the response
	if !reflect.DeepEqual(got, want) {
		t.Errorf("UpgradeDemoUser returned %v, want %v", got, want)
	}

	// Check the mock account client
	if mockAccountClient.UpgradeDemoUserCalls() != 1 {
		t.Errorf("UpgradeDemoUser was not called")
	}
}
