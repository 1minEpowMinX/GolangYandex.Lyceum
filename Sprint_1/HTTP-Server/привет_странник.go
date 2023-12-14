package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprint(w, "hello stranger")
		} else if !strings.ContainsAny(name, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz") {
			fmt.Fprint(w, "hello dirty hacker")
		} else {
			fmt.Fprint(w, "hello ", name)
		}
	})

	http.ListenAndServe(":8080", nil)
}
