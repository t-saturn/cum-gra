# ============ Builder ============
FROM golang:1.24-bookworm AS builder

ENV CGO_ENABLED=0 \
  GO111MODULE=on \
  GOTOOLCHAIN=auto \
  GOPROXY=https://proxy.golang.org,direct \
  GOFLAGS=-mod=mod
# asegura que go pueda escribir go.sum si hace falta

WORKDIR /src

# 1. Pre-cache: solo go.mod (y go.sum si existiera en el repo)
COPY go.mod ./
# Si tienes go.sum en el repo, INCLÚYELO para más cache:
# COPY go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# 2. Copia el resto del código
COPY . .

# 3. Genera/actualiza go.sum con TODAS las dependencias (transitivas incluidas)
RUN --mount=type=cache,target=/go/pkg/mod go mod tidy

# 4. Compila
RUN --mount=type=cache,target=/go/pkg/mod go build -o /out/server ./cmd/server/main.go

# (opcional) crea carpeta de logs para runtime distroless:nonroot
RUN mkdir -p /out/logs

# ============ Runtime ============
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app

COPY --from=builder /out/server /app/server
# copia la carpeta logs con ownership nonroot
COPY --from=builder --chown=nonroot:nonroot /out/logs /app/logs

EXPOSE 9190
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
