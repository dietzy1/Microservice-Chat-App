package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/dietzy1/chatapp/pkg/clients"
	accountv1 "github.com/dietzy1/chatapp/services/apigateway/accountgateway/v1"
	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/chatroomgateway/v1"
	iconv1 "github.com/dietzy1/chatapp/services/apigateway/icongateway/v1"
	messagev1 "github.com/dietzy1/chatapp/services/apigateway/messagegateway/v1"
	userv1 "github.com/dietzy1/chatapp/services/apigateway/usergateway/v1"

	//import the generated protobuf code straight from their source
	accountclientv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	chatroomclientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
	iconclientv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
	messageclientv1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	userclientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

type server struct {
	authv1.UnimplementedAuthGatewayServiceServer
	messagev1.UnimplementedMessageGatewayServiceServer
	userv1.UnimplementedUserGatewayServiceServer
	chatroomv1.UnimplementedChatroomGatewayServiceServer
	accountclientv1.UnimplementedAccountServiceServer
	iconclientv1.UnimplementedIconServiceServer

	authClient authclientv1.AuthServiceClient

	userClient userclientv1.UserServiceClient

	messageClient messageclientv1.MessageServiceClient

	chatroomClient chatroomclientv1.ChatroomServiceClient

	accountClient accountclientv1.AccountServiceClient

	iconClient iconclientv1.IconServiceClient

	logger *zap.Logger
}

// Create a new server object and inject the cache and clients
func newServer(authClient authclientv1.AuthServiceClient, userClient userclientv1.UserServiceClient, messageClient messageclientv1.MessageServiceClient, chatroomClient chatroomclientv1.ChatroomServiceClient, accountClient accountclientv1.AccountServiceClient, iconClient iconclientv1.IconServiceClient, logger *zap.Logger) *server {
	return &server{authClient: authClient, userClient: userClient, messageClient: messageClient, chatroomClient: chatroomClient, accountClient: accountClient, iconClient: iconClient, logger: logger}
}

// run the generated GRPC gateway server
func runGateway(authClient authclientv1.AuthServiceClient) error {
	log := grpclog.NewLoggerV2WithVerbosity(os.Stdout, io.Discard, io.Discard, 1)
	grpclog.SetLoggerV2(log)

	//The reverse proxy connects to the GRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		/* "dns:///0.0.0.0:8080", */
		"dns:///0.0.0.0"+os.Getenv("GRPC"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}

	//intercepts the response and reads the cookie
	incomingHeader := incomingHeaderMatcherWrapper()

	//intercepts the response and sets the cookie
	forwardResponse := withForwardResponseOptionWrapper()

	//intercepts the request and sets the cookie
	withMetaData := withMetaDataWrapper()

	//Main mux where options are added in
	gwmux := runtime.NewServeMux(incomingHeader, forwardResponse, withMetaData)

	//Registration of the gateway
	if err = authv1.RegisterAuthGatewayServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}
	if err = userv1.RegisterUserGatewayServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}
	if err = messagev1.RegisterMessageGatewayServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}
	if err = chatroomv1.RegisterChatroomGatewayServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}
	if err = accountv1.RegisterAccountGatewayServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}
	if err = iconv1.RegisterIconGatewayServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}

	gatewayAddress := os.Getenv("GATEWAY")
	//gatewayAddress := ":8090"

	//middleware chaining
	mw := wrapperAuthMiddleware(authClient)
	middleware := mw(logger(cors(gwmux)))

	//middleware := logger(cors(gwmux))

	gwServer := &http.Server{
		Addr:    gatewayAddress,
		Handler: middleware,
	}

	log.Info("Serving gRPC-Gateway", gatewayAddress)
	log.Fatalln(gwServer.ListenAndServe())

	return nil
}

// This is the function called in main.go
func Start(logger *zap.Logger) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("GRPC")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}
	//initiate dependencies for the server

	authClient := clients.NewAuthClient()
	userClient := clients.NewUserClient()
	chatroomClient := clients.NewChatRoomClient()
	messageClient := clients.NewMessageClient()
	accountClient := clients.NewAccountClient()
	iconClient := clients.NewIconClient()

	dependencies := newServer(*authClient, *userClient, *messageClient, *chatroomClient, *accountClient, *iconClient, logger)

	//Inject dependencies into the server
	s := grpc.NewServer()

	authv1.RegisterAuthGatewayServiceServer(s, dependencies)
	messagev1.RegisterMessageGatewayServiceServer(s, dependencies)
	userv1.RegisterUserGatewayServiceServer(s, dependencies)
	chatroomv1.RegisterChatroomGatewayServiceServer(s, dependencies)
	accountv1.RegisterAccountGatewayServiceServer(s, dependencies)

	// Serve gRPC Server
	logger.Info("Serving gRPC on http://", zap.String("address", addr))
	go func() {
		log.Fatal(s.Serve(lis))

	}()

	//create new stdlib mux
	r := http.DefaultServeMux

	//Attach REST endpoints to the mux
	r.HandleFunc("/api/v1/icons", dependencies.uploadIconHandler)
	//r.Handle("/swagger-ui/", serveSwaggerUI())

	srv := &http.Server{
		Handler:      r,
		Addr:         os.Getenv("ADDR") + os.Getenv("PORT"), //Required addr for railway.app deployment.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	//Serve the REST Server
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	//Run the GRPC gateway server
	err = runGateway(*authClient)
	log.Fatalln(err)
}
