# Start with the official Golang image
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the application port
EXPOSE 3000

# installs air
RUN go install github.com/air-verse/air@latest

# installs delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Command to run the application
CMD ["air"] 
# "dlv", "debug", "--listen=0.0.0.0:3030", "--headless=true", "--api-version=2", "--accept-multiclient",  "--continue"]
