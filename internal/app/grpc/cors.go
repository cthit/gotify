package grpc

import (
	"net/http"
)

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)

			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
				w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE")
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
