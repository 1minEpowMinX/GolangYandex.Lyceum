package main

import (
	"fmt"
	"net/http"
)

var requestCount int

func get_fib(n int) int { // !итеративный метод лучше
	if n < 2 {
		return n
	} else {
		return get_fib(n-1) + get_fib(n-2)
	}
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, get_fib(requestCount))
	requestCount++
}
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			requestCount++
		}
		next.ServeHTTP(w, r)
	})
}
