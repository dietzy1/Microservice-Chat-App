package server

import (
	"net/http"
)

func serveSwaggerUI(pathToSwaggerUI string) http.Handler {
	swaggerUI := http.FileServer(http.Dir(pathToSwaggerUI))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swagger-ui/" {
			http.Redirect(w, r, "/swagger-ui/index.html", http.StatusFound)
			return
		}
		swaggerUI.ServeHTTP(w, r)
	})
}
