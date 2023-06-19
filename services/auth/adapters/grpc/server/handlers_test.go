package server

import (
	"context"
	"testing"

	"github.com/dietzy1/chatapp/services/auth/domain"
	authv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
)

type mockAuth struct {
	LoginFunc        func(ctx context.Context, cred domain.Credentials) (domain.Credentials, error)
	LogoutFunc       func(ctx context.Context, session string, useruuid string) error
	AuthenticateFunc func(ctx context.Context, session string, useruuid string) (domain.Credentials, error)
	InvalidateFunc   func(ctx context.Context, userUUid string) error
}

func (m *mockAuth) Login(ctx context.Context, cred domain.Credentials) (domain.Credentials, error) {
	return m.LoginFunc(ctx, cred)
}

func (m *mockAuth) Logout(ctx context.Context, session string, useruuid string) error {
	return m.LogoutFunc(ctx, session, useruuid)
}

func (m *mockAuth) Authenticate(ctx context.Context, session string, useruuid string) (domain.Credentials, error) {
	return m.AuthenticateFunc(ctx, session, useruuid)
}

func (m *mockAuth) Invalidate(ctx context.Context, userUUid string) error {
	return m.InvalidateFunc(ctx, userUUid)
}

func TestLogin(t *testing.T) {
	mock := &mockAuth{
		LoginFunc: func(ctx context.Context, cred domain.Credentials) (domain.Credentials, error) {
			return domain.Credentials{
				Session: "test-session",
				Uuid:    "test-uuid",
			}, nil
		},
	}

	s := &server{auth: mock}

	req := &authv1.LoginRequest{
		Username: "test-username",
		Password: "test-password",
	}

	resp, err := s.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Session != "test-session" {
		t.Errorf("expected session to be %q, but got %q", "test-session", resp.Session)
	}

	if resp.UserUuid != "test-uuid" {
		t.Errorf("expected user UUID to be %q, but got %q", "test-uuid", resp.UserUuid)
	}
}

func TestLogout(t *testing.T) {
	mock := &mockAuth{
		LogoutFunc: func(ctx context.Context, session string, useruuid string) error {
			return nil
		},
	}

	s := &server{auth: mock}

	req := &authv1.LogoutRequest{
		Session:  "test-session",
		UserUuid: "test-uuid",
	}

	_, err := s.Logout(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAuthenticate(t *testing.T) {
	mock := &mockAuth{
		AuthenticateFunc: func(ctx context.Context, session string, useruuid string) (domain.Credentials, error) {
			return domain.Credentials{
				Session: "test-session",
				Uuid:    "test-uuid",
			}, nil
		},
	}

	s := &server{auth: mock}

	req := &authv1.AuthenticateRequest{
		Session:  "test-session",
		UserUuid: "test-uuid",
	}

	resp, err := s.Authenticate(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Session != "test-session" {
		t.Errorf("expected session to be %q, but got %q", "test-session", resp.Session)
	}

	if resp.UserUuid != "test-uuid" {
		t.Errorf("expected user UUID to be %q, but got %q", "test-uuid", resp.UserUuid)
	}
}

func TestInvalidate(t *testing.T) {
	mock := &mockAuth{
		InvalidateFunc: func(ctx context.Context, userUUid string) error {
			return nil
		},
	}

	s := &server{auth: mock}

	req := &authv1.InvalidateRequest{
		UserUuid: "test-uuid",
	}

	_, err := s.Invalidate(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
