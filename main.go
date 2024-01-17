package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	port := 8080
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		fmt.Println(err)
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	response := map[string]string{"message": "pong"}

	json.NewEncoder(w).Encode(response)
}
