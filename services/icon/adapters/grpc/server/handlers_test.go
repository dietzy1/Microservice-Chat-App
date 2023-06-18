package server

import (
	"bytes"
	"context"
	"testing"

	"github.com/dietzy1/chatapp/services/icon/domain"
	iconv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func TestUploadIcon(t *testing.T) {
	// Create a mock domain object that returns a new icon
	mockDomain := &mockDomain{
		icon: domain.Icon{
			Uuid: "1",
			Link: "http://example.com/icon1.png",
		},
	}

	// Create a new server object with the mock domain
	server := &server{domain: mockDomain}

	// Create a new UploadIcon request
	req := &iconv1.UploadIconRequest{
		Data: &iconv1.UploadIconRequest_Info{
			Info: &iconv1.ImageInfo{
				Kindof:    ".png",
				OwnerUuid: "1",
			},
		},
	}

	// Create a new stream
	stream := &mockIconService_UploadIconServer{
		recv: func() (*iconv1.UploadIconRequest, error) {
			return req, nil
		},
		sendAndClose: func(resp *iconv1.UploadIconResponse) error {
			if resp.Uuid != "1" {
				t.Fatalf("UploadAvatar returned incorrect UUID: %s", resp.Uuid)
			}
			if resp.Link != "http://example.com/icon1.png" {
				t.Fatalf("UploadAvatar returned incorrect link: %s", resp.Link)
			}
			return nil
		},
	}

	// Call the UploadAvatar function
	err := server.UploadIcon(stream)
	if err != nil {
		t.Fatalf("UploadAvatar returned an error: %v", err)
	}
}

// Define a mock stream object that implements the iconv1.IconService_UploadIconServer interface
type mockIconService_UploadIconServer struct {
	recv         func() (*iconv1.UploadIconRequest, error)
	sendAndClose func(*iconv1.UploadIconResponse) error
	grpc.ServerStream
}

func (m *mockIconService_UploadIconServer) Recv() (*iconv1.UploadIconRequest, error) {
	req, err := m.recv()
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (m *mockIconService_UploadIconServer) SendAndClose(resp *iconv1.UploadIconResponse) error {
	return m.sendAndClose(resp)
}

func TestGetIcons(t *testing.T) {
	// Create a mock domain object that returns a list of icons
	mockDomain := &mockDomain{
		icons: []domain.Icon{
			{Uuid: "1", Link: "http://example.com/icon1.png", Kindof: "user", OwnerUuid: "1"},
			{Uuid: "2", Link: "http://example.com/icon2.png", Kindof: "user", OwnerUuid: "2"},
			{Uuid: "3", Link: "http://example.com/icon3.png", Kindof: "group", OwnerUuid: "1"},
		},
	}

	// Create a new server object with the mock domain
	server := &server{domain: mockDomain}

	// Create a new GetIcons request
	req := &iconv1.GetIconsRequest{}

	// Call the GetIcons function
	resp, err := server.GetIcons(context.Background(), req)
	if err != nil {
		t.Fatalf("GetIcons returned an error: %v", err)
	}

	// Check that the response contains the correct number of icons
	if len(resp.Icons) != 3 {
		t.Fatalf("GetIcons returned %d icons, expected 3", len(resp.Icons))
	}

	// Check that the response contains the correct icons
	expectedIcons := []*iconv1.Icon{
		{Uuid: "1", Link: "http://example.com/icon1.png", Kindof: "user", OwnerUuid: "1"},
		{Uuid: "2", Link: "http://example.com/icon2.png", Kindof: "user", OwnerUuid: "2"},
		{Uuid: "3", Link: "http://example.com/icon3.png", Kindof: "group", OwnerUuid: "1"},
	}
	for i, expectedIcon := range expectedIcons {
		if !proto.Equal(resp.Icons[i], expectedIcon) {
			t.Fatalf("GetIcons returned incorrect icon at index %d", i)
		}
	}
}

// Define a mock domain object that implements the domain.Domain interface
type mockDomain struct {
	icon  domain.Icon
	icons []domain.Icon
}

func (m *mockDomain) GetIcons(ctx context.Context) ([]domain.Icon, error) {
	return m.icons, nil
}

func (m *mockDomain) UploadIcon(ctx context.Context, imageData bytes.Buffer) (domain.Icon, error) {
	return domain.Icon{}, nil
}
