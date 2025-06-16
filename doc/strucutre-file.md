```txt
auth-service/
├── cmd/
│   └── api/
│       └── main.go                    # Punto de entrada de la aplicación
├── internal/
│   ├── core/                          # Capa de dominio (Business Logic)
│   │   ├── domain/                    # Entidades y agregados
│   │   │   ├── user.go
│   │   │   ├── application.go
│   │   │   ├── role.go
│   │   │   ├── module.go
│   │   │   └── oauth_token.go
│   │   ├── ports/                     # Interfaces (contratos)
│   │   │   ├── repositories/          # Interfaces de repositorios
│   │   │   │   ├── user_repository.go
│   │   │   │   ├── application_repository.go
│   │   │   │   ├── role_repository.go
│   │   │   │   └── oauth_repository.go
│   │   │   └── services/              # Interfaces de servicios externos
│   │   │       ├── email_service.go
│   │   │       ├── token_service.go
│   │   │       └── hash_service.go
│   │   └── services/                  # Casos de uso (Use Cases)
│   │       ├── auth_service.go
│   │       ├── user_service.go
│   │       ├── application_service.go
│   │       ├── role_service.go
│   │       └── oauth_service.go
│   ├── adapters/                      # Capa de infraestructura
│   │   ├── handlers/                  # HTTP Handlers (Controllers)
│   │   │   ├── auth_handler.go
│   │   │   ├── user_handler.go
│   │   │   ├── application_handler.go
│   │   │   ├── role_handler.go
│   │   │   └── oauth_handler.go
│   │   ├── repositories/              # Implementaciones de repositorios
│   │   │   ├── postgres/
│   │   │   │   ├── user_repository.go
│   │   │   │   ├── application_repository.go
│   │   │   │   ├── role_repository.go
│   │   │   │   ├── oauth_repository.go
│   │   │   │   └── migrations/
│   │   │   │       ├── 001_create_users_table.go
│   │   │   │       ├── 002_create_applications_table.go
│   │   │   │       ├── 003_create_roles_table.go
│   │   │   │       └── 004_create_oauth_tokens_table.go
│   │   │   └── redis/                 # Cache repository si es necesario
│   │   │       └── session_repository.go
│   │   └── external/                  # Servicios externos
│   │       ├── email/
│   │       │   └── smtp_service.go
│   │       ├── crypto/
│   │       │   └── bcrypt_service.go
│   │       └── jwt/
│   │           └── jwt_service.go
│   ├── infrastructure/                # Configuración de infraestructura
│   │   ├── database/
│   │   │   ├── postgres.go           # Configuración de PostgreSQL
│   │   │   └── redis.go              # Configuración de Redis
│   │   ├── server/
│   │   │   ├── fiber.go              # Configuración de Fiber
│   │   │   ├── routes.go             # Definición de rutas
│   │   │   └── middleware/
│   │   │       ├── auth.go
│   │   │       ├── cors.go
│   │   │       ├── logger.go
│   │   │       └── rate_limit.go
│   │   └── config/
│   │       └── config.go             # Configuración de la aplicación
│   └── shared/                       # Código compartido
│       ├── dto/                      # Data Transfer Objects
│       │   ├── auth_dto.go
│       │   ├── user_dto.go
│       │   ├── application_dto.go
│       │   └── response_dto.go
│       ├── errors/                   # Errores personalizados
│       │   ├── domain_errors.go
│       │   ├── http_errors.go
│       │   └── error_handler.go
│       ├── utils/                    # Utilidades
│       │   ├── validator.go
│       │   ├── pagination.go
│       │   └── response.go
│       └── constants/
│           ├── permissions.go
│           └── status.go
├── pkg/                              # Paquetes públicos reutilizables
│   ├── logger/
│   │   └── logger.go
│   ├── validator/
│   │   └── custom_validators.go
│   └── database/
│       └── connection.go
├── scripts/                          # Scripts de utilidad
│   ├── migration.sh
│   ├── seed.sh
│   └── docker-compose.yml
├── docs/                            # Documentación
│   ├── api/
│   │   └── swagger.yaml
│   └── README.md
├── tests/                           # Tests
│   ├── integration/
│   │   ├── auth_test.go
│   │   └── user_test.go
│   ├── unit/
│   │   ├── services/
│   │   │   ├── auth_service_test.go
│   │   │   └── user_service_test.go
│   │   └── repositories/
│   │       └── user_repository_test.go
│   └── fixtures/
│       └── test_data.go
├── .env.example                     # Variables de entorno ejemplo
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile                         # Comandos de desarrollo
└── README.md
```

# Archivos principales de configuración:

## go.mod

```go
module auth-service

go 1.21

require (
    github.com/gofiber/fiber/v2 v2.52.0
    github.com/gofiber/contrib/jwt v1.0.8
    gorm.io/gorm v1.25.5
    gorm.io/driver/postgres v1.5.4
    github.com/golang-jwt/jwt/v5 v5.2.0
    github.com/redis/go-redis/v9 v9.3.0
    github.com/spf13/viper v1.17.0
    github.com/go-playground/validator/v10 v10.16.0
    golang.org/x/crypto v0.17.0
    github.com/google/uuid v1.4.0
    github.com/stretchr/testify v1.8.4
)
```

## .env.example

# Database

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=auth_service
DB_SSLMODE=disable

# Redis

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT

JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRE_HOURS=24
REFRESH_TOKEN_EXPIRE_HOURS=168

# Server

SERVER_PORT=8080
SERVER_HOST=localhost
ENV=development

# Email

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASSWORD=your-app-password

## Makefile

.PHONY: build run test clean docker-up docker-down migrate seed

# Variables

BINARY_NAME=auth-service
DOCKER_COMPOSE_FILE=docker-compose.yml

# Build the application

build:
`go build -o bin/$(BINARY_NAME) cmd/server/main.go`

# Run the application

run:
`go run cmd/server/main.go`

# Run tests

test:
`go test -v ./...`

# Run tests with coverage

test-coverage:
`go test -v -coverprofile=coverage.out ./...`
`go tool cover -html=coverage.out`

# Clean build artifacts

clean:
`go clean`
`rm -f bin/$(BINARY_NAME)`

# Docker commands

docker-up:
`docker-compose -f $(DOCKER_COMPOSE_FILE) up -d`

docker-down:
`docker-compose -f $(DOCKER_COMPOSE_FILE) down`

docker-build:
`docker build -t $(BINARY_NAME) .`

# Database commands

migrate:
`go run cmd/migrate/main.go`

seed:
`go run cmd/seed/main.go`

# Development

dev:
`air`

# Install dependencies

deps:
`go mod download`
`go mod tidy`

# Lint

lint:
`golangci-lint run`

# Format code

fmt:
`go fmt ./...`
