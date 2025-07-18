# CONFIGURACIÓN
BINARY_NAME := github.com/t-saturn/central-user-manager
APP_NAME := central-user-manager
DOCKER_COMPOSE_FILE := docker-compose.yml
MIGRATE_PATH := internal/adapters/repositories/postgres/migrations
MIGRATE_CMD := go run cmd/migrate/main.go

.PHONY: help build run test clean deps fmt lint dev
.PHONY: docker-up docker-down docker-build docker-logs
.PHONY: migrate-up migrate-down migrate-up-1 migrate-down-1 migrate-force migrate-version migrate-drop
.PHONY: create-migration migrate-status migrate-validate migrate-backup migrate-up-safe
.PHONY: seed test-coverage setup check rebuild

# AYUDA Y COMANDOS PRINCIPALES
help:
	@echo "=== Central User Manager - Comandos Disponibles ==="
	@echo ""
	@echo " Aplicación:"
	@echo "  build              - Compilar la aplicación"
	@echo "  run                - Ejecutar la aplicación"
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
	@echo "  docker-build       - Construir imagen Docker"
	@echo "  docker-logs        - Ver logs de Docker Compose"
	@echo ""
	@echo " Base de Datos:"
	@echo "  migrate-up         - Ejecutar todas las migraciones"
	@echo "  migrate-down       - Revertir todas las migraciones"
	@echo "  migrate-up-1       - Ejecutar 1 migración"
	@echo "  migrate-down-1     - Revertir 1 migración"
	@echo "  migrate-status     - Ver estado de migraciones"
	@echo "  create-migration   - Crear nueva migración (NAME=nombre)"
	@echo "  seed               - Ejecutar seeds de datos"
	@echo ""
	@echo " Desarrollo:"
	@echo "  deps               - Descargar y limpiar dependencias"
	@echo "  fmt                - Formatear código"
	@echo "  lint               - Ejecutar linter"
	@echo ""
	@echo " Ejemplos:"
	@echo "  make create-migration NAME=add_users_table"
	@echo "  make migrate-force VERSION=001"

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
	@command -v air >/dev/null 2>&1 || { echo "ERROR: Air no está instalado. Instala con: go install github.com/air-verse/air@latest"; exit 1; }
	@air

clean:
	@echo "Limpiando artefactos..."
	@go clean
	@rm -f bin/$(APP_NAME)
	@echo "Limpieza completada"

# TESTING
test:
	@echo "Ejecutando tests..."
	@go test -v ./...

test-coverage:
	@echo "Ejecutando tests con cobertura..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@echo "Reporte de cobertura generado: coverage.out"

# DOCKER
docker-up:
	@echo "Levantando servicios Docker..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Servicios Docker iniciados"

docker-down:
	@echo "Bajando servicios Docker..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "Servicios Docker detenidos"

docker-build:
	@echo "Construyendo imagen Docker..."
	@docker build -t $(APP_NAME) .
	@echo "Imagen Docker construida: $(APP_NAME)"

docker-logs:
	@echo "Mostrando logs de Docker Compose..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# MIGRACIONES
migrate-up:
	@echo "Ejecutando migraciones..."
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=up
	@echo "Migraciones completadas"

migrate-down:
	@echo "Revirtiendo migraciones..."
	@read -p "¿Continuar? [y/N]: " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=down; \
		echo "Migraciones revertidas"; \
	else \
		echo "Operación cancelada"; \
	fi

migrate-up-1:
	@echo "Ejecutando 1 migración..."
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=up -steps=1

migrate-down-1:
	@echo "Revirtiendo 1 migración..."
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=down -steps=1

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: Especifica VERSION=X"; \
		echo "Ejemplo: make migrate-force VERSION=001"; \
		exit 1; \
	fi
	@echo "Forzando migración a versión $(VERSION)..."
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=force -version=$(VERSION)

migrate-version:
	@echo "Versión actual de migración:"
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=version

migrate-drop:
	@echo "PELIGRO: Esto eliminará TODAS las tablas"
	@read -p "¿Estás COMPLETAMENTE seguro? [y/N]: " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=drop; \
		echo "Tablas eliminadas"; \
	else \
		echo "Operación cancelada"; \
	fi

