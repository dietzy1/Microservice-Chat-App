package server

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func withForwardResponseOptionWrapper() runtime.ServeMuxOption {
	ok := runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, m proto.Message) error {

		//I need to read the cookie from the grpc context and set it as a header in the response

		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: "value",
		})

		return nil
	})
	return ok
}

func incomingHeaderMatcherWrapper() runtime.ServeMuxOption {
	ok := runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {

		log.Default().Println(key)
		if key == "session_token" {

			log.Default().Println("session_token recieved")
			return key, true
		}
		return runtime.DefaultHeaderMatcher(key)
	})
	return ok
}

func withMetaDataWrapper() runtime.ServeMuxOption {
	ok := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {

		cookie, err := req.Cookie("session_token")
		if err != nil {
			log.Default().Println(err)
		}
		log.Default().Println(cookie)

		md := metadata.Pairs("session_token", "value")

		// ADD THE COOKIE TO THE METADATA

		return md
	})
	return ok
}
