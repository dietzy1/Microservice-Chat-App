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
	"github.com/dietzy1/chatapp/services/apigateway/cache"
	client "github.com/dietzy1/chatapp/services/apigateway/clients"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

type server struct {
	authv1.UnimplementedAuthGatewayServiceServer
	cache      Cache
	authClient client.AuthClient
}

/* type server struct {
	authv1.UnimplementedAuthGatewayServiceServer
	cache      Cache
	authClient *client.AuthClient
}
*/

// Create a new server object and inject the cache
func newServer(cache Cache, authClient client.AuthClient) *server {
	return &server{cache: cache, authClient: authClient}
}

//authv1.AuthGatewayServiceServer

/* func NewClients(authClient *client.AuthClient) *server {
	return &server{authClient: authClient}
} */

// Define Cache interface where it is being used

// run the generated GRPC gateway server
func runGateway() error {
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
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
	gwServer := &http.Server{
		Addr:    gatewayAddress,
		Handler: gwmux,
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
	//initiate dependencies for the server -- cache
	lruCache := cache.New(1000)
	authClient := client.NewAuthClient()
	//NewClients(authClient)
	dependencies := newServer(&lruCache, *authClient)

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
