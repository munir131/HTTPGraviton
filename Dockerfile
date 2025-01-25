# Use the official Golang image as a build stage
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifest
COPY go.mod ./

# Download the Go modules (this will be empty if there are no external dependencies)
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpgraviton main.go

# Use a minimal base image for the final image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/httpgraviton /bin/httpgraviton

# Set the default port environment variable
ENV PORT=8080

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["httpgraviton"] 