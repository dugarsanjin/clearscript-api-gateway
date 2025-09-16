# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based API gateway called `clearscript-api-gateway`. The project follows trunk-based development workflow with plans to implement feature flag management in the future.

## Architecture

### Core Structure
- **Entry Point**: `cmd/app/main.go` - Sets up HTTP server with Chi router, logging, and middleware
- **Configuration**: `internal/config/config.go` - Uses cleanenv for YAML/environment variable configuration  
- **HTTP Layer**: `internal/http/` - Contains handlers and middleware
  - **Handlers**: Structured handlers in `internal/http/handlers/` using handler pattern with methods returning `http.HandlerFunc`
  - **User Handler**: `internal/http/handlers/user.go` - User-related endpoints
  - **Content Handler**: `internal/http/handlers/content.go` - Lesson content endpoints
- **Middleware**: Custom logging middleware in `internal/http/middleware/logger/`

### Configuration System
The application requires a `CONFIG_PATH` environment variable pointing to a YAML configuration file. Local development uses `config/local.yml`.

### Logging
Uses Go's structured logging (slog) with environment-specific handlers:
- **local**: Text format with debug level
- **dev**: JSON format with debug level  
- **prod**: JSON format with info level

### HTTP Framework
Built on Chi router (go-chi/chi/v5) with standard middleware:
- Request ID generation
- Custom structured logging
- Panic recovery
- URL formatting

## Development Commands

### Build and Run
```bash
# Build the application
go build -o bin/clearscript-api-gateway ./cmd/app

# Run directly
CONFIG_PATH=config/local.yml go run ./cmd/app

# Format code
go fmt ./...

# Download dependencies  
go mod download

# Tidy dependencies
go mod tidy
```

### Testing
```bash
# Run tests (if any exist)
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/config
```

### Swagger Documentation
```bash
# Generate Swagger documentation from code annotations
swag init -g cmd/app/main.go

# Install swag tool (if not already installed)
go install github.com/swaggo/swag/cmd/swag@latest

# Access Swagger UI (when server is running)
# http://localhost:8080/clearscript-api-gateway/swagger/index.html
```

### Docker Commands
```bash
# Build Docker image
docker build -t clearscript-api-gateway .

# Run with Docker (not recommended, use docker-compose)
docker run -p 8080:8080 \
  -v $(pwd)/config:/root/config:ro \
  -e CONFIG_PATH=/root/config/local.yml \
  clearscript-api-gateway

# Using docker-compose (recommended)
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f clearscript-api-gateway

# Rebuild and restart
docker-compose up --build -d

# Different environments:
# Local development (default)
docker-compose up -d

# Development environment  
CONFIG_PATH=/root/config/dev.yml docker-compose up -d

# Production with external config
CONFIG_PATH=/root/config/prod.yml HOST_CONFIG_PATH=/etc/clearscript/config docker-compose up -d

# Using .env file (copy .env.example to .env and modify)
cp .env.example .env
docker-compose up -d
```

## Project Dependencies
- **go-chi/chi/v5**: HTTP router and middleware
- **ilyakaznacheev/cleanenv**: Configuration management
- **BurntSushi/toml**: TOML parsing support
- **joho/godotenv**: Environment variable loading
- **swaggo/swag**: Swagger documentation generation
- **swaggo/http-swagger**: Swagger UI middleware for HTTP servers
- **swaggo/files**: Static file handling for Swagger

## Project Files
- **.gitignore**: Excludes build artifacts, IDE files, logs, and sensitive configs from Git
- **.dockerignore**: Optimizes Docker build by excluding unnecessary files
- **.env.example**: Template for environment variables (copy to .env)

## API Documentation
The project includes Swagger/OpenAPI documentation:
- **Swagger UI**: Available at `/swagger/index.html` when server is running
- **API Spec**: Auto-generated from code annotations using swaggo/swag
- **Endpoints**: 
  - `GET /api/v1/users/{id}` - Retrieve user information
  - `GET /api/v1/lessons/{id}` - Retrieve lesson content

## Development Notes
- Handler architecture uses structured approach: each resource has its own handler struct with methods returning `http.HandlerFunc`
- Response types are prefixed with resource name (e.g., `UserGetResponse`, `ContentGetResponse`) to avoid naming conflicts
- Mock responses are implemented for all endpoints during development
- No tests are currently present in the codebase
- Swagger documentation is auto-generated from code annotations in handler methods

## Docker Configuration Notes
- **Important**: For Docker deployment, use `address: "0.0.0.0:8080"` in config files (not `localhost:8080`)
- Config files are mounted as volumes, allowing changes without rebuilding the image
- Use `HOST_CONFIG_PATH` environment variable to specify external config directory for production
- Default configuration uses files from the repository (`./config` directory)