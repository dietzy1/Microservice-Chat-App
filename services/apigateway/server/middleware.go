package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// CORS middleware wrapper that allows origins -- configured in ENV
func cors(h http.Handler) http.Handler {
	log.Println("CORS WAS HIT")
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

func recoverer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func wrapperAuthMiddleware(authClient authclientv1.AuthServiceClient) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//if the request path is /login or /register /authenticate, then skip the auth middleware

			fmt.Println("request path: ", r.URL.Path)

			if r.URL.Path == "/v1/auth/login" || r.URL.Path == "/v1/auth/register" || r.URL.Path == "/v1/auth/authenticate" {
				h.ServeHTTP(w, r)
				return
			}
			//Unsure if I actually need to do the cookie thing here or if its done at some other part of the middleware

			//Call the auth service to check if the session token is valid
			/* cookie, err := r.Cookie("session_token")
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			//Call the auth service to check if the session token is valid
			ctx := context.Background()
			ctx = metadata.AppendToOutgoingContext(ctx, "session_token", cookie.Value)
			_, err = authClient.Authenticate(ctx, &authclientv1.AuthenticateRequest{})
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			} */

			h.ServeHTTP(w, r)

		})
	}
}

// Called on the response from the GRPC call - is used to set session token cookies
func withForwardResponseOptionWrapper() runtime.ServeMuxOption {

	ok := runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, m proto.Message) error {
		//I need to read the cookie from the grpc context and set it as a header in the response
		md, ok := runtime.ServerMetadataFromContext(ctx)
		if !ok {
			log.Println("no metadata")
			return nil
		}
		fmt.Println("2nd hit")

		//Specificly look for the session_token key in the metadata
		token := md.HeaderMD.Get("session_token")
		if len(token) == 0 {
			log.Println("no session token")
			return nil
		}

		//perform check if token is set to "logout" and if so, delete the cookie
		if token[0] == "" {
			log.Println("Deleting cookie")
			// Add the session token to the cookie header
			http.SetCookie(w, &http.Cookie{
				Name:  "session_token",
				Value: "",

				MaxAge: -1,
			})
			return nil
		}

		//Add the session token to the cookie header
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   token[0],
			Expires: time.Now().Add(60 * time.Minute),
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
