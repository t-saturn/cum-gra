.PHONY: build run test clean docker-up docker-down migrate seed dev deps fmt lint

# Variables
BINARY_NAME=github.com/t-saturn/central-user-manager # modify this path
DOCKER_COMPOSE_FILE=docker-compose.yml

# ======================================
# Compilar
# ======================================

build:
	go build -o bin/$(BINARY_NAME) cmd/api/main.go

# ======================================
# Ejecutar app
# ======================================

run:
	go run cmd/api/main.go

# ======================================
# Tests
# ======================================

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# ======================================
# Limpiar artefactos de compilación
# ======================================

clean:
	go clean
	rm -f bin/$(BINARY_NAME)

# ======================================
# Docker
# ======================================

docker-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

docker-build:
	docker build -t $(BINARY_NAME) .

# ======================================
# Base de datos
# ======================================

migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/seed/main.go

# ======================================
# Desarrollo en caliente (requiere github.com/cosmtrek/air)
# ======================================

dev:
	air

# ======================================
# Dependencias
# ======================================

deps:
	go mod download
	go mod tidy

# ======================================
# Formateo y lint
# ======================================

fmt:
	go fmt ./...

lint:
	golangci-lint run
