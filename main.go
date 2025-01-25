package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Read the port from the environment variable or use default
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port
    }

    http.HandleFunc("/", handleRequest)
    log.Printf("Starting server on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    // Log the request details
    log.Printf("URL: %s", r.URL.Path)
    log.Printf("Method: %s", r.Method)
    log.Printf("Headers: %v", r.Header)

    // Read the request body
    body := make([]byte, r.ContentLength)
    _, err := r.Body.Read(body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Printf("Body: %s", body)

    // Respond with the appropriate status code
    switch r.Method {
    case http.MethodGet:
        w.WriteHeader(http.StatusOK)
    case http.MethodPost:
        w.WriteHeader(http.StatusCreated)
    case http.MethodPut:
        w.WriteHeader(http.StatusNoContent)
    case http.MethodDelete:
        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
