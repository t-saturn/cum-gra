# ============ Builder ============
FROM golang:1.24.4-bookworm AS builder

ENV CGO_ENABLED=0 \
  GO111MODULE=on \
  GOTOOLCHAIN=auto \
  GOPROXY=https://proxy.golang.org,direct \
  GOFLAGS=-mod=mod

WORKDIR /src

# 1) Pre-cache: solo go.mod (si tienes go.sum en el repo y quieres usarlo, puedes copiarlo también)
COPY go.mod ./
# COPY go.sum ./
# (opcional con BuildKit) cache de módulos
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# 2) Código
COPY . .

# 3) Asegura que se resuelvan TODAS las dependencias transitivas y se genere go.sum
RUN --mount=type=cache,target=/go/pkg/mod go mod tidy

# 4) Binarios: server, migrate y seed (mejor por paquete que por archivo)
RUN --mount=type=cache,target=/go/pkg/mod go build -ldflags="-s -w" -o /out/server  ./cmd/server
RUN --mount=type=cache,target=/go/pkg/mod go build -ldflags="-s -w" -o /out/migrate ./cmd/migrate
RUN --mount=type=cache,target=/go/pkg/mod go build -ldflags="-s -w" -o /out/seed    ./cmd/seed

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
COPY internal/data     /app/internal/data

# Entrypoint: espera a Postgres, ejecuta migrate/seed y arranca el server
COPY <<'SH' /entrypoint.sh
#!/bin/sh
set -eu

DB_HOST="${DB_HOST:-postgres}"
DB_PORT="${DB_PORT:-5432}"

echo "[entrypoint] Esperando a PostgreSQL en ${DB_HOST}:${DB_PORT}..."
until nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 2
done

echo "[entrypoint] Ejecutando migraciones..."
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
