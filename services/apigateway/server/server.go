package server

//	authv1 "github.com/dietzy1/chatapp/services/apigateway/gen/go/auth/v1"
import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	//authv1 "github.com/dietzy1/chatapp/services/apigateway/gen/go/auth/v1"

	accountv1 "github.com/dietzy1/chatapp/services/apigateway/accountgateway/v1"
	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/chatroomgateway/v1"
	messagev1 "github.com/dietzy1/chatapp/services/apigateway/messagegateway/v1"
	userv1 "github.com/dietzy1/chatapp/services/apigateway/usergateway/v1"

	client "github.com/dietzy1/chatapp/services/apigateway/clients"

	//import the generated protobuf code straight from their source
	accountclientv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	chatroomclientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
	messageclientv1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	userclientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"

	"github.com/dietzy1/chatapp/services/apigateway/cache"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	cache Cache

	//authClient client.AuthServiceClient
	authClient authclientv1.AuthServiceClient

	//messageClient  client.MessageClient
	userClient userclientv1.UserServiceClient
	//chatroomClient client.ChatRoomClient

	messageClient messageclientv1.MessageServiceClient

	chatroomClient chatroomclientv1.ChatroomServiceClient

	accountClient accountclientv1.AccountServiceClient
}

// Create a new server object and inject the cache and clients
func newServer(cache Cache, authClient authclientv1.AuthServiceClient, userClient userclientv1.UserServiceClient, messageClient messageclientv1.MessageServiceClient, chatroomClient chatroomclientv1.ChatroomServiceClient, accountClient accountclientv1.AccountServiceClient) *server {
	return &server{cache: cache, authClient: authClient, userClient: userClient, messageClient: messageClient, chatroomClient: chatroomClient, accountClient: accountClient}
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
func Start() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("GRPC")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	//initiate dependencies for the server

	lruCache := cache.New(1000)
	authClient := client.NewAuthClient()
	userClient := client.NewUserClient()
	chatroomClient := client.NewChatRoomClient()
	messageClient := client.NewMessageClient()
	accountClient := client.NewAccountClient()

	dependencies := newServer(&lruCache, *authClient, *userClient, *messageClient, *chatroomClient, *accountClient)

	//Inject dependencies into the server
	s := grpc.NewServer()

	authv1.RegisterAuthGatewayServiceServer(s, dependencies)
	messagev1.RegisterMessageGatewayServiceServer(s, dependencies)
	userv1.RegisterUserGatewayServiceServer(s, dependencies)
	chatroomv1.RegisterChatroomGatewayServiceServer(s, dependencies)
	accountv1.RegisterAccountGatewayServiceServer(s, dependencies)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()
	//Run the GRPC gateway server
	err = runGateway(*authClient)
	log.Fatalln(err)
}
