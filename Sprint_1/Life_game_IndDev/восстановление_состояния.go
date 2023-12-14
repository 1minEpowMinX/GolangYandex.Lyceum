package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type State struct {
	Fill int `json:"fill"`
}

func Reset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	bytes, err := os.ReadFile("state.cfg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var state State
	n, err := strconv.Atoi(strings.Trim(string(bytes), "%%\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state.Fill = n

	jsonData, err := json.Marshal(&state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/reset", Reset)
	http.ListenAndServe(":8081", nil)
}
