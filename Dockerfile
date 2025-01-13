# Use a lightweight Go image
FROM golang:1.23.3-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application (specify the entry point in cmd/main.go)
RUN go build -o bin/core ./cmd/core/main.go

# Expose the application port
EXPOSE 8090

# Command to run the application with the configuration file
CMD ["./bin/core", "--config=./config/server.yaml"]
