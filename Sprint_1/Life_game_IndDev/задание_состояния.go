package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type State struct {
	Fill int `json:"fill"`
}

func SetState(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("state.cfg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var state State
	err = json.Unmarshal(body, &state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(file, "%d%%", state.Fill)
}

func main() {
	http.HandleFunc("/setstate", SetState)
	http.ListenAndServe(":8081", nil)
}
