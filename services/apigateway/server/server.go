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

	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/chatroomgateway/v1"
	messagev1 "github.com/dietzy1/chatapp/services/apigateway/messagegateway/v1"
	userv1 "github.com/dietzy1/chatapp/services/apigateway/usergateway/v1"

	client "github.com/dietzy1/chatapp/services/apigateway/clients"

	//import the generated protobuf code straight from their source
	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"

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

	cache Cache

	//authClient client.AuthServiceClient
	authClient authclientv1.AuthServiceClient

	/*
		 	messageClient  client.MessageClient
			userClient     client.UserClient
			chatroomClient client.ChatRoomClient
	*/
}

// Create a new server object and inject the cache and clients
func newServer(cache Cache, authClient authclientv1.AuthServiceClient /* messageClient client.MessageClient, userClient client.UserClient, chatRoomClient client.ChatRoomClient */) *server {
	return &server{cache: cache, authClient: authClient /* messageClient: messageClient, userClient: userClient, chatroomClient: chatRoomClient */}
}

// run the generated GRPC gateway server
func runGateway() error {
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

	gatewayAddress := os.Getenv("GATEWAY")
	//gatewayAddress := ":8090"

	//middleware chaining
	middleware := logger(cors(gwmux))

	gwServer := &http.Server{
		Addr: gatewayAddress,
		//Handler: cors(gwmux),
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
	//addr := ":8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	//initiate dependencies for the server
	lruCache := cache.New(1000)
	authClient := client.NewAuthClient()
	//chatroomclient := client.NewChatRoomClient()

	/* messageClient := client.NewMessageClient()
	userClient := client.NewUserClient()
	chatRoomClient := client.NewChatRoomClient() */

	dependencies := newServer(&lruCache, *authClient /* *messageClient, *userClient, *chatRoomClient */)

	//Inject dependencies into the server
	s := grpc.NewServer()
	authv1.RegisterAuthGatewayServiceServer(s, dependencies)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()
	//Run the GRPC gateway server
	err = runGateway()
	log.Fatalln(err)
}
