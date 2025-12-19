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
	@echo "=== Central User Manager - Comandos Disponibles ==="
	@echo ""
	@echo " Aplicación:"
	@echo "  build              - Compilar la aplicación"
	@echo "  run                - Ejecutar el servidor principal"
	@echo "  dev                - Desarrollo en caliente (requiere air)"
	@echo "  clean              - Limpiar artefactos de compilación"
	@echo ""
	@echo " Testing:"
	@echo "  test               - Ejecutar tests"
	@echo "  test-coverage      - Ejecutar tests con cobertura"
	@echo ""
	@echo " Docker:"
	@echo "  docker-up          - Levantar servicios con Docker Compose"
	@echo "  docker-down        - Bajar servicios Docker"
	@echo "  docker-logs        - Ver logs de Docker"
	@echo ""
	@echo " Base de Datos:"
	@echo "  migrate-up         - Ejecutar migración UP"
	@echo "  migrate-down       - Revertir migración DOWN"
	@echo "  migrate-reset      - Ejecutar limpieza de datos (RESET)"
	@echo "  migrate-force      - Forzar base de datos a una versión específica (usa VERSION=N)"
	@echo "  migrate-version    - Mostrar versión actual de la base de datos"
	@echo "  migrate-drop       - Eliminar todas las tablas migradas (peligroso)"
	@echo "  reset-data         - Vaciar datos de todas las tablas (sin eliminar estructuras)"
	@echo ""
	@echo " Seeds & Backups:"
	@echo "  seed               - Ejecutar semillas"
	@echo "  backup             - Generar backup de la base de datos"
	@echo ""
	@echo " Desarrollo:"
	@echo "  deps               - Descargar y limpiar dependencias"
	@echo "  fmt                - Formatear código"
	@echo "  lint               - Ejecutar linter"
	@echo ""
	@echo " Ejemplos:"
	@echo "  make migrate-force VERSION=001"

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
	@$(MIGRATE_CMD) -cmd=up

migrate-down:
	@echo "Revirtiendo migración DOWN..."
	@$(MIGRATE_CMD) -cmd=down

migrate-reset:
	@echo "Ejecutando limpieza de datos (RESET)..."
	@$(MIGRATE_CMD) -cmd=reset

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
