package main

import (
	"fmt"
	"net/http"
	"sync"
)

type ServerState struct {
	fib_count     int
	metrics_count int
	mutex         sync.Mutex
}

func get_fib(n int) int { // !итеративный метод лучше
	if n < 2 {
		return n
	} else {
		return get_fib(n-1) + get_fib(n-2)
	}
}
func main() {
	state := ServerState{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		state.mutex.Lock()
		defer state.mutex.Unlock()

		fmt.Fprintf(w, "%d", get_fib(state.fib_count))
		state.fib_count++
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		state.mutex.Lock()
		defer state.mutex.Unlock()

		fmt.Fprintf(w, "rpc_duration_milliseconds_count %d\n", state.metrics_count)
		state.metrics_count++
	})

	http.ListenAndServe(":8080", nil)
}
