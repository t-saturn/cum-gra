# CONFIGURACIÓN
# Variables principales
BINARY_NAME := github.com/t-saturn/central-user-manager
APP_NAME := central-user-manager
DOCKER_COMPOSE_FILE := docker-compose.yml
MIGRATE_PATH := internal/adapters/repositories/postgres/migrations
MIGRATE_CMD := go run cmd/migrate/main.go
# Colores para output
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m # No Color
.PHONY: help build run test clean deps fmt lint dev
.PHONY: docker-up docker-down docker-build docker-logs
.PHONY: migrate-up migrate-down migrate-up-1 migrate-down-1 migrate-force migrate-version migrate-drop
.PHONY: create-migration migrate-status migrate-validate migrate-backup migrate-up-safe
.PHONY: seed test-coverage

# AYUDA Y COMANDOS PRINCIPALES
# Comando por defecto
help:
	@echo "$(GREEN)=== Central User Manager - Comandos Disponibles ===$(NC)"
	@echo ""
	@echo "$(YELLOW) Aplicación:$(NC)"
	@echo "  build              - Compilar la aplicación"
	@echo "  run                - Ejecutar la aplicación"
	@echo "  dev                - Desarrollo en caliente (requiere air)"
	@echo "  clean              - Limpiar artefactos de compilación"
	@echo ""
	@echo "$(YELLOW) Testing:$(NC)"
	@echo "  test               - Ejecutar tests"
	@echo "  test-coverage      - Ejecutar tests con cobertura"
	@echo ""
	@echo "$(YELLOW) Docker:$(NC)"
	@echo "  docker-up          - Levantar servicios con Docker Compose"
	@echo "  docker-down        - Bajar servicios Docker"
	@echo "  docker-build       - Construir imagen Docker"
	@echo "  docker-logs        - Ver logs de Docker Compose"
	@echo ""
	@echo "$(YELLOW)🗄️  Base de Datos:$(NC)"
	@echo "  migrate-up         - Ejecutar todas las migraciones"
	@echo "  migrate-down       - Revertir todas las migraciones"
	@echo "  migrate-up-1       - Ejecutar 1 migración"
	@echo "  migrate-down-1     - Revertir 1 migración"
	@echo "  migrate-status     - Ver estado de migraciones"
	@echo "  create-migration   - Crear nueva migración (NAME=nombre)"
	@echo "  seed               - Ejecutar seeds de datos"
	@echo ""
	@echo "$(YELLOW)🔧 Desarrollo:$(NC)"
	@echo "  deps               - Descargar y limpiar dependencias"
	@echo "  fmt                - Formatear código"
	@echo "  lint               - Ejecutar linter"
	@echo ""
	@echo "$(YELLOW)Ejemplos:$(NC)"
	@echo "  make create-migration NAME=add_users_table"
	@echo "  make migrate-force VERSION=001"

# COMPILACIÓN Y EJECUCIÓN
build:
	@echo "$(GREEN) Compilando aplicación...$(NC)"
	@go build -o bin/$(APP_NAME) cmd/server/main.go
	@echo "$(GREEN) Compilación completada: bin/$(APP_NAME)$(NC)"
run:
	@echo "$(GREEN) Ejecutando aplicación...$(NC)"
	@go run cmd/server/main.go
dev:
	@echo "$(GREEN) Iniciando desarrollo en caliente...$(NC)"
	@command -v air >/dev/null 2>&1 || echo "$(RED) Air no está instalado. Instala con: go install github.com/cosmtrek/air@latest$(NC)"; exit 1;
	@air
clean:
	@echo "$(YELLOW) Limpiando artefactos...$(NC)"
	@go clean
	@rm -f bin/$(APP_NAME)
	@echo "$(GREEN) Limpieza completada$(NC)"

# TESTING
test:
	@echo "$(GREEN) Ejecutando tests...$(NC)"
	@go test -v ./...
test-coverage:
	@echo "$(GREEN) Ejecutando tests con cobertura...$(NC)"
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@echo "$(GREEN) Reporte de cobertura generado: coverage.out$(NC)"

# DOCKER
docker-up:
	@echo "$(GREEN) Levantando servicios Docker...$(NC)"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "$(GREEN) Servicios Docker iniciados$(NC)"
docker-down:
	@echo "$(YELLOW) Bajando servicios Docker...$(NC)"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "$(GREEN) Servicios Docker detenidos$(NC)"
docker-build:
	@echo "$(GREEN) Construyendo imagen Docker...$(NC)"
	@docker build -t $(APP_NAME) .
	@echo "$(GREEN) Imagen Docker construida: $(APP_NAME)$(NC)"
docker-logs:
	@echo "$(GREEN) Mostrando logs de Docker Compose...$(NC)"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# MIGRACIONES DE BASE DE DATOS
migrate-up:
	@echo "$(GREEN)  Ejecutando migraciones...$(NC)"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=up
	@echo "$(GREEN) Migraciones completadas$(NC)"
migrate-down:
	@echo "$(YELLOW)  Revirtiendo migraciones...$(NC)"
	@echo "$(RED)  CUIDADO: Esto revertirá TODAS las migraciones$(NC)"
	@read -p "¿Continuar? [y/N]: " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=down; \
		echo "$(GREEN) Migraciones revertidas$(NC)"; \
	else \
		echo "$(YELLOW) Operación cancelada$(NC)"; \
	fi
migrate-up-1:
	@echo "$(GREEN)  Ejecutando 1 migración...$(NC)"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=up -steps=1
migrate-down-1:
	@echo "$(YELLOW)  Revirtiendo 1 migración...$(NC)"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=down -steps=1
migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "$(RED) Error: Especifica VERSION=X$(NC)"; \
		echo "$(YELLOW)Ejemplo: make migrate-force VERSION=001$(NC)"; \
		exit 1; \
	fi
	@echo "$(YELLOW) Forzando migración a versión $(VERSION)...$(NC)"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=force -version=$(VERSION)
migrate-version:
	@echo "$(GREEN) Versión actual de migración:$(NC)"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=version
migrate-drop:
	@echo "$(RED)  PELIGRO: Esto eliminará TODAS las tablas$(NC)"
	@read -p "¿Estás COMPLETAMENTE seguro? [y/N]: " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=drop; \
		echo "$(GREEN) Tablas eliminadas$(NC)"; \
	else \
		echo "$(YELLOW) Operación cancelada$(NC)"; \
	fi
create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "$(RED) Error: Especifica NAME=nombre_de_la_migracion$(NC)"; \
		echo "$(YELLOW)Ejemplo: make create-migration NAME=add_users_table$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN) Creando migración: $(NAME)$(NC)"
	@TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	UP_FILE="$(MIGRATE_PATH)/$${TIMESTAMP}_$(NAME).up.sql"; \
	DOWN_FILE="$(MIGRATE_PATH)/$${TIMESTAMP}_$(NAME).down.sql"; \
	mkdir -p $(MIGRATE_PATH); \
	echo "-- Migration: $(NAME)" > $$UP_FILE; \
	echo "-- Created: $$(date)" >> $$UP_FILE; \
	echo "" >> $$UP_FILE; \
	echo "-- Add your UP migration here" >> $$UP_FILE; \
	echo "" >> $$UP_FILE; \
	echo "-- Migration: $(NAME)" > $$DOWN_FILE; \
	echo "-- Created: $$(date)" >> $$DOWN_FILE; \
	echo "" >> $$DOWN_FILE; \
	echo "-- Add your DOWN migration here" >> $$DOWN_FILE; \
	echo "" >> $$DOWN_FILE; \
	echo "$(GREEN) Migración creada:$(NC)"; \
	echo "   UP:   $$UP_FILE"; \
	echo "   DOWN: $$DOWN_FILE"
migrate-status:
	@echo "$(GREEN)=== Estado de Migraciones ===$(NC)"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=version
	@echo ""
	@echo "$(GREEN)=== Archivos de migración disponibles ===$(NC)"
	@ls -la $(MIGRATE_PATH)/*.sql 2>/dev/null | head -10 || echo "$(YELLOW)No hay archivos de migración$(NC)"
	@TOTAL=$$(ls $(MIGRATE_PATH)/*.sql 2>/dev/null | wc -l); \
	if [ $$TOTAL -gt 10 ]; then \
		echo "$(YELLOW)... y $$(($$TOTAL - 10)) archivos más$(NC)"; \
	fi
migrate-validate:
	@echo "$(GREEN) Validando migraciones...$(NC)"
	@ERROR=0; \
	for file in $(MIGRATE_PATH)/*.up.sql; do \
		if [ -f "$$file" ]; then \
			basename="$$(basename $$file .up.sql)"; \
			down_file="$(MIGRATE_PATH)/$${basename}.down.sql"; \
			if [ ! -f "$$down_file" ]; then \
				echo "$(RED) Error: No existe el archivo DOWN para $$file$(NC)"; \
				ERROR=1; \
			fi; \
		fi; \
	done; \
	if [ $$ERROR -eq 0 ]; then \
		echo "$(GREEN) Todas las migraciones son válidas$(NC)"; \
	else \
		exit 1; \
	fi
migrate-backup:
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "$(RED) Error: DATABASE_URL no está definida$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN) Creando backup de la base de datos...$(NC)"
	@TIMESTAMP=$$(date +%Y%m%d_%H%M%S); \
	BACKUP_FILE="backups/db_backup_$${TIMESTAMP}.sql"; \
	mkdir -p backups; \
	pg_dump $(DATABASE_URL) > $$BACKUP_FILE 2>/dev/null; \
	if [ $$? -eq 0 ]; then \
		echo "$(GREEN) Backup creado: $$BACKUP_FILE$(NC)"; \
	else \
		echo "$(RED) Error creando backup$(NC)"; \
		exit 1; \
	fi
migrate-up-safe: migrate-backup migrate-up
	@echo "$(GREEN) Migración completada con backup de seguridad$(NC)"

# DATOS DE PRUEBA
seed:
	@echo "$(GREEN) Ejecutando seeds...$(NC)"
	@go run cmd/seed/main.go
	@echo "$(GREEN) Seeds completados$(NC)"

# HERRAMIENTAS DE DESARROLLO
deps:
	@echo "$(GREEN) Descargando dependencias...$(NC)"
	@go mod download
	@go mod tidy
	@echo "$(GREEN) Dependencias actualizadas$(NC)"
fmt:
	@echo "$(GREEN) Formateando código...$(NC)"
	@go fmt ./...
	@echo "$(GREEN) Código formateado$(NC)"
lint:
	@echo "$(GREEN) Ejecutando linter...$(NC)"
	@command -v golangci-lint >/dev/null 2>&1 || { echo "$(RED) golangci-lint no está instalado$(NC)"; exit 1; }
	@golangci-lint run
	@echo "$(GREEN) Linting completado$(NC)"

# COMANDOS COMBINADOS
# Preparar entorno completo
setup: deps docker-up migrate-up seed
	@echo "$(GREEN)🎉 Entorno configurado completamente$(NC)"
# Verificación completa
check: fmt lint test
	@echo "$(GREEN) Verificación completa exitosa$(NC)"
# Reconstruir todo
rebuild: clean build
	@echo "$(GREEN) Reconstrucción completada$(NC)"
