.PHONY: build run test clean docker-up docker-down migrate seed dev deps fmt lint

# Variables
BINARY_NAME=github.com/t-saturn/central-user-manager # modify this path
DOCKER_COMPOSE_FILE=docker-compose.yml
MIGRATE_PATH := internal/adapters/repositories/postgres/migrations
MIGRATE_CMD := go run cmd/migrate/main.go


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

# Ayuda
.PHONY: help
help:
	@echo "Comandos disponibles para migraciones:"
	@echo "  migrate-up         - Ejecutar todas las migraciones pendientes"
	@echo "  migrate-down       - Revertir todas las migraciones"
	@echo "  migrate-up-1       - Ejecutar 1 migración hacia arriba"
	@echo "  migrate-down-1     - Revertir 1 migración"
	@echo "  migrate-force      - Forzar migración a una versión específica (usar VERSION=X)"
	@echo "  migrate-version    - Mostrar versión actual de migración"
	@echo "  migrate-drop       - Eliminar todas las tablas (¡CUIDADO!)"
	@echo "  create-migration   - Crear nueva migración (usar NAME=nombre_migracion)"

# Ejecutar todas las migraciones
.PHONY: migrate-up
migrate-up:
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=up

# Revertir todas las migraciones
.PHONY: migrate-down
migrate-down:
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=down

# Ejecutar 1 migración hacia arriba
.PHONY: migrate-up-1
migrate-up-1:
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=up -steps=1

# Revertir 1 migración
.PHONY: migrate-down-1
migrate-down-1:
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=down -steps=1

# Forzar migración a versión específica
.PHONY: migrate-force
migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: Especifica VERSION=X"; \
		exit 1; \
	fi
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=force -version=$(VERSION)

# Mostrar versión actual
.PHONY: migrate-version
migrate-version:
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=version

# Eliminar todas las tablas
.PHONY: migrate-drop
migrate-drop:
	@echo "¿Estás seguro de que quieres eliminar todas las tablas? [y/N]"
	@read -r REPLY; \
	if [ "$$REPLY" = "y" ] || [ "$$REPLY" = "Y" ]; then \
		$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=drop; \
	else \
		echo "Operación cancelada"; \
	fi

# Crear nueva migración
.PHONY: create-migration
create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: Especifica NAME=nombre_de_la_migracion"; \
		exit 1; \
	fi
	@TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	UP_FILE="$(MIGRATE_PATH)/$${TIMESTAMP}_$(NAME).up.sql"; \
	DOWN_FILE="$(MIGRATE_PATH)/$${TIMESTAMP}_$(NAME).down.sql"; \
	echo "-- $(NAME).up.sql" > $$UP_FILE; \
	echo "" >> $$UP_FILE; \
	echo "-- Agregar cambios aquí" >> $$UP_FILE; \
	echo "-- $(NAME).down.sql" > $$DOWN_FILE; \
	echo "" >> $$DOWN_FILE; \
	echo "-- Revertir cambios aquí" >> $$DOWN_FILE; \
	echo "Migración creada:"; \
	echo "  UP: $$UP_FILE"; \
	echo "  DOWN: $$DOWN_FILE"

# Verificar estado de migraciones
.PHONY: migrate-status
migrate-status:
	@echo "=== Estado de Migraciones ==="
	$(MIGRATE_CMD) -path=$(MIGRATE_PATH) -cmd=version
	@echo ""
	@echo "=== Archivos de migración disponibles ==="
	@ls -la $(MIGRATE_PATH)/*.sql 2>/dev/null || echo "No hay archivos de migración"

# Validar migraciones
.PHONY: migrate-validate
migrate-validate:
	@echo "Validando migraciones..."
	@for file in $(MIGRATE_PATH)/*.up.sql; do \
		if [ -f "$$file" ]; then \
			basename="$$(basename $$file .up.sql)"; \
			down_file="$(MIGRATE_PATH)/$${basename}.down.sql"; \
			if [ ! -f "$$down_file" ]; then \
				echo "Error: No existe el archivo down para $$file"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "Todas las migraciones son válidas"

# Backup de base de datos antes de migrar
.PHONY: migrate-backup
migrate-backup:
	@echo "Creando backup de la base de datos..."
	@TIMESTAMP=$$(date +%Y%m%d_%H%M%S); \
	BACKUP_FILE="backups/db_backup_$${TIMESTAMP}.sql"; \
	mkdir -p backups; \
	pg_dump $(DATABASE_URL) > $$BACKUP_FILE; \
	echo "Backup creado: $$BACKUP_FILE"

# Migrar con backup automático
.PHONY: migrate-up-safe
migrate-up-safe: migrate-backup migrate-up
	@echo "Migración completada con backup de seguridad"
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
