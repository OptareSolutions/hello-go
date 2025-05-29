package main

import (
    "encoding/json"
    "net/http"
    "os"
)

type response struct {
    Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    resp := response{Message: "Hello World desde Go!"}
    json.NewEncoder(w).Encode(resp)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":"+port, nil)
}