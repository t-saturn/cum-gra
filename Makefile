SHELL := /bin/bash

# CONFIGURACIÓN
APP_NAME := central-user-manager
MIGRATE_CMD := go run cmd/migrate/main.go
SEED_CMD := go run cmd/seed/main.go
BACKUP_CMD := go run cmd/backup/main.go
DOCKER_COMPOSE_FILE := docker-compose.yml
GOLANGCI_LINT := ./tools/bin/golangci-lint

.PHONY: help build run dev clean deps fmt lint docker-up docker-down docker-logs migrate seed backup

# AYUDA
help:
	@echo "=== Comandos disponibles ==="
	@echo "  make run              - Ejecutar servidor"
	@echo "  make dev              - Desarrollo en caliente (requiere air)"
	@echo "  make build            - Compilar aplicación"
	@echo "  make migrate-up       - Ejecutar migración UP"
	@echo "  make migrate-down     - Revertir migración DOWN"
	@echo "  make reset-data       - Vaciar datos de todas las tablas (sin eliminar estructuras)"
	@echo "  make migrate-reset    - Ejecutar limpieza de datos sin borrar estructuras"
	@echo "  make migrate-force    - Forzar base de datos a una versión específica (usa -version=N)"
	@echo "  make migrate-version  - Mostrar versión actual de la base de datos"
	@echo "  make migrate-drop     - Eliminar todas las tablas migradas (peligroso)"
	@echo "  make seed             - Ejecutar semillas"
	@echo "  make backup           - Generar backup de la base de datos"
	@echo "  make deps             - Descargar dependencias"
	@echo "  make fmt              - Formatear código"
	@echo "  make lint             - Ejecutar linter"
	@echo "  make docker-up        - Levantar servicios Docker"
	@echo "  make docker-down      - Bajar servicios Docker"
	@echo "  make docker-logs      - Ver logs de Docker"

# COMANDOS PRINCIPALES
run:
	@echo "Ejecutando servidor..."
	@go run cmd/server/main.go

dev:
	@echo "Iniciando desarrollo con Air..."
	@air

build:
	@echo "Compilando aplicación..."
	@go build -o bin/$(APP_NAME) cmd/server/main.go

clean:
	@echo "Limpiando artefactos..."
	@go clean
	@rm -rf bin tmp

# MIGRACIONES / SEEDS / BACKUP
migrate-up:
	@echo "Ejecutando migración UP..."
	@$(MIGRATE_CMD) -cmd=up -path=internal/database/migrations

migrate-down:
	@echo "Revirtiendo migración DOWN..."
	@$(MIGRATE_CMD) -cmd=down -path=internal/database/migrations

reset-data:
	@echo "Vaciando datos de todas las tablas..."
	@psql $(DB_URL) -f internal/database/clean/reset_data.sql

migrate-reset:
	@echo "Ejecutando limpieza de datos (RESET)..."
	@$(MIGRATE_CMD) -cmd=reset -path=internal/database/migrations

migrate-force:
	@echo "Forzando migración a versión $(VERSION)..."
	@$(MIGRATE_CMD) -cmd=force -version=$(VERSION) -path=internal/database/migrations

migrate-version:
	@echo "Mostrando versión actual de migraciones..."
	@$(MIGRATE_CMD) -cmd=version -path=internal/database/migrations

migrate-drop:
	@echo "Eliminando todas las tablas migradas..."
	@$(MIGRATE_CMD) -cmd=drop -path=internal/database/migrations

seed:
	@echo "Ejecutando seeds..."
	@$(SEED_CMD)

backup:
	@echo "Ejecutando backup..."
	@$(BACKUP_CMD)

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

# DOCKER
docker-up:
	@echo "Levantando Docker..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-down:
	@echo "Bajando Docker..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

docker-logs:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f
