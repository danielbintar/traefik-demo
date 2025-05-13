package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/hello", helloHandler)
	http.HandleFunc("/api/v3/hello", helloHandler)

	port := ":8080"
	log.Printf("listen %s", port)
	http.ListenAndServe(port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, world!", "path": r.URL.Path})
}