create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: Especifica NAME=nombre_de_la_migracion"; \
		echo "Ejemplo: make create-migration NAME=add_users_table"; \
		exit 1; \
	fi
	@echo "Creando migración: $(NAME)"
	@TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	UP_FILE="$(MIGRATE_PATH)/$${TIMESTAMP}_$(NAME).up.sql"; \
	DOWN_FILE="$(MIGRATE_PATH)/$${TIMESTAMP}_$(NAME).down.sql"; \
	mkdir -p $(MIGRATE_PATH); \
	echo "-- Migration: $(NAME)" > $$UP_FILE; \
	echo "-- Created: $$(date)" >> $$UP_FILE; \
	echo "" >> $$UP_FILE; \
	echo "-- Add your UP migration here" >> $$UP_FILE; \
	echo "-- Migration: $(NAME)" > $$DOWN_FILE; \
	echo "-- Created: $$(date)" >> $$DOWN_FILE; \
	echo "-- Add your DOWN migration here" >> $$DOWN_FILE; \
	echo "Migración creada:"; \
	echo "   UP:   $$UP_FILE"; \
	echo "   DOWN: $$DOWN_FILE"

migrate-status:
	@echo "=== Estado de Migraciones ==="
	@$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=version
	@echo ""
	@echo "=== Archivos de migración disponibles ==="
	@ls -la $(MIGRATE_PATH)/*.sql 2>/dev/null | head -10 || echo "No hay archivos de migración"
	@TOTAL=$$(ls $(MIGRATE_PATH)/*.sql 2>/dev/null | wc -l); \
	if [ $$TOTAL -gt 10 ]; then \
		echo "... y $$(($$TOTAL - 10)) archivos más"; \
	fi

migrate-validate:
	@echo "Validando migraciones..."
	@ERROR=0; \
	for file in $(MIGRATE_PATH)/*.up.sql; do \
		if [ -f "$$file" ]; then \
			basename="$$(basename $$file .up.sql)"; \
			down_file="$(MIGRATE_PATH)/$${basename}.down.sql"; \
			if [ ! -f "$$down_file" ]; then \
				echo "Error: Falta el archivo DOWN para $$file"; \
				ERROR=1; \
			fi; \
		fi; \
	done; \
	if [ $$ERROR -eq 0 ]; then \
		echo "Todas las migraciones son válidas"; \
	else \
		exit 1; \
	fi

migrate-backup:
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "Error: DATABASE_URL no está definida"; \
		exit 1; \
	fi
	@echo "Creando backup de la base de datos..."
	@TIMESTAMP=$$(date +%Y%m%d_%H%M%S); \
	BACKUP_FILE="backups/db_backup_$${TIMESTAMP}.sql"; \
	mkdir -p backups; \
	pg_dump $(DATABASE_URL) > $$BACKUP_FILE 2>/dev/null; \
	if [ $$? -eq 0 ]; then \
		echo "Backup creado: $$BACKUP_FILE"; \
	else \
		echo "Error creando backup"; \
		exit 1; \
	fi

migrate-up-safe: migrate-backup migrate-up
	@echo "Migración completada con backup de seguridad"

# DATOS DE PRUEBA
seed:
	@echo "Ejecutando seeds..."
	@go run cmd/seed/main.go
	@echo "Seeds completados"

# HERRAMIENTAS DE DESARROLLO
deps:
	@echo "Descargando dependencias..."
	@go mod download
	@go mod tidy
	@echo "Dependencias actualizadas"

fmt:
	@echo "Formateando código..."
	@go fmt ./...
	@echo "Código formateado"

lint:
	@echo "Ejecutando linter..."
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint no está instalado"; exit 1; }
	@golangci-lint run
	@echo "Linting completado"

# COMANDOS COMBINADOS
setup: deps docker-up migrate-up seed
	@echo "🎉 Entorno configurado completamente"

check: fmt lint test
	@echo "Verificación completa exitosa"

rebuild: clean build
	@echo "Reconstrucción completada"
