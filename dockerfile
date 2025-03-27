# Start with the official Golang image
FROM golang:1.23 as builder

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

# Command to run the application
CMD ["air"]