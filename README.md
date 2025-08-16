# project clearscript-api-gateway

### Branching strategy: 
Trunk-based development
(предполагается внедрение управления фича флагами в будущем)

### Configuration:
задать переменную окружения `CONFIG_PATH` с путем к файлу конфигурации

### Endpoints:

- GET /api/v1/users/{id} - возвращает информацию о пользователе
- GET /api/v1/lessons/{id} - возвращает контент урока с письменными символами

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

