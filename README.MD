![tiny-trail-banner](https://github.com/user-attachments/assets/aa4e2525-4829-46b6-8e77-e2abba432e65)

# Tiny-Trail: A URL Shortening Service

Tiny-Trail is a URL shortening service built with Go, using Fiber as the web framework and Redis for data storage. It provides functionality for creating short, custom URLs and includes rate limiting to manage excessive requests. The service can run both as a standalone application or in Docker containers for easy deployment.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [Configuration](#configuration)
- [Scaling](#scaling)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features

- Shorten long URLs with custom aliases.
- Rate limiting to prevent abuse of the service.
- Redis for fast storage and retrieval.
- Easy setup with Docker support.

### Getting Started

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/thisisamr/tiny-trail.git
   cd tiny-trail/api
   ```

2. **Install Dependencies**

   Make sure you have Go installed. Then, install dependencies by running:

   ```bash
   go mod tidy
   ```

3. **Install Air for Development**

   If you don't have Air installed for hot-reloading, install it using:

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

### Running the Application

#### Option 1: Running Locally

1. **Run the Redis Server**

   Make sure Redis is running on your local machine or use Docker:

   ```bash
   docker run --name redis-db -p 6379:6379 -d redis
   ```

2. **Start the Server**

   Run the Go application locally:

   ```bash
   air
   ```

#### Option 2: Running with Docker

1. **Build and Run Docker Containers**

   Use Docker Compose to build and run the application:

   ```bash
   docker-compose up --build
   ```

### Configuration

To configure the application, create a `.env` file in the `api` directory. Use the following example:

```text
DB_ADDRESS=redis-db-service:6379
DB_PASSWORD=apollo123
APP_PORT=3000
DOMAIN=localhost:3000
```

- `DB_ADDRESS`: Address of the Redis server.
- `DB_PASSWORD`: Password for the Redis server.
- `APP_PORT`: Port on which the application will run.
- `DOMAIN`: Domain for the service.

## Scaling

Tiny-Trail is designed to handle high traffic. To scale:

- Use Docker to deploy multiple instances of the Go service behind a load balancer.
- Scale Redis separately depending on the load (consider Redis Cluster for high availability).
- Use Kubernetes for orchestrating containers if needed.

## Project Structure

The project is structured as follows:

```text
tiny-trail/
├── api
│   ├── db
│   │   └── db.go                # Redis client setup
│   ├── Dockerfile               # Dockerfile for Go API
│   ├── go.mod                   # Go module file
│   ├── go.sum                   # Go dependency file
│   ├── main.go                  # Entry point of the Go application
│   ├── Middleware
│   │   ├── rateLimiting.go      # Rate limiting middleware
│   │   └── validation.go        # Validation middleware
│   ├── routes
│   │   ├── clipurl.go           # Handler for shortening URLs
│   │   └── resolve.go           # Handler for resolving URLs
│   └── tmp
│       └── main                 # Temporary directory for development
├── docker-compose.yml           # Docker Compose configuration
├── README.md                    # Project README
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.