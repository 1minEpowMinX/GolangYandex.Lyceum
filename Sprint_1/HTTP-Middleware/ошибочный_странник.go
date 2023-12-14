package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var name string
var mutex sync.Mutex

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		name = r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ContainsAny(name, "!@#$%^&*()_+.") {
			mutex.Lock()
			name = "dirty hacker"
			mutex.Unlock()
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.Handle("/", SetDefaultName(Sanitize(http.HandlerFunc(HelloHandler))))

	http.ListenAndServe(":8080", nil)
}
