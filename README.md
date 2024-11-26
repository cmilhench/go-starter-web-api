# Golang Standard Library Web Server

## Overview

This project demonstrates a lightweight Go web server that:

- RESTful API design
- In-memory data management

Using only the standard library, this example is designed to give a clear,
beginner-friendly approach to working with web services in Go.

## Features

- Pure Go implementation using standard library
- Modular package structure
- Pattern-based routing
  - Simplified routing with pattern matching
  - More intuitive path value extraction
  - Built-in support for different HTTP methods
  - No need for external routing libraries
- In-memory data store
- JSON serialization/deserialization
- Concurrent-safe operations
- Request validation
- Health check endpoint

## Getting Started

### Prerequisites
- Go 1.22 or higher
- Basic knowledge of Go, HTTP, and JSON.

### Installation

1. Clone the repository
1. Run `go mod tidy` to ensure dependencies are downloaded
1. Run the application with `go run cmd/main.go`

## Example Requests

### Create User
```bash
# List Users
curl http://localhost:8080/api/users

# Get Specific User
curl http://localhost:8080/api/users/1

# Create User
curl -X POST http://localhost:8080/api/users \
     -H "Content-Type: application/json" \
     -d '{"username":"Charlie Brown","email":"charlie@example.com"}'

# Update User
curl -X PUT http://localhost:8080/api/users/1 \
     -H "Content-Type: application/json" \
     -d '{"username":"Alice Updated","email":"alice.updated@example.com"}'

# Delete User
curl -X DELETE http://localhost:8080/api/users/1
```

## Best Practices Demonstrated
- Modular code structure
- Context-aware request handling
- Minimal external dependencies

## Improvements

- Error Handling & Logging
- Middleware Implementation
- External Configuration
- Version Management
- Graceful Shutdown
- Testing
- Rate Limiting
- Metrics & Monitoring
- Development Tools
