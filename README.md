# HTTPGraviton

HTTPGraviton is a simple HTTP server written in Go that logs incoming HTTP requests. It is designed for reverse engineering devices in your local network by capturing the routes and payloads they call.

## Features

- Handles all types of HTTP requests (GET, POST, PUT, DELETE)
- Logs request details including URL, method, headers, and body
- Easy to deploy using Docker and Kubernetes

## Usage

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd HTTPGraviton
   ```

2. Build the application:
   ```bash
   go build -o httpgraviton main.go
   ```

3. Run the application:
   ```bash
   ./httpgraviton
   ```

4. Access the server at `http://localhost:8080`.

## Docker

To build and run the application in a Docker container, use the provided Dockerfile.

## Kubernetes

A Kubernetes manifest is provided for easy deployment in a Kubernetes cluster. 