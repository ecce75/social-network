# Use the official Golang image as the base image
FROM golang:1.21.2-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Install gcc and other necessary build tools
RUN apk update && apk add --no-cache gcc musl-dev

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Enable CGO for the build
ENV CGO_ENABLED=1

# Build the Go application
RUN go build -o main server.go

# Set up a smaller base image for the runtime container
FROM arm64v8/alpine:latest  

# Install runtime dependencies for CGO and SQLite
RUN apk add --no-cache libc6-compat

# Set the working directory inside the runtime container
WORKDIR /root/

# Copy the compiled executable from the builder to the runtime container
COPY --from=builder /app/main .

# Copy the pkg directory with database files and migrations
COPY --from=builder /app/pkg/db /root/pkg/db

# Expose the port on which the application will run
EXPOSE 8080

# Check the contents of the current directory (for debugging)
RUN ls -l

# Ensure the main executable has execution permissions
RUN chmod +x main

# Run the Go application
CMD ["./main"]