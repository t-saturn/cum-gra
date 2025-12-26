# ============ Builder ============
FROM golang:1.24.4-bookworm AS builder

ENV CGO_ENABLED=0 \
  GO111MODULE=on \
  GOTOOLCHAIN=auto \
  GOPROXY=https://proxy.golang.org,direct \
  GOFLAGS=-mod=mod

WORKDIR /src

COPY go.mod ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod go mod tidy

RUN --mount=type=cache,target=/go/pkg/mod go build -ldflags="-s -w" -o /out/server  ./cmd/server
RUN --mount=type=cache,target=/go/pkg/mod go build -ldflags="-s -w" -o /out/migrate ./cmd/migrate
RUN --mount=type=cache,target=/go/pkg/mod go build -ldflags="-s -w" -o /out/seed    ./cmd/seed

# ============ Runtime ============
FROM alpine:3.20

RUN apk add --no-cache ca-certificates netcat-openbsd

WORKDIR /app

COPY --from=builder /out/server  /app/server
COPY --from=builder /out/migrate /app/migrate
COPY --from=builder /out/seed    /app/seed

COPY internal/database /app/internal/database
COPY internal/data     /app/internal/data

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
  /app/migrate -cmd=up && break
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