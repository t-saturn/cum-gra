# ============ Builder ============
FROM golang:1.24.4-bookworm AS builder

ENV CGO_ENABLED=0 \
  GO111MODULE=on \
  GOTOOLCHAIN=auto

WORKDIR /src

# Cache de dependencias
COPY go.mod ./
# COPY go.sum ./   # descomenta si lo tienes para cache más estable
RUN go mod download

# Código
COPY . .

# Binarios: server, migrate y seed
RUN go build -ldflags="-s -w" -o /out/server  ./cmd/server
RUN go build -ldflags="-s -w" -o /out/migrate ./cmd/migrate
RUN go build -ldflags="-s -w" -o /out/seed    ./cmd/seed

# ============ Runtime ============
# Usamos Alpine para poder ejecutar un script de arranque (sh) y nc
FROM alpine:3.20

RUN apk add --no-cache ca-certificates netcat-openbsd

WORKDIR /app

# Binarios
COPY --from=builder /out/server  /app/server
COPY --from=builder /out/migrate /app/migrate
COPY --from=builder /out/seed    /app/seed

# Archivos de migraciones/seeds (ajusta si tu repo los tiene en otra ruta)
COPY internal/database /app/internal/database

# Entrypoint: espera a Postgres, ejecuta migrate/seed y arranca el server
COPY <<'SH' /entrypoint.sh
#!/bin/sh
set -eu

DB_HOST="${DB_HOST:-postgres}"
DB_PORT="${DB_PORT:-5432}"

echo "[entrypoint] Esperando a PostgreSQL en ${DB_HOST}:${DB_PORT}..."
# Espera TCP
until nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 2
done

echo "[entrypoint] Ejecutando migraciones..."
# Reintentos por si el motor acepta TCP pero aún no está listo
for i in 1 2 3 4 5; do
  /app/migrate -cmd=up -path=/app/internal/database/migrations && break
  echo "[entrypoint] migrate fallo, reintento $i/5..."
  sleep 3
done

echo "[entrypoint] Ejecutando seeds..."
for i in 1 2 3; do
  /app/seed && break
  echo "[entrypoint] seed fallo, reintento $i/3..."
  sleep 2
done

echo "[entrypoint] Iniciando servidor..."
exec /app/server
SH

RUN chmod +x /entrypoint.sh

EXPOSE 9191
ENTRYPOINT ["/entrypoint.sh"]
