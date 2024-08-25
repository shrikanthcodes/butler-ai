# Use the official Golang image as a build stage
FROM golang:1.22.6-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/butler

# Use a minimal image to run the application
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the port that your application runs on
EXPOSE 8080

# Run the Go app
CMD ["./main"]
