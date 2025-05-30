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
    http.Error(w, "Error de prueba: handler roto", http.StatusInternalServerError)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":"+port, nil)
}