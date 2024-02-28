# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the environment variables
ENV DB_NAME=database.db
ENV DB_PATH=backend/pkg/db/

# Create the data directory
RUN mkdir -p $DB_PATH

# Expose the port on which the application will run
# 8080 443
EXPOSE 8080

# Run the Go application
CMD ["./main"]
