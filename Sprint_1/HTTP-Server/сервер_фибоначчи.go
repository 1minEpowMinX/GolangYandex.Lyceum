package main

import (
	"fmt"
	"net/http"
	"sync"
)

type ServerState struct {
	count int
	mutex sync.Mutex
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

		fmt.Fprint(w, get_fib(state.count))
		state.count++
	})

	http.ListenAndServe(":8080", nil)
}
