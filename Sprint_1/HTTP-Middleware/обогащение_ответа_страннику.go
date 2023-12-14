package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Server struct {
	Greetings string     `json:"greetings"`
	Name      string     `json:"name"`
	Mutex     sync.Mutex `json:"-"`
}

type RpcResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

var server *Server // ! можно отказаться от структуры и написать через мапу, а также отказаться от JSON-нов
var rpc *RpcResponse

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rpc = &RpcResponse{Status: "error", Result: struct{}{}}
			}

			jsonData, err := json.Marshal(rpc)
			if err != nil {
				return
			}

			fmt.Fprint(w, string(jsonData))
		}()
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.Mutex.Lock()
		server.Name = r.URL.Query().Get("name")
		if server.Name == "" {
			server.Name = "stranger"
		}
		server.Mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ContainsAny(server.Name, "!@#$%^&*()_+.") {
			panic("Invalid name")
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	server.Greetings = "hello"
	rpc = &RpcResponse{Status: "ok", Result: *server}
}

func main() { // ! не прошло в ЛМС, но работает
	server = &Server{}
	rpc = &RpcResponse{}
	http.Handle("/hello", RPC(SetDefaultName(Sanitize(http.HandlerFunc(HelloHandler)))))
	http.ListenAndServe(":8080", nil)
}
