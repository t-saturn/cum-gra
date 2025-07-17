SHELL := /bin/bash

# CONFIGURACIÓN
APP_NAME := central-user-manager
MIGRATE_CMD := go run cmd/migrate/main.go
SEED_CMD := go run cmd/seed/main.go
BACKUP_CMD := go run cmd/backup/main.go
DOCKER_COMPOSE_FILE := docker-compose.yml
GOLANGCI_LINT := ./tools/bin/golangci-lint

GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m

.PHONY: help build run dev clean deps fmt lint docker-up docker-down docker-logs migrate seed backup

# AYUDA
help:
	@echo "$(GREEN)=== Comandos disponibles ===$(NC)"
	@echo "  make run         - Ejecutar servidor"
	@echo "  make dev         - Desarrollo en caliente (requiere air)"
	@echo "  make build       - Compilar aplicación"
	@echo "  make migrate     - Ejecutar migraciones"
	@echo "  make seed        - Ejecutar semillas"
	@echo "  make backup      - Generar backup de la base de datos"
	@echo "  make deps        - Descargar dependencias"
	@echo "  make fmt         - Formatear código"
	@echo "  make lint        - Ejecutar linter"
	@echo "  make docker-up   - Levantar servicios Docker"
	@echo "  make docker-down - Bajar servicios Docker"
	@echo "  make docker-logs - Ver logs de Docker"

# COMANDOS PRINCIPALES
run:
	@echo "$(GREEN) Ejecutando servidor...$(NC)"
	@go run cmd/server/main.go

dev:
	@echo "$(GREEN) Iniciando desarrollo con Air...$(NC)"
	@air

build:
	@echo "$(GREEN) Compilando aplicación...$(NC)"
	@go build -o bin/$(APP_NAME) cmd/server/main.go

clean:
	@echo "$(YELLOW) Limpiando artefactos...$(NC)"
	@go clean
	@rm -rf bin tmp

# MIGRACIONES / SEEDS / BACKUP
migrate:
	@echo "$(GREEN) Ejecutando migraciones...$(NC)"
	@$(MIGRATE_CMD)

seed:
	@echo "$(GREEN) Ejecutando seeds...$(NC)"
	@$(SEED_CMD)

backup:
	@echo "$(GREEN) Ejecutando backup...$(NC)"
	@$(BACKUP_CMD)

# DEPENDENCIAS Y FORMATO
deps:
	@echo "$(GREEN) Descargando dependencias...$(NC)"
	@go mod tidy

fmt:
	@echo "$(GREEN) Formateando código...$(NC)"
	@go fmt ./...

lint:
	@echo Ejecutando linter...
	@if not exist tools\bin\golangci-lint.exe ( \
		echo golangci-lint no está instalado. Ejecuta: make install-tools & \
		exit /b 1 \
	)
	@tools\bin\golangci-lint.exe run

# lint:
# 	@echo "Ejecutando linter..."
# 	@if [ ! -x "$(GOLANGCI_LINT)" ]; then \
# 		echo "golangci-lint no está instalado. Ejecuta: make install-tools"; \
# 		exit 1; \
# 	fi
# 	@$(GOLANGCI_LINT) run
# lint:
# 	@echo "$(GREEN) Ejecutando linter...$(NC)"
# 	@command -v golangci-lint >/dev/null 2>&1 || { echo "$(RED) golangci-lint no está instalado$(NC)"; exit 1; }
# 	@golangci-lint run

install-tools:
	@echo "Instalando herramientas locales..."
	@mkdir -p tools/bin
	@GOBIN=$(PWD)/tools/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# DOCKER
docker-up:
	@echo "$(GREEN) Levantando Docker...$(NC)"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-down:
	@echo "$(YELLOW) Bajando Docker...$(NC)"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

docker-logs:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f
