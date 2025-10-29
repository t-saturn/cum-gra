# CONFIGURACIÓN
BINARY_NAME := github.com/t-saturn/auth-service-server
APP_NAME := auth-service-server
DOCKER_COMPOSE_FILE := docker-compose.yml
MIGRATE_PATH := internal/adapters/repositories/postgres/migrations
MIGRATE_CMD := go run cmd/migrate/main.go

.PHONY: help build run test clean deps fmt lint dev
.PHONY: docker-up docker-down docker-build docker-logs
.PHONY: migrate-up migrate-down migrate-up-1 migrate-down-1 migrate-force migrate-version migrate-drop
.PHONY: create-migration migrate-status migrate-validate migrate-backup migrate-up-safe
.PHONY: seed test-coverage setup check rebuild

help:
	@echo "=== Central User Manager - Comandos Disponibles ==="
	@echo ""
	@echo " Aplicación:"
	@echo "  build              - Compilar la aplicación"
	@echo "  run                - Ejecutar la aplicación"
	@echo "  dev                - Desarrollo en caliente (requiere air)"
	@echo "  clean              - Limpiar artefactos de compilación"
	@echo ""
	@echo " Desarrollo:"
	@echo "  deps               - Descargar y limpiar dependencias"
	@echo "  fmt                - Formatear código"
	@echo "  lint               - Ejecutar linter"
	@echo ""
	@echo " Keys:"
	@echo "  gen-keys           - Generar claves RSA para JWT"

# COMPILACIÓN Y EJECUCIÓN
build:
	@echo "Compilando aplicación..."
	@go build -o bin/$(APP_NAME) cmd/server/main.go
	@echo "Compilación completada: bin/$(APP_NAME)"
run:
	@echo "Ejecutando aplicación..."
	@go run cmd/server/main.go
dev:
	@echo "Iniciando desarrollo en caliente..."
	@command -v air >/dev/null 2>&1 || { echo "Air no está instalado. Instala con: go install github.com/cosmtrek/air@latest"; exit 1; }
	@air
clean:
	@echo "Limpiando artefactos..."
	@go clean
	@rm -f bin/$(APP_NAME)
	@echo "Limpieza completada"

# DEPENDENCIAS Y FORMATO
deps:
	@echo "Descargando dependencias..."
	@go mod tidy

fmt:
	@echo "Formateando código..."
	@go fmt ./...

lint:
	@echo "Ejecutando linter..."
	@if [ ! -f tools/bin/golangci-lint ]; then \
		echo "golangci-lint no está instalado. Ejecuta: make install-tools"; \
		exit 1; \
	fi
	@tools/bin/golangci-lint run

install-tools:
	@echo "Instalando herramientas locales..."
	@mkdir -p tools/bin
	@GOBIN=$(PWD)/tools/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

gen-keys:
	@echo "Generando claves RSA para JWT..."
	@mkdir -p keys
	@openssl genrsa -out keys/jwtRS256.key 2048
	@openssl rsa -in keys/jwtRS256.key -pubout -out keys/jwtRS256.key.pub
	@echo "Claves generadas en carpeta ./keys"

# COMANDOS COMBINADOS
setup: deps docker-up migrate-up seed
check: fmt lint test
rebuild: clean build
