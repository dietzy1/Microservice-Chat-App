package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// CORS middleware wrapper that allows origins -- configured in ENV
func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType, Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)

		//the following are some prior configurations used in production
		/* w.Header().Set("Access-Control-Allow-Origin", "https://pepe-api.vercel.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true") */
	})
}

// Reads ENV file and determines if origin should be * or regex matching
func allowedOrigin(origin string) bool {
	if os.Getenv("CORS") == "*" {

		return true
	}
	if matched, _ := regexp.MatchString(os.Getenv(("CORS")), origin); matched {
		return true
	}
	return false
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)

		//I need to make this logger way more visible
	})
}

// Called on the response from the GRPC call - is used to set session token cookies
func withForwardResponseOptionWrapper() runtime.ServeMuxOption {

	ok := runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, m proto.Message) error {
		fmt.Println("withForwardResponseOptionWrapper")
		//I need to read the cookie from the grpc context and set it as a header in the response
		md, ok := runtime.ServerMetadataFromContext(ctx)
		if !ok {
			log.Println("no metadata")
			return nil
		}

		//Specificly look for the session_token key in the metadata
		token := md.HeaderMD.Get("session_token")
		if len(token) == 0 {
			log.Println("no session token")
			return nil
		}

		//Add the session token to the cookie header
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   token[0],
			Expires: time.Now().Add(15 * time.Minute),
			MaxAge:  time.Now().Minute() + 15,
			Path:    "/",
		})
		log.Println("session token set")

		return nil
	})
	return ok
}

// Called before GRPC call is made
// Function that decides which headers to forward to the gRPC server as metadata
func incomingHeaderMatcherWrapper() runtime.ServeMuxOption {
	ok := runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {

		//List of allowed headers that can be forwarded to the gRPC server
		log.Println(key)
		if key == "session_token" {
			log.Default().Println("session_token recieved")
			return key, true
		}

		return runtime.DefaultHeaderMatcher(key)
	})

	return ok
}

// Functuon that looks at the request and extracts the cookies into metadata
func withMetaDataWrapper() runtime.ServeMuxOption {
	ok := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		log.Println("withMetaData called")

		cookie, err := req.Cookie("session_token")
		if err != nil {
			log.Println(err)
			return nil
		}
		log.Println(cookie)

		md := metadata.Pairs("session_token", cookie.Value)

		// ADD THE COOKIE TO THE METADATA

		return md
	})
	return ok
}
