package main

import (
	"context"
	"fmt"
	"net/http"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req_id := r.Header.Get("X-Request-ID")
		if req_id != "" {
			req_id = "RequestID: " + req_id
		} else {
			req_id = "RequestID not found"
		}
		ctx := context.WithValue(r.Context(), "X-Request-ID", req_id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	req_id := r.Context().Value("X-Request-ID")
	fmt.Fprintf(w, "Hello! %s", req_id)
}

func main() {
	http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))

	http.ListenAndServe(":8080", nil)
}
