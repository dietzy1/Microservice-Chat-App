package domain

import (
	"context"
	"testing"

	"github.com/dietzy1/chatapp/pkg/hashing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAuth struct {
	mock.Mock
}

func (m *mockAuth) Login(ctx context.Context, username string) (string, error) {
	args := m.Called(ctx, username)
	return args.String(0), args.Error(1)
}

func (m *mockAuth) Logout(ctx context.Context, username string) error {
	args := m.Called(ctx, username)
	return args.Error(0)
}

func (m *mockAuth) Authenticate(ctx context.Context, token string) (string, error) {
	args := m.Called(ctx, token)
	return args.String(0), args.Error(1)
}

func (m *mockAuth) UpdateToken(ctx context.Context, username string, token string) (string, error) {
	args := m.Called(ctx, username, token)
	return args.String(0), args.Error(1)
}

type mockCache struct {
	mock.Mock
}

func (m *mockCache) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *mockCache) Set(key string, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *mockCache) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestLogin(t *testing.T) {
	mockAuth := new(mockAuth)
	mockCache := new(mockCache)

	d := New(mockAuth, mockCache)

	ctx := context.Background()

	// Mock the hashing functions
	mockHash, err := hashing.GenerateHash("testpassword")
	assert.NoError(t, err)

	// Set up mock behavior for auth.Login
	mockAuth.On("Login", ctx, "testuser").Return(mockHash, nil)

	// Set up mock behavior for auth.UpdateToken
	mockAuth.On("UpdateToken", ctx, "testuser", mock.AnythingOfType("string")).Return("testuuid", nil)

	// Set up mock behavior for cache.Set
	mockCache.On("Set", "testuuid", mock.AnythingOfType("string")).Return(nil)

	// Call the Login method with test credentials
	creds := Credentials{Username: "testuser", Password: "testpassword"}
	result, err := d.Login(ctx, creds)

	// Check that the method returned the expected result
	assert.NoError(t, err)
	assert.Equal(t, "testuuid", result.Uuid)
	assert.NotEmpty(t, result.Session)

	// Check that the mock methods were called with the expected arguments
	mockAuth.AssertCalled(t, "Login", ctx, "testuser")
	mockAuth.AssertCalled(t, "UpdateToken", ctx, "testuser", mock.AnythingOfType("string"))
	mockCache.AssertCalled(t, "Set", "testuuid", mock.AnythingOfType("string"))
}

func TestLogout(t *testing.T) {
	mockAuth := new(mockAuth)
	mockCache := new(mockCache)

	d := New(mockAuth, mockCache)

	ctx := context.Background()

	// Set up mock behavior for auth.Authenticate
	mockAuth.On("Authenticate", ctx, "testuuid").Return("testsession", nil)

	// Set up mock behavior for auth.Logout
	mockAuth.On("Logout", ctx, "testuuid").Return(nil)

	// Set up mock behavior for cache.Get
	mockCache.On("Get", "testuuid").Return("testsession", nil)

	// Set up mock behavior for cache.Set
	mockCache.On("Set", "testuuid", "").Return(nil)

	// Call the Logout method with test credentials
	err := d.Logout(ctx, "testsession", "testuuid")

	// Check that the method returned the expected result
	assert.NoError(t, err)

	// Check that the mock methods were called with the expected arguments
	mockAuth.AssertCalled(t, "Authenticate", ctx, "testuuid")
	mockAuth.AssertCalled(t, "Logout", ctx, "testuuid")
	mockCache.AssertCalled(t, "Get", "testuuid")
	mockCache.AssertCalled(t, "Set", "testuuid", "")
}
