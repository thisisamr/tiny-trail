# Stage 1: Build the Go application
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY api/go.mod api/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the .env file into the container
COPY ./api/.env .

# Build the Go application
RUN cd api && go build -o api-service

# Stage 2: Run the Go application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

Copy api/.env .

# Copy the pre-built binary file from the buildlocalhoster stage
COPY --from=builder /app/api/api-service .


# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./api-service"]

