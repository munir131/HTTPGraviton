package main

import (
    "encoding/json"
    "io/ioutil"
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
    // Set the Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Log the request details
    log.Printf("URL: %s", r.URL.Path)
    log.Printf("Method: %s", r.Method)
    log.Printf("Headers: %v", r.Header)

    // Read the request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()
    log.Printf("Body: %s", body)

    // Create a response object
    response := map[string]interface{}{
        "status": "success",
    }

    // Respond with the appropriate status code and JSON response
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
        return
    }

    // Write the JSON response
    json.NewEncoder(w).Encode(response)
}
